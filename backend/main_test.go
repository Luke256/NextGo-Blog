package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nextgoBlog/cmd"
	v1 "nextgoBlog/router/v1"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup() *httptest.Server {
	e := echo.New()
	server := cmd.NewServer(e)
	return httptest.NewServer(server.Router)
}

func TestHello(t *testing.T) {
	server := setup()
	defer server.Close()

	r, err := server.Client().Get(server.URL + "/api/hello")
	if err != nil {
		t.Fatal(err)
	}

	if r.StatusCode != http.StatusOK {
		assert.Equal(t, http.StatusOK, r.StatusCode)
	}

	// expected: json data like {"message":"Hello, World from Docker compose v2!"}
	
	body := make([]byte, 1024)
	n, _ := r.Body.Read(body)
	actual := ByteToMessage(body[:n])

	expct := v1.Message{Message: "Hello, World from Docker compose v2!"}
	
	assert.Equal(t, expct, actual)
}

func TestHelloName(t *testing.T) {
	server := setup()
	defer server.Close()

	targets := []struct {
		name string
		code int
		body v1.Message
	}{
		{"Alice", http.StatusOK, v1.Message{Message: "Hello, Alice!"}},
		{"Bob", http.StatusOK, v1.Message{Message: "Hello, Bob!"}},
	}

	for _, target := range targets {
		r, err := server.Client().Get(server.URL + "/api/hello/" + target.name)
		if err != nil {
			t.Fatal(err)
		}

		if r.StatusCode != target.code {
			assert.Equal(t, target.code, r.StatusCode)
		}

		body := make([]byte, 1024)
		n, _ := r.Body.Read(body)
		actual := ByteToMessage(body[:n])

		assert.Equal(t, target.body, actual)
	}
}

func ByteToMessage(b []byte) v1.Message {
	var m v1.Message
	json.Unmarshal(b, &m)
	return m
}

// import (
// 	"net/http"
// 	"net/http/cookiejar"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestHello(t *testing.T) {
// 	e := setupEcho()

// 	req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, hello(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, "Hello, World from Docker compose v2!\n", rec.Body.String())
// 	}
// }

// func TestHelloWithName(t *testing.T) {
// 	e := setupEcho()

// 	targets := []struct {
// 		name string
// 		code int
// 		body string
// 	}{
// 		{"Alice", http.StatusOK, "Hello, Alice!\n"},
// 		{"Bob", http.StatusOK, "Hello, Bob!\n"},
// 	}

// 	for _, target := range targets {
// 		req := httptest.NewRequest(http.MethodGet, "/api/hello/"+target.name, nil)
// 		rec := httptest.NewRecorder()

// 		c := e.NewContext(req, rec)

// 		c.SetParamNames("name")
// 		c.SetParamValues(target.name)

// 		if assert.NoError(t, helloByName(c)) {
// 			assert.Equal(t, target.code, rec.Code)
// 			assert.Equal(t, target.body, rec.Body.String())
// 		}
// 	}
// }

// func TestCreateSession(t *testing.T) {
// 	e := setupEcho()

// 	server := httptest.NewServer(e)
// 	defer server.Close()

// 	r, err := http.Get(server.URL + "/api/create-session")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, http.StatusOK, r.StatusCode)
// }

// func TestReadSessionWithoutSession(t *testing.T) {
// 	e := setupEcho()
// 	server := httptest.NewServer(e)
// 	defer server.Close()

// 	r, err := http.Get(server.URL + "/api/user/read-session")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, http.StatusUnauthorized, r.StatusCode)
// }

// func TestReadSessionWithSession(t *testing.T) {
// 	e := setupEcho()
// 	server := httptest.NewServer(e)
// 	defer server.Close()

// 	jar, err := cookiejar.New(nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	c := http.Client{Jar: jar}
// 	r, err := c.Get(server.URL + "/api/create-session")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, http.StatusOK, r.StatusCode)

// 	r, err = c.Get(server.URL + "/api/user/read-session")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, http.StatusOK, r.StatusCode)
// 	assert.Equal(t, "foo=bar\n", readBody(r))
// }

// func readBody(r *http.Response) string {
// 	buf := make([]byte, 1024)
// 	n, _ := r.Body.Read(buf)
// 	return string(buf[:n])
// }
