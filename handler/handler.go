package handler

import (
	"net/http"
	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello World " + username)
	}
}

func JsonPage() echo.HandlerFunc {
	jsonMap := map[string]string {
		"foo": "bar",
		"hoge": "fuga",
	}
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, jsonMap)
	}
}
