package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UploadForm(c echo.Context) error {
	return c.Render(http.StatusOK, "upload", UploadFormData{Title: "", Error: ""})
}

func (h *Handler) Upload(c echo.Context) error {
	title := strings.TrimSpace(c.FormValue("title"))
	if len(title) == 0 {
		return c.Render(http.StatusOK, "upload.upload-form", UploadFormData{Title: title, Error: "Missing title"})
	}
	return c.Render(http.StatusOK, "upload", nil)
}

type UploadFormData struct {
	Title string
	Error string
}
