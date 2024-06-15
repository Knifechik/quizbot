package telegram

import "github.com/go-telegram/bot/models"

func createKeyboard(answers []string) models.InlineKeyboardMarkup {
	return models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: answers[0], CallbackData: answers[0]},
				{Text: answers[1], CallbackData: answers[1]},
			}, {
				{Text: answers[2], CallbackData: answers[2]},
				{Text: answers[3], CallbackData: answers[3]},
			},
		},
	}
}
