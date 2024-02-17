package main

import (
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		component := hello("Johnny")
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/other", func(c echo.Context) error {
		component := other()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
