package handler

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var allowedImageTypes = []string{
	"image/jpeg", "image/png", "image/gif",
}

func (h *Handler) UploadForm(c echo.Context) error {
	return c.Render(http.StatusOK, "upload", UploadFormData{Title: "", Error: ""})
}

func (h *Handler) Upload(c echo.Context) error {
	title := strings.TrimSpace(c.FormValue("title"))
	if len(title) == 0 {
		return c.Render(http.StatusOK, "upload.upload-form", UploadFormData{Title: title, Error: "Missing title"})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Render(http.StatusOK, "upload.upload-form", UploadFormData{Title: title, Error: "Missing image"})
	}
	src, err := file.Open()
	if err != nil {
		return c.Render(http.StatusOK, "upload.upload-form", UploadFormData{Title: title, Error: "Invalid image uploaded"})
	}
	defer src.Close()

	if err := validateMimeType(src); err != nil {
		return c.Render(http.StatusOK, "upload.upload-form", UploadFormData{Title: title, Error: err.Error()})
	}

	return c.Render(http.StatusOK, "upload.upload-form", UploadFormData{Title: title, Error: "Okay!"})
}

func validateMimeType(f multipart.File) error {
	contents, err := io.ReadAll(f)
	if err != nil {
		return errors.New("invalid image uploaded")
	}

	mimeType := http.DetectContentType(contents)
	for _, allowedMimeType := range allowedImageTypes {
		if allowedMimeType == mimeType {
			return nil
		}
	}

	return errors.New("invalid file type. Must be one of: " + strings.Join(allowedImageTypes, ", "))
}

type UploadFormData struct {
	Title string
	Error string
}
