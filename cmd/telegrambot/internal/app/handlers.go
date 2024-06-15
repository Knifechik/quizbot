package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
)

func (a *App) CheckChatExist(ctx context.Context, chatID int) bool {
	check, err := a.repo.CheckChatExist(ctx, chatID)
	if err != nil {
		slog.Error("repo.CheckChatExist", err)
		return true
	}
	return check
}

func (a *App) CheckFinished(ctx context.Context, chatID int) bool {
	check, err := a.repo.CheckFinished(ctx, chatID)
	if err != nil {
		slog.Error("repo.CheckFinished", err)
		return false
	}
	return check
}

func (a *App) CreateChat(ctx context.Context, chatID int, messageID int) error {
	upd := UserInfo{
		ChatID:           chatID,
		QuestNumber:      0,
		LastMessageID:    messageID,
		CountRightAnswer: 0,
		Answer:           "",
		Finished:         false,
		Quests:           getQuestion(),
		UserAnswers:      make([]string, 0),
	}

	err := a.repo.CreateChat(ctx, upd)
	if err != nil {
		return fmt.Errorf("repo.CreateChat: %w", err)
	}

	return nil
}

func (a *App) SaveMessage(ctx context.Context, user UserInfo) error {
	err := a.repo.SaveMessage(ctx, user)
	if err != nil {
		return fmt.Errorf("repo.SaveMessage: %w", err)
	}

	return nil
}

func (a *App) PlusCounter(ctx context.Context, u UserInfo) error {
	u.QuestNumber++

	err := a.repo.PlusCounter(ctx, u)
	if err != nil {
		return fmt.Errorf("repo.PlusCounter: %w", err)
	}

	return nil
}

func (a *App) GetInfo(ctx context.Context, chatID int) (UserInfo, error) {
	upd, err := a.repo.GetInfo(ctx, chatID)
	if err != nil {
		return UserInfo{}, fmt.Errorf("repo.GetInfo: %v", err)
	}

	return upd, nil
}

func (a *App) CheckAnswer(ctx context.Context, answer string, upd UserInfo) error {
	quests := upd.Quests[upd.QuestNumber]
	log.Println(quests.GoodAnswer, "=+=+=", answer)
	if quests.GoodAnswer == answer {
		upd.CountRightAnswer++

		err := a.repo.PlusAnswer(ctx, upd)
		if err != nil {
			return fmt.Errorf("repo.PlusAnswer: %v", err)
		}
	}

	if upd.QuestNumber+1 == len(upd.Quests) {
		err := a.repo.SetFinished(ctx, upd)
		if err != nil {
			return fmt.Errorf("repo.SetFinished: %v", err)
		}
	}

	return nil
}
