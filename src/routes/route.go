package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gitlab.com/croc3/umbrellacorptask/src/controllers"
	mdw "gitlab.com/croc3/umbrellacorptask/src/middlewares"
	"gitlab.com/croc3/umbrellacorptask/src/services"
)

type Route struct {
	s    *services.DateService
	c    controllers.Controller
	echo *echo.Echo
}

func NewRoute() (*Route, error) {
	r := &Route{}

	r.s = services.NewDateService()
	r.c = controllers.NewController(r.s)
	r.echo = echo.New()

	r.echo.Use(mdw.IsUserAdmin)

	r.echo.GET("/days", r.c.HandleRequest)

	return r, nil
}

func (r *Route) Process() error {
	fmt.Printf("Server is running on 4000 port")

	err := r.echo.Start(":4000")
	if err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}

	return nil
}
