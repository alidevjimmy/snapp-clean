package service

import "context"

type User interface {
	Login(ctx context.Context, username , passwords string) error
	Register(ctx context.Context, username , passwords string) error

}