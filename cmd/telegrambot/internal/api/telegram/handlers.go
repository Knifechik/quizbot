package telegram

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log/slog"
	"strconv"
	"tgbot/cmd/telegrambot/internal/app"
)

const (
	callbackButton1 = "1"
	CallbackButton2 = "2"
	CallbackButton3 = "3"
	CallbackButton4 = "4"
)

func (a *api) CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	err := a.CallbackHandle(ctx, b, update)
	if err != nil {
		slog.Error("api.CallbackHandle", err)
	}

}

func (a *api) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	err := a.DefaultHandle(ctx, b, update)
	if err != nil {
		slog.Error("api.DefaultHandle:", err)
	}
}

func (a *api) CallbackHandle(ctx context.Context, b *bot.Bot, update *models.Update) error {
	answer := update.CallbackQuery.Data

	upd, err := a.app.CheckAnswer(ctx, answer, int(update.CallbackQuery.Message.Message.Chat.ID))
	if err != nil {
		return fmt.Errorf("app.CheckAnswer %w", err)
	}

	finish, err := a.app.CheckFinished(ctx, int(update.CallbackQuery.Message.Message.Chat.ID))
	if err != nil {
		return fmt.Errorf("app.CheckFinished %w", err)
	}

	if finish {
		err = a.app.Save(ctx, app.UserInfo{
			ChatID:           upd.ChatID,
			QuestNumber:      upd.QuestNumber,
			LastMessageID:    upd.LastMessageID,
			CountRightAnswer: upd.CountRightAnswer,
			Answer:           upd.Answer,
			Finished:         upd.Finished,
			Quests:           upd.Quests,
		})
		if err != nil {
			return fmt.Errorf("app.Save %w", err)
		}

		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    upd.ChatID,
			MessageID: upd.LastMessageID,
			Text: "Поздравляю, Вы закончили, правильных ответов: " +
				strconv.Itoa(upd.CountRightAnswer),
		})
		if err != nil {
			slog.Error("app.EditMessageText:", err)
			return fmt.Errorf("app.EditMessageText %w", err)
		}

		return nil
	}

	keyboard := models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: upd.Quests[upd.QuestNumber].Answers[0], CallbackData: callbackButton1},
				{Text: upd.Quests[upd.QuestNumber].Answers[1], CallbackData: CallbackButton2},
			}, {
				{Text: upd.Quests[upd.QuestNumber].Answers[2], CallbackData: CallbackButton3},
				{Text: upd.Quests[upd.QuestNumber].Answers[3], CallbackData: CallbackButton4},
			},
		},
	}

	message, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   upd.LastMessageID,
		Text:        upd.Quests[upd.QuestNumber].Quest,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		return fmt.Errorf("b.EditMessageText %w", err)
	}

	_, err = b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		return fmt.Errorf("b.AnswerCallbackQuery %w", err)
	}

	err = a.app.Save(ctx, app.UserInfo{
		ChatID:           upd.ChatID,
		QuestNumber:      upd.QuestNumber,
		LastMessageID:    message.ID,
		CountRightAnswer: upd.CountRightAnswer,
		Answer:           upd.Answer,
		Finished:         upd.Finished,
		Quests:           upd.Quests,
	})
	if err != nil {
		return fmt.Errorf("app.Update %w", err)
	}

	return nil
}

func (a *api) DefaultHandle(ctx context.Context, b *bot.Bot, update *models.Update) error {

	user, err := a.app.Create(ctx, int(update.Message.Chat.ID))

	if err != nil {
		_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: update.Message.ID,
		})

		if err != nil {
			return fmt.Errorf("app.DeleteMessage: %w", err)
		}
		return fmt.Errorf("app.Create: %w", err)
	}

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Поехали", CallbackData: "button_0"},
			},
		},
	}

	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Попробуй решить наши задачки, там еще всякое добро, короче разбирайся, а я тебя писать буду.",
		ReplyMarkup: kb,
	})
	if err != nil {
		return fmt.Errorf("app.SendMessage: %w", err)
	}

	err = a.app.Save(ctx, app.UserInfo{
		ChatID:           user.ChatID,
		QuestNumber:      user.QuestNumber,
		LastMessageID:    message.ID,
		CountRightAnswer: user.CountRightAnswer,
		Answer:           user.Answer,
		Finished:         user.Finished,
		Quests:           user.Quests,
	})
	if err != nil {
		return fmt.Errorf("app.Update: %w", err)
	}

	return nil
}
