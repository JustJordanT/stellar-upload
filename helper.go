package main

import "github.com/labstack/echo/v4"

func KeyAuth(c echo.Context) string {
	return c.Request().Header.Get("KEY")
}

func SecretKeyAuth(c echo.Context) string {
	return c.Request().Header.Get("SECRET_KEY")
}
