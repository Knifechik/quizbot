package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
)

func (a *App) CheckExist(ctx context.Context, chatID int) error {
	_, err := a.repo.Get(ctx, chatID)
	switch {
	case errors.Is(err, ErrNotFound):
		return nil
	default:
		return fmt.Errorf("chat exist or repo's error: %w", err)
	}
}

func (a *App) CheckFinished(ctx context.Context, chatID int) (bool, error) {
	user, err := a.repo.Get(ctx, chatID)
	if err != nil {
		return false, fmt.Errorf("repo.Get: %w", err)
	}
	return user.Finished, nil
}

func (a *App) Create(ctx context.Context, chatID int) (UserInfo, error) {

	upd := UserInfo{
		ChatID:           chatID,
		QuestNumber:      0,
		LastMessageID:    0,
		CountRightAnswer: 0,
		Answer:           "",
		Finished:         false,
		Quests:           getQuestion(),
	}

	err := a.repo.Create(ctx, upd)
	if err != nil {
		return UserInfo{}, fmt.Errorf("repo.Create: %w", err)
	}

	return upd, nil
}

func (a *App) Save(ctx context.Context, user UserInfo) error {
	slog.Info("Save:", user.LastMessageID)
	err := a.repo.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("repo.Update: %w", err)
	}

	return nil
}

func (a *App) Get(ctx context.Context, chatID int) (UserInfo, error) {
	upd, err := a.repo.Get(ctx, chatID)
	if err != nil {
		return UserInfo{}, fmt.Errorf("repo.Get: %v", err)
	}

	return upd, nil
}

func (a *App) CheckAnswer(ctx context.Context, answer string, chatID int) (UserInfo, error) {
	user, err := a.repo.Get(ctx, chatID)
	if err != nil {
		return UserInfo{}, fmt.Errorf("repo.Get: %v", err)
	}

	if answer == "button_0" {
		return user, nil
	}

	user.Quests[user.QuestNumber].UserAnswer = answer
	quests := user.Quests[user.QuestNumber]

	log.Println(quests.GoodAnswer, "=+=+=", answer)

	if quests.GoodAnswer == answer {
		user.CountRightAnswer++
	}

	if user.QuestNumber+1 == len(user.Quests) {
		user.Finished = true
	}

	user.QuestNumber++

	err = a.repo.Update(ctx, user)
	if err != nil {
		return UserInfo{}, fmt.Errorf("repo.SetFinished: %v", err)
	}

	return user, nil
}

//func (a *App) SaveAnswer(ctx context.Context, user UserInfo, answer string, message int) error {
//	if answer != "button_0" {
//		user.Quests[user.QuestNumber].UserAnswer = answer
//	}
//	user.LastMessageID = message
//	err := a.repo.Update(ctx, user)
//	if err != nil {
//		return fmt.Errorf("repo.PlusAnswer: %v", err)
//	}
//	return nil
//}
