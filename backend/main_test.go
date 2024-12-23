package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World from Docker compose v2!\n", rec.Body.String())
	}
}

func TestHelloWithName(t *testing.T) {
	e := setup()

	targets := []struct {
		name string
		code int
		body string
	}{
		{"Alice", http.StatusOK, "Hello, Alice!\n"},
		{"Bob", http.StatusOK, "Hello, Bob!\n"},
	}

	for _, target := range targets {
		req := httptest.NewRequest(http.MethodGet, "/hello/"+target.name, nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		c.SetParamNames("name")
		c.SetParamValues(target.name)

		if assert.NoError(t, helloByName(c)) {
			assert.Equal(t, target.code, rec.Code)
			assert.Equal(t, target.body, rec.Body.String())
		}
	}
}


func TestCreateSession(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodGet, "/create-session", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte(SECRET_KEY)))

	if assert.NoError(t, createSession(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestReadSessionWithoutSession(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodGet, "/sess/read-session", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte(SECRET_KEY)))

	if assert.NoError(t, readSession(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, "invalid session", rec.Body.String())
	}
}