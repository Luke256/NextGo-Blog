package main

import (
	"net/http"
	"time"

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

const (
	// Key used for JWT_SECRET_KEY generation
	JWT_SECRET_KEY = "SECRET_KEY"
)

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
		SigningKey: []byte(JWT_SECRET_KEY),
	}
	private.Use(echojwt.WithConfig(config))
	private.GET("", privateHello)

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

	if username != "user" || password != "password" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		"Luke Skywalker",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func privateHello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)

	return c.String(http.StatusOK, "Welcome, "+claims.Name+"!\n")
}