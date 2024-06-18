package telegram

import (
	"context"
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
	answer := update.CallbackQuery.Data

	upd, err := a.app.CheckAnswer(ctx, answer, int(update.CallbackQuery.Message.Message.Chat.ID))
	if err != nil {
		slog.Error("app.CheckAnswer:", err)
		return
	}

	finish, err := a.app.CheckFinished(ctx, int(update.CallbackQuery.Message.Message.Chat.ID))
	if err != nil {
		slog.Error("app.CheckFinished:", err)
		return
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
			slog.Error("app.Save:", err)
			return
		}

		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    upd.ChatID,
			MessageID: upd.LastMessageID,
			Text: "Поздравляю, Вы закончили, правильных ответов: " +
				strconv.Itoa(upd.CountRightAnswer),
		})
		if err != nil {
			slog.Error("app.EditMessageText:", err)
			return
		}

		return
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
		slog.Error("b.EditMessageText:", err)
		return
	}

	_, err = b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		slog.Error("b.AnswerCallbackQuery:", err)
		return
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
		slog.Error("app.Update:", err)
		return
	}

}

func (a *api) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	err := a.app.CheckExist(ctx, int(update.Message.Chat.ID))

	if err != nil {
		_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: update.Message.ID,
		})
		if err != nil {
			slog.Error("app.DeleteMessage:", err)
			return
		}
		return
	}

	user, err := a.app.Create(ctx, int(update.Message.Chat.ID))
	if err != nil {
		slog.Error("app.Create:", err)
		return
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
		slog.Error("app.SendMessage:", err)
		return
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
		slog.Error("app.Update:", err)
		return
	}

}
