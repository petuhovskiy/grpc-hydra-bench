package logic

import (
	"context"
	"fmt"
	"log"
	"strings"

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

	// actual logic

	sendText := func(text string) {
		desc.Bot.Send(tgbotapi.NewMessage(
			chatID,
			text,
		))
	}

	text := msg.Text
	args := strings.Split(text, " ")

	if len(args) == 0 {
		sendText("only text messages are supported for now")
		return
	}

	r := initRouter()
	r.on("ping", func() { sendText("pong") })
	r.on("count", func() {
		s.Counter++
		sendText(fmt.Sprintf("Counter: %v", s.Counter))
	})
	r.on("token", func() {
		token := args[1]
		s.Token = token
		sendText("token set")
	})
	r.on("userinfo", func() {
		// TODO:
	})

	commandNotFound := func() {
		help := "Command not found.\n\nAvailable commands:"
		for k := range r.cmds {
			help += fmt.Sprintf("\n%s", k)
		}
		sendText(help)
	}

	r.route(args, func() {
		commandNotFound()
	})
}
