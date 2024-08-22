package main

import (
	"github.com/AidansCode/img-shr/app"
	"github.com/AidansCode/img-shr/renderer"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = renderer.NewRenderer()
	e.Static("/static", "static")

	app := app.NewApp("imgshr.sqlite")
	app.HandleRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
