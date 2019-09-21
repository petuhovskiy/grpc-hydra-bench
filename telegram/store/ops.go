package store

import (
	"encoding/json"
	"fmt"

	"github.com/petuhovskiy/grpc-hydra-bench/telegram/state"
	bolt "go.etcd.io/bbolt"
)

var name = []byte("sessions")

func Init(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(name)
		return err
	})
}

func Load(db *bolt.DB, chatID int64) (*state.Session, error) {
	var s state.Session

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(name)
		val := bucket.Get(str(chatID))
		if val == nil {
			return nil
		}

		return json.Unmarshal(val, &s)
	})

	return &s, err
}

func Save(db *bolt.DB, chatID int64, s *state.Session) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(name)
		val, err := json.Marshal(s)
		if err != nil {
			return err
		}

		return bucket.Put(str(chatID), val)
	})
}

func str(key int64) []byte {
	return []byte(fmt.Sprintf("%d", key))
}
