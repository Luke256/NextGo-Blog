package v1

import (
	"nextgoBlog/repository"
	"nextgoBlog/router/session"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Repo repository.Repository
	ss session.Store
}

func (h *Handler) Setup(e *echo.Group) {
	e.GET("/hello", h.Hello)
	e.GET("/hello/:name", h.HelloName)
}