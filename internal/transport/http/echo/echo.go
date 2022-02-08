package echo


import (
	"context"
	"github.com/alidevjimmy/snapp-clean/internal/transport/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

type rest struct {
	echo *echo.Echo
	//user
}

func New() http.Rest {
	return &rest{
		echo : echo.New(),
	}
}

func (r *rest) Start(address string) error {
	r.echo.Use(middleware.Recover())
	r.routing()

	return r.echo.Start(address)
}

func (r *rest) ShutDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return r.echo.Shutdown(ctx)
}