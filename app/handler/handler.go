package handler

import (
	"github.com/AidansCode/img-shr/app/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	PostService service.PostService
}

type ErrorResponse struct {
	StatusCode int
	ErrorMsg string
}

func (h *Handler) renderError(c echo.Context, statusCode int, errorMsg string) error {
	return c.Render(statusCode, "error.error", ErrorResponse{StatusCode: statusCode, ErrorMsg: errorMsg})
}
