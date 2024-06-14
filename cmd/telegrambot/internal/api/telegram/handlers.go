package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
	"strings"
	"tgbot/cmd/telegrambot/internal/app"
)

func (a *api) CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	upd := a.app.GetInfo(ctx, int(update.CallbackQuery.Message.Message.Chat.ID))

	answer := update.CallbackQuery.Data

	if !strings.Contains(answer, "button_0") {
		a.app.PlusCounter(ctx, upd)
		a.app.CheckAnswer(ctx, answer, upd)
	}

	upd = a.app.GetInfo(ctx, int(update.CallbackQuery.Message.Message.Chat.ID))

	if a.app.CheckFinished(ctx, int(update.CallbackQuery.Message.Message.Chat.ID)) {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    upd.ChatID,
			MessageID: upd.LastMessageID,
			Text: "Поздравляю, Вы закончили, правильных ответов: " +
				strconv.Itoa(upd.CountRightAnswer),
		})
		return
	}

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: upd.Quests[upd.QuestNumber].Answers[0], CallbackData: "button_1"},
				{Text: upd.Quests[upd.QuestNumber].Answers[1], CallbackData: "button_2"},
			}, {
				{Text: upd.Quests[upd.QuestNumber].Answers[2], CallbackData: "button_3"},
				{Text: upd.Quests[upd.QuestNumber].Answers[3], CallbackData: "button_4"},
			},
		},
	}

	message, _ := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      upd.ChatID,
		MessageID:   upd.LastMessageID,
		Text:        upd.Quests[upd.QuestNumber].Quest,
		ReplyMarkup: kb,
	})

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	a.app.SaveMessage(ctx, app.UserInfo{
		ChatID:           upd.ChatID,
		QuestNumber:      upd.QuestNumber,
		LastMessageID:    message.ID,
		CountRightAnswer: upd.CountRightAnswer,
		Answer:           upd.Answer,
		Finished:         upd.Finished,
		Quests:           upd.Quests,
	})

}

func (a *api) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if a.app.CheckChatExist(ctx, int(update.Message.Chat.ID)) {
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: update.Message.ID,
		})
		return
	}

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Поехали", CallbackData: "button_0"},
			},
		},
	}

	message, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Попробуй решить наши задачки, там еще всякое добро, короче разбирайся, а я тебя писать буду.",
		ReplyMarkup: kb,
	})

	a.app.CreateChat(ctx, int(update.Message.Chat.ID), message.ID)

}
