package auth

import (
	"nextgoBlog/router/session"

	"github.com/labstack/echo/v4"
)

type Provider struct {
	ss session.Store
}

func NewProvider(ss session.Store) *Provider {
	return &Provider{
		ss: ss,
	}
}

func (p *Provider) IssueSession(c echo.Context) error {
	return SessionIssuer(p.ss)(c)
}

func (p *Provider) Setup(e *echo.Group) {
	e.GET("/create-session", p.IssueSession)
}