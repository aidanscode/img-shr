package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AidansCode/img-shr/app/service"
	"github.com/labstack/echo/v4"
)

type ViewData struct {
	Title string
	DirectUrl string
}

func (h *Handler) View(c echo.Context) error {
	id, err := getPostId(c)
	if err != nil {
		return h.renderError(c, http.StatusNotFound, err.Error())
	}

	post, err := h.PostService.Get(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return h.renderError(c, http.StatusNotFound, "Post not found")
		}
		return h.renderError(c, http.StatusInternalServerError, "Error occurred while finding image")
	}

	return c.Render(http.StatusOK, "image", ViewData{Title: post.Title, DirectUrl: post.GetDirectUrl()})
}

func (h *Handler) ViewRaw(c echo.Context) error {
	id, err := getPostId(c)
	if err != nil {
		return h.renderError(c, http.StatusNotFound, err.Error())
	}

	post, err := h.PostService.Get(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return h.renderError(c, http.StatusNotFound, "Post not found")
		}
		return h.renderError(c, http.StatusInternalServerError, "Error occurred while finding image")
	}

	return c.File(post.ImgPath)
}

func getPostId(c echo.Context) (int, error) {
	rawId := c.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return 0, errors.New("invalid id given")
	}

	return id, nil
}
