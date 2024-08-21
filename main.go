package main

import (
	"net/http"

	"github.com/AidansCode/img-shr/renderer"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = renderer.NewRenderer()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/upload", func(c echo.Context) error {
		return c.Render(http.StatusOK, "upload", nil)
	})

	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8000"))
}