package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"tgbot/cmd/telegrambot/internal/adapters/repo"
	"tgbot/cmd/telegrambot/internal/api/telegram"
	"tgbot/cmd/telegrambot/internal/app"
)

// Send any text message to the bot after the bot has been started

var token = ""

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db, err := repo.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			slog.Warn("db.Close:", err)
		}
	}()
	myapp := app.New(db)
	bot := telegram.New(myapp, token)

	log.Println("Bot is now running.  Press CTRL-C to exit.")

	bot.Start(ctx)
}
