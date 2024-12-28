package main

import (
	"net/http"
	"fmt"
	
    "github.com/gorilla/sessions"
    "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	SECRET_KEY = "SECRET_KEY"
)

func main() {
	// Echoの新しいインスタンスを作成
	e := setup()


	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}

func setup() *echo.Echo{
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{
			echo.HeaderOrigin, 
			echo.HeaderContentType, 
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowCredentials,
		},
		AllowCredentials: true,
	}))
	
	api := e.Group("/api")
	api.GET("/hello", hello)
	api.GET("/hello/:name", helloByName)
	api.GET("/create-session",createSession)
	
	sess := api.Group("/sess")
	sess.Use(readSessionMiddleware)
	sess.GET("/read-session", readSession)
	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World from Docker compose v2!\n")
}

func helloByName(c echo.Context) error {
	name := c.Param("name")

	return c.String(http.StatusOK, "Hello, "+name+"!\n")
}

func createSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   10000000,
		HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	sess.Values["name"] = "Luke"
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func readSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("foo=%v\n", sess.Values["foo"]))
}

func readSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		if sess.Values["foo"] == nil {
			return c.String(http.StatusUnauthorized, "invalid session")
		}
		return next(c)
	}
}