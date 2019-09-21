package main

import (
	"context"
	"log"
	"os"

	"github.com/petuhovskiy/grpc-hydra-bench/telegram/desc"
	"github.com/petuhovskiy/grpc-hydra-bench/telegram/logic"
	"github.com/petuhovskiy/grpc-hydra-bench/telegram/store"
	"github.com/petuhovskiy/grpc-hydra-bench/telegram/xctx"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	bolt "go.etcd.io/bbolt"
)

func createBot(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	return bot
}

func createUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	return updates
}

func main() {
	botToken := os.Getenv("BOT_TOKEN")

	db, err := bolt.Open("./bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = store.Init(db)
	if err != nil {
		log.Fatal(err)
	}

	bot := createBot(botToken)
	updates := createUpdates(bot)

	dsc := &desc.Main{
		Bot: bot,
		DB:  db,
	}

	root := context.Background()
	root = context.WithValue(root, xctx.Desc, dsc)

	log.Println("Bot initialized")

	for update := range updates {
		ctx := root
		ctx = context.WithValue(ctx, xctx.Update, update)

		logic.HandleUpdate(ctx, update)
	}
}
