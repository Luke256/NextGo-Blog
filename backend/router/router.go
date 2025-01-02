package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"nextgoBlog/repository"
	"nextgoBlog/router/auth"
	"nextgoBlog/router/session"
	"nextgoBlog/router/v1"
)

type Router struct {
	e *echo.Echo
	v1 *v1.Handler
	p *auth.Provider
}

func Setup(e *echo.Echo, db *gorm.DB, repo repository.Repository) *echo.Echo {
	r := NewRouter(e, db, repo)

	api := r.e.Group("/api")
	r.v1.Setup(api)
	r.p.Setup(api)

	return e
}

func NewRouter(e *echo.Echo, db *gorm.DB, repo repository.Repository) *Router {
	handler := &v1.Handler {
		Repo: repo,
	}
	store := session.NewSessionStore(db)
	p := auth.NewProvider(store)

	r := &Router {
		e: e,
		v1: handler,
		p: p,
	}

	api := r.e.Group("/api")
	r.v1.Setup(api)
	r.p.Setup(api)

	return r
}