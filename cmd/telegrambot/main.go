package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"tgbot/cmd/telegrambot/internal/adapters/repo"
	"tgbot/cmd/telegrambot/internal/api/telegram"
	"tgbot/cmd/telegrambot/internal/app"
)

// Send any text message to the bot after the bot has been started

var token = "7285075184:AAFpkz2skzgxN0nV2kndLiRoLGqHVKryBMs"

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db := repo.New(ctx)
	defer db.Close()

	myapp := app.New(db)
	bot := telegram.New(myapp, token)

	log.Println("Bot is now running.  Press CTRL-C to exit.")

	bot.Start(ctx)
}
