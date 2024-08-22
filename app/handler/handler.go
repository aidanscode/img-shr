package handler

import (
	"net/http"

	"github.com/AidansCode/img-shr/app/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	PostService service.PostService
}

func (h *Handler) Home(c echo.Context) error {
	return c.Render(http.StatusOK, "index", h.PostService.Latest(5))
}
