package v1

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Hello(c echo.Context) error {

	msg := h.Repo.Hello()

	res := Message {
		Message: msg,
	}

	return c.JSON(200, res)
}

func (h *Handler) HelloName(c echo.Context) error {
	name := c.Param("name")

	msg := h.Repo.HelloName(name)

	res := Message {
		Message: msg,
	}

	return c.JSON(200, res)
}
