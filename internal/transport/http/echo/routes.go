package echo

func (r *rest) routing() {
	r.echo.GET("/api/v1/login", r.user.Login)
	r.echo.GET("/api/v1/register", r.user.Register)
}