package handler

import (
	"net/http"

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

func (h *Handler) Home(c echo.Context) error {
	posts, err := h.PostService.Latest(5)
	if err != nil {
		return h.renderError(c, http.StatusInternalServerError, "Failed to load latest posts")
	}

	return c.Render(http.StatusOK, "index", posts)
}
