package app

import (
	"github.com/AidansCode/img-shr/app/db"
	"github.com/AidansCode/img-shr/app/handler"
	"github.com/AidansCode/img-shr/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	Handler handler.Handler
}

func NewApp(dbPath string) *App {
	database, err := db.NewDatabase(dbPath)
	if err != nil {
		panic(err)
	}
	database.Migrate()

	handler := handler.Handler{
		PostService: service.NewPostService(database),
	}
	return &App{Handler: handler}
}

func (app *App) HandleRoutes(e *echo.Echo) {
	e.GET("/", app.Handler.Home)
	e.GET("/upload", app.Handler.UploadForm)
	e.POST("/upload", app.Handler.Upload)
}
