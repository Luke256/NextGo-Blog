package auth

import (
	"net/http"
	"nextgoBlog/router/session"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	cookieName   = "session"
	cookieMaxAge = 60 * 60 * 24 * 7
)

func SessionIssuer(ss session.Store) echo.HandlerFunc {
	return func(c echo.Context) error {

		data := map[string]interface{}{
			"foo": "bar",
		}

		s, err := ss.IssueSession("Luke", data)
		if err != nil {
			return err
		}

		token := s.Token()

		c.SetCookie(&http.Cookie{
			Name:     "session",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(cookieMaxAge * time.Second),
			MaxAge:   cookieMaxAge,
			HttpOnly: true,
		})

		return c.NoContent(http.StatusOK)
	}
}
