package users

import "github.com/jmoiron/sqlx"

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) FindByUsername(username string) (*User, error) {
	var u User
	err := r.db.Get(&u, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
