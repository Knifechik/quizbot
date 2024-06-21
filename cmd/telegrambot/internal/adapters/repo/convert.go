package repo

import "tgbot/cmd/telegrambot/internal/app"

type UserInfo struct {
	ChatID           int         `db:"chat_id"`
	QuestNumber      int         `db:"quest_number"`
	LastMessageID    int         `db:"last_message"`
	CountRightAnswer int         `db:"right_answer"`
	Finished         bool        `db:"finished"`
	Quests           []Questions `db:"quests"`
}

type Questions struct {
	Quest      string
	Answers    []string
	GoodAnswer string
	UserAnswer string
}

func (u *UserInfo) convert(q []app.Questions) *app.UserInfo {
	res := &app.UserInfo{
		ChatID:           u.ChatID,
		QuestNumber:      u.QuestNumber,
		LastMessageID:    u.LastMessageID,
		CountRightAnswer: u.CountRightAnswer,
		Finished:         u.Finished,
	}
	res.Quests = q
	return res
}
