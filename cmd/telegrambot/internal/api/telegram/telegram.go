package telegram

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"tgbot/cmd/telegrambot/internal/app"
)

type application interface {
	Create(context.Context, int) (*app.UserInfo, error)
	Save(context.Context, app.UserInfo) error
	Get(context.Context, int) (*app.UserInfo, error)
	CheckFinished(context.Context, int) (bool, error)
	CheckAnswer(context.Context, string, int) (*app.UserInfo, error)
}

type api struct {
	app application
}

func New(a application, token string) (*bot.Bot, error) {
	api := api{
		app: a,
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(api.DefaultHandler),
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, api.CallbackHandler),
	}
	b, err := bot.New(token, opts...)
	if nil != err {
		return nil, fmt.Errorf("create telegram bot: %w", err)
	}

	return b, nil
}
