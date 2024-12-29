package main

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
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
		req := httptest.NewRequest(http.MethodGet, "/api/hello/"+target.name, nil)
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
	
	server := httptest.NewServer(e)
	defer server.Close()

	r, err := http.Get(server.URL + "/api/create-session")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, r.StatusCode)
}

func TestReadSessionWithoutSession(t *testing.T) {
	e := setup()
	server := httptest.NewServer(e)
	defer server.Close()

	r, err := http.Get(server.URL + "/api/user/read-session")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusUnauthorized, r.StatusCode)
}

func TestReadSessionWithSession(t *testing.T) {
	e := setup()
	server := httptest.NewServer(e)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	c := http.Client{Jar: jar}
	r, err := c.Get(server.URL + "/api/create-session")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, r.StatusCode)

	r, err = c.Get(server.URL + "/api/user/read-session")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Equal(t, "foo=bar\n", readBody(r))
}

func readBody(r *http.Response) string {
	buf := make([]byte, 1024)
	n, _ := r.Body.Read(buf)
	return string(buf[:n])
}