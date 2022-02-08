package echo

import (
	"github.com/labstack/echo/v4"
)

func (r *rest) routing() {
	r.echo.GET("/api/v1/test", func(ctx echo.Context) error{
		return ctx.String( 200, "salam")
	})
}