package v1

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Hello(c echo.Context) error {

	res := h.Repo.Hello()

	return c.JSON(200, res)
}

func (h *Handler) HelloName(c echo.Context) error {
	name := c.Param("name")

	res := h.Repo.HelloName(name)

	return c.JSON(200, res)
}
