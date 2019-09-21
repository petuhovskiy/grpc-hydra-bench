package desc

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	bolt "go.etcd.io/bbolt"
)

type Main struct {
	Bot *tgbotapi.BotAPI
	DB  *bolt.DB
}
