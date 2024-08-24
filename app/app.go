package app

import (
	"errors"
	"os"

	"github.com/AidansCode/img-shr/app/db"
	"github.com/AidansCode/img-shr/app/handler"
	"github.com/AidansCode/img-shr/app/service"
	"github.com/AidansCode/img-shr/flags"
	"github.com/labstack/echo/v4"
)

type App struct {
	Handler handler.Handler
}

func NewApp(appConfig flags.AppFlags) *App {
	database := getDatabase(appConfig)
	imageDirectory := getImageDirectory(appConfig)

	handler := handler.Handler{
		PostService: service.NewPostService(database, imageDirectory),
	}
	return &App{Handler: handler}
}

func (app *App) HandleRoutes(e *echo.Echo) {
	e.GET("/", app.Handler.Home)
	e.GET("/i/:id", app.Handler.View)
	e.GET("/i/:id/raw", app.Handler.ViewRaw)
	e.GET("/upload", app.Handler.UploadForm)
	e.POST("/upload", app.Handler.Upload)
}

func getDatabase(appConfig flags.AppFlags) *db.Database {
	database, err := db.NewDatabase(appConfig.DBPath)
	if err != nil {
		panic(err)
	}
	database.Migrate()

	return database
}

func getImageDirectory(appConfig flags.AppFlags) string {
	imgPath := appConfig.ImagesPath
	fileInfo, err := os.Stat(imgPath)
	if err != nil  {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}

		os.Mkdir(imgPath, 0754)
	} else if !fileInfo.IsDir() {
		panic("specified image direcory is NOT a directory")
	}

	return imgPath
}
