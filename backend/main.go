package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func main() {
	// Echoの新しいインスタンスを作成
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", hello)
	e.GET("/hello/:name", helloByName)

	e.POST("/login", login)

	private := e.Group("/private")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	private.Use(echojwt.WithConfig(config))

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World from Docker compose v2!\n")
}

func helloByName(c echo.Context) error {
	name := c.Param("name")

	return c.String(http.StatusOK, "Hello, "+name+"!\n")
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "user" && password == "password" {
		return c.String(http.StatusOK, "Login success!\n")
	} else {
		return c.String(http.StatusUnauthorized, "Login failed!\n")
	}
}
