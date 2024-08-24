package main

import (
	"fmt"
	"os"

	"github.com/AidansCode/img-shr/app"
	"github.com/AidansCode/img-shr/flags"
	"github.com/AidansCode/img-shr/renderer"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = renderer.NewRenderer()
	e.Static("/static", "static")

	appConfig, err := flags.GetAppFlags()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	app := app.NewApp(*appConfig)
	app.HandleRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
