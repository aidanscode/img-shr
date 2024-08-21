package app

import (
	"github.com/AidansCode/img-shr/app/handler"
	"github.com/AidansCode/img-shr/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	Handler handler.Handler
}

func NewApp() *App {
	handler := handler.Handler{PostService: service.NewPostService()}
	return &App{Handler: handler}
}

func (app *App) HandleRoutes(e *echo.Echo) {
	e.GET("/", app.Handler.Home)
	e.GET("/upload", app.Handler.Upload)
}
