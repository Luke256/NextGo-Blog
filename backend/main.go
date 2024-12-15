package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echoの新しいインスタンスを作成
	e := echo.New()

	// 「/hello」というエンドポイントを設定する
	e.GET("/hello", Hello)

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World from Docker compose v2!\n")
}