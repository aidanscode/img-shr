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

	p := model.NewPost(1, title, "https://placehold.co/300x200")
	p, err = h.PostService.Save(p)
	if err != nil {
		return c.Render(http.StatusOK, "upload.upload-form", UploadFormData{Title: title, Error: "Error saving new post"})
	}

	return c.Render(http.StatusOK, "upload.upload-form", UploadFormData{Title: title, Error: fmt.Sprintf("Okay! Saved id: %d", p.Id)})
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
