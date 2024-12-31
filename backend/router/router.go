package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"nextgoBlog/repository"
	"nextgoBlog/router/v1"
)

type Router struct {
	e *echo.Echo
	v1 *v1.Handler
}

func Setup(e *echo.Echo, db *gorm.DB, repo repository.Repository) *echo.Echo {
	handler := &v1.Handler {
		Repo: repo,
	}
	r := &Router {
		e: e,
		v1: handler,
	}

	api := r.e.Group("/api")
	r.v1.Setup(api)

	return e
}