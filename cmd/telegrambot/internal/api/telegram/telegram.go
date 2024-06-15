package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
	"tgbot/cmd/telegrambot/internal/app"
)

type application interface {
	CheckChatExist(context.Context, int) bool
	CheckFinished(context.Context, int) bool
	CreateChat(context.Context, int, int) error
	SaveMessage(context.Context, app.UserInfo) error
	GetInfo(context.Context, int) (app.UserInfo, error)
	PlusCounter(context.Context, app.UserInfo) error
	CheckAnswer(context.Context, string, app.UserInfo) error
}

type api struct {
	app application
}

func New(a application, token string) *bot.Bot {
	api := api{
		app: a,
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(api.DefaultHandler),
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, api.CallbackHandler),
	}
	b, err := bot.New(token, opts...)
	if nil != err {
		log.Fatal(err)
	}

	return b
}
