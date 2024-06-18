package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
	"tgbot/cmd/telegrambot/internal/app"
)

type application interface {
	Create(context.Context, int) (app.UserInfo, error)
	Save(context.Context, app.UserInfo) error
	Get(context.Context, int) (app.UserInfo, error)
	CheckExist(context.Context, int) error
	CheckFinished(context.Context, int) (bool, error)
	CheckAnswer(context.Context, string, int) (app.UserInfo, error)
	//SaveAnswer(context.Context, app.UserInfo, string, int) error
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
