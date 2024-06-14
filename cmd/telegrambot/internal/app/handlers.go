package app

import (
	"context"
	"log"
)

func (a *App) CheckChatExist(ctx context.Context, chatID int) bool {
	return a.repo.CheckChatExist(ctx, chatID)
}

func (a *App) CheckFinished(ctx context.Context, chatID int) bool {
	return a.repo.CheckFinished(ctx, chatID)
}

func (a *App) CreateChat(ctx context.Context, chatID int, messageID int) {
	upd := UserInfo{
		ChatID:           chatID,
		QuestNumber:      0,
		LastMessageID:    messageID,
		CountRightAnswer: 0,
		Answer:           "",
		Finished:         false,
		Quests:           getQuestion(),
	}
	a.repo.CreateChat(ctx, upd)
}

func (a *App) SaveMessage(ctx context.Context, user UserInfo) {
	a.repo.SaveMessage(ctx, user)
}

func (a *App) PlusCounter(ctx context.Context, u UserInfo) {
	u.QuestNumber++
	a.repo.PlusCounter(ctx, u)
}

func (a *App) GetInfo(ctx context.Context, chatID int) UserInfo {
	upd := a.repo.GetInfo(ctx, chatID)

	return upd
}

//func (a *App) GetQuestions(ctx context.Context, u UserInfo) UserInfo {
//	u.Quests = getQuestion()
//	//quest := QuizEasy[u.QuestNumber]
//	return u
//}

func (a *App) CheckAnswer(ctx context.Context, answer string, upd UserInfo) {
	quests := upd.Quests[upd.QuestNumber]
	log.Println(quests.GoodAnswer, answer, upd.CountRightAnswer)
	if quests.GoodAnswer == answer {
		upd.CountRightAnswer++
		a.repo.PlusAnswer(ctx, upd)
	}

	if upd.QuestNumber+1 == len(upd.Quests) {
		a.repo.SetFinished(ctx, upd)
	}
}
