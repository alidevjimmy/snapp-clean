package controller

import (
	"github.com/alidevjimmy/snapp-clean/internal/pkg/logger"
	"github.com/alidevjimmy/snapp-clean/internal/service"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	logger logger.Logger
	userService service.User
}

func (u *UserController) Login(ctx echo.Context) error {
	panic("not implemented!")
}

func (u *UserController) Register(ctx echo.Context) error {
	panic("not implemented!")
}

