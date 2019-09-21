package logic

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/petuhovskiy/grpc-hydra-bench/telegram/store"
	"github.com/petuhovskiy/grpc-hydra-bench/telegram/xctx"
)

func HandleUpdate(ctx context.Context, upd tgbotapi.Update) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic", r)
		}
	}()

	if upd.Message == nil {
		return
	}

	msg := upd.Message
	chatID := msg.Chat.ID

	if msg.From == nil || chatID != int64(msg.From.ID) {
		log.Println("Message is not PM")
		return
	}

	desc := xctx.GetDesc(ctx)
	db := desc.DB

	s, err := store.Load(db, chatID)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := store.Save(db, chatID, s)
		if err != nil {
			panic(err)
		}
	}()

	// init session
	if s.ChatID == 0 {
		s.ChatID = chatID
	}

	s.Counter++

	spew.Dump(msg)
	spew.Dump(s)

	desc.Bot.Send(tgbotapi.NewMessage(
		chatID,
		fmt.Sprintf("Counter: %v", s.Counter),
	))
}
