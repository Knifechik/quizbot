package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"log/slog"
	"strconv"
	"strings"
	"tgbot/cmd/telegrambot/internal/app"
)

func (a *api) CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	answer := update.CallbackQuery.Data

	upd, err := a.app.GetInfo(ctx, int(update.CallbackQuery.Message.Message.Chat.ID))
	if err != nil {
		slog.Error("app.GetInfo:", err)
		return
	}

	if !strings.Contains(answer, "button_0") {
		a.app.PlusCounter(ctx, upd)
		a.app.CheckAnswer(ctx, answer, upd)
	}

	upd, err = a.app.GetInfo(ctx, int(update.CallbackQuery.Message.Message.Chat.ID))
	if err != nil {
		slog.Error("app.GetInfo:", err)
		return
	}

	log.Println(upd.CountRightAnswer)
	if a.app.CheckFinished(ctx, int(update.CallbackQuery.Message.Message.Chat.ID)) {
		a.app.SaveMessage(ctx, app.UserInfo{
			ChatID:           upd.ChatID,
			QuestNumber:      upd.QuestNumber,
			LastMessageID:    upd.LastMessageID,
			CountRightAnswer: upd.CountRightAnswer,
			Answer:           upd.Answer,
			Finished:         upd.Finished,
			Quests:           upd.Quests,
			UserAnswers:      append(upd.UserAnswers[1:], answer),
		})
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    upd.ChatID,
			MessageID: upd.LastMessageID,
			Text: "Поздравляю, Вы закончили, правильных ответов: " +
				strconv.Itoa(upd.CountRightAnswer),
		})
		return
	}

	keyboard := createKeyboard(upd.Quests[upd.QuestNumber].Answers)

	message, _ := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      upd.ChatID,
		MessageID:   upd.LastMessageID,
		Text:        upd.Quests[upd.QuestNumber].Quest,
		ReplyMarkup: keyboard,
	})

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	err = a.app.SaveMessage(ctx, app.UserInfo{
		ChatID:           upd.ChatID,
		QuestNumber:      upd.QuestNumber,
		LastMessageID:    message.ID,
		CountRightAnswer: upd.CountRightAnswer,
		Answer:           upd.Answer,
		Finished:         upd.Finished,
		Quests:           upd.Quests,
		UserAnswers:      append(upd.UserAnswers, answer),
	})
	if err != nil {
		slog.Error("app.SaveMessage:", err)
		return
	}

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

	err := a.app.CreateChat(ctx, int(update.Message.Chat.ID), message.ID)
	if err != nil {
		slog.Error("app.CreateChat:", err)
		return
	}

}
