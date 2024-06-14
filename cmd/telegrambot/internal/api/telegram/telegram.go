package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"tgbot/cmd/telegrambot/internal/app"
)

type application interface {
	CheckChatExist(context.Context, int) bool
	CheckFinished(context.Context, int) bool
	CreateChat(context.Context, int, int)
	SaveMessage(context.Context, app.UserInfo)
	GetInfo(context.Context, int) app.UserInfo
	PlusCounter(context.Context, app.UserInfo)
	//GetQuestions(context.Context, app.UserInfo) app.UserInfo
	CheckAnswer(context.Context, string, app.UserInfo)
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
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, api.CallbackHandler),
	}
	b, err := bot.New(token, opts...)
	if nil != err {
		panic(err)
	}

	return b
}
