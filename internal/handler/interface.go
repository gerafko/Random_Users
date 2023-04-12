package handler

import (
	"context"
	"test/internal/model"
	"test/pkg/randomuser"
)

type userAccessor interface {
	RandomUserGET(ctx context.Context) (*randomuser.GetRandomUserResponse, error)
}

type userStorage interface {
	InsertList(users []model.User) error
}
