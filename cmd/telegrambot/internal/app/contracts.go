package app

import "context"

type Repo interface {
	CheckChatExist(context.Context, int) bool
	CreateChat(context.Context, UserInfo)
	SaveMessage(context.Context, UserInfo)
	GetInfo(context.Context, int) UserInfo
	CheckFinished(context.Context, int) bool
	PlusCounter(context.Context, UserInfo)
	PlusAnswer(context.Context, UserInfo)
	SetFinished(context.Context, UserInfo)
}
