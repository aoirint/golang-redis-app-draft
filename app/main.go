// https://mem-archive.com/2020/02/24/post-2396/
package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", hello)

  e.Logger.Fatal(e.Start(":8000"))
}

func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
