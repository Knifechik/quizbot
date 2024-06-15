package app

import "context"

type Repo interface {
	CheckChatExist(context.Context, int) (bool, error)
	CreateChat(context.Context, UserInfo) error
	SaveMessage(context.Context, UserInfo) error
	GetInfo(context.Context, int) (UserInfo, error)
	CheckFinished(context.Context, int) (bool, error)
	PlusCounter(context.Context, UserInfo) error
	PlusAnswer(context.Context, UserInfo) error
	SetFinished(context.Context, UserInfo) error
}
