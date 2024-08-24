package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Home(c echo.Context) error {
	posts, err := h.PostService.Latest(25)
	if err != nil {
		return h.renderError(c, http.StatusInternalServerError, "Failed to load latest posts")
	}

	return c.Render(http.StatusOK, "index", posts)
}
