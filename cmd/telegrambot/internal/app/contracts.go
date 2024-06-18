package app

import "context"

type Repo interface {
	Create(context.Context, UserInfo) error
	Update(context.Context, UserInfo) error
	Get(context.Context, int) (UserInfo, error)
}
