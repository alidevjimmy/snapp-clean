package service

import (
	"context"
	"github.com/alidevjimmy/snapp-clean/internal/entity/model"
)

type User interface {
	Login(ctx context.Context, username , passwords string) error
	Register(ctx context.Context, user model.User) error
}