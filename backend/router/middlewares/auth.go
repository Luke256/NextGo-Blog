package middleware

import (
	"net/http"

	"nextgoBlog/repository"
	"nextgoBlog/router/consts"
	"nextgoBlog/router/session"

	"github.com/labstack/echo/v4"
)

func Auth(repo repository.Repository, ss session.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := ss.GetSession(c)
			if err != nil {
				return err
			}

			if sess == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "you are not logged in")
			}

			c.Set(consts.KeyUserID, sess.UserID)

			return next(c)
		}
	}
}