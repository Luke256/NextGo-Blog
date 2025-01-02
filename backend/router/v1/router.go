package v1

import (
	"nextgoBlog/repository"
	middleware "nextgoBlog/router/middlewares"
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

	sess := e.Group("/user", middleware.Auth(h.Repo, h.ss));
	sess.GET("/read-session", h.ReadSession)
}