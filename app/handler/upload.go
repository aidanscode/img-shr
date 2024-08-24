package handler

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/AidansCode/img-shr/model"
	"github.com/labstack/echo/v4"
)

type UploadFormData struct {
	Title string
	Error string
}

var allowedImageTypes = []string{
	"image/jpeg", "image/png", "image/gif",
}

func (h *Handler) UploadForm(c echo.Context) error {
	return c.Render(http.StatusOK, "upload", UploadFormData{Title: "", Error: ""})
}

func (h *Handler) Upload(c echo.Context) error {
	title := strings.TrimSpace(c.FormValue("title"))
	if len(title) == 0 {
		return c.Render(http.StatusUnprocessableEntity, "upload.upload-form", UploadFormData{Title: title, Error: "Missing title"})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Render(http.StatusUnprocessableEntity, "upload.upload-form", UploadFormData{Title: title, Error: "Missing image"})
	}

	if file.Size > h.maxUploadSizeBytes {
		return c.Render(http.StatusUnprocessableEntity, "upload.upload-form", UploadFormData{Title: title, Error: fmt.Sprintf("File size must be less than %v bytes", h.maxUploadSizeBytes)})
	}

	src, err := file.Open()
	if err != nil {
		return c.Render(http.StatusUnprocessableEntity, "upload.upload-form", UploadFormData{Title: title, Error: "Invalid image uploaded"})
	}
	defer src.Close()

	mimeType, err := validateMimeType(src)
	if err != nil {
		return c.Render(http.StatusUnprocessableEntity, "upload.upload-form", UploadFormData{Title: title, Error: err.Error()})
	}

	p := model.NewPost(1, title)
	p, err = h.PostService.Save(p, &src, mimeType)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "upload.upload-form", UploadFormData{Title: title, Error: "Error saving new post"})
	}

	c.Response().Header().Set("HX-Redirect", p.GetPostUrl())
	return c.String(http.StatusNoContent, "")
}

func validateMimeType(f multipart.File) (string, error) {
	contents, err := io.ReadAll(f)
	defer f.Seek(0, io.SeekStart)
	if err != nil {
		return "", errors.New("invalid image uploaded")
	}

	mimeType := http.DetectContentType(contents)
	for _, allowedMimeType := range allowedImageTypes {
		if allowedMimeType == mimeType {
			return allowedMimeType, nil
		}
	}

	return "", errors.New("invalid file type. Must be one of: " + strings.Join(allowedImageTypes, ", "))
}
