// test for main package

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// Echoの新しいインスタンスを作成
	e := echo.New()

	// テスト用のリクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	// テスト用のレスポンスを作成
	rec := httptest.NewRecorder()

	// テスト用のリクエストとレスポンスを紐付ける
	c := e.NewContext(req, rec)

	// テスト対象の関数を実行
	Hello(c)

	// レスポンスのステータスコードが200番であることを検証
	assert.Equal(t, http.StatusOK, rec.Code)

	// レスポンスのボディが「Hello, World.」であることを検証
	assert.Equal(t, "Hello, World from Docker compose v2!\n", rec.Body.String())
}