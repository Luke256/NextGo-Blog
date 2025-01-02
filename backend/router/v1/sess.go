package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)



func (h *Handler) ReadSession(c echo.Context) error {
	s, err := h.ss.GetSession(c)
	if err != nil {
		return err
	}

	d, err := s.Get("foo")
	if err != nil {
		return err
	}

	res := ReadSessionResponse{ 
		Foo: d.(string),
	};

	return c.JSON(http.StatusOK, res)
}