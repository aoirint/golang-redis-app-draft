// https://mem-archive.com/2020/02/24/post-2396/
package main

import (
  "time"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "context"
  "github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
  Addr: "redis:6379",
  Password: "",
  DB: 0,
})

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", getval)
  e.GET("/set", setval)

  e.Logger.Fatal(e.Start(":8000"))
}

func getval(c echo.Context) error {
  val, err := rdb.Get(ctx, "mykey").Result()
  if err == redis.Nil {
    return c.String(http.StatusOK, "no data (nil)")
  } else if err != nil {
      panic(err)
  }

  return c.String(http.StatusOK, val)
}

func setval(c echo.Context) error {
  err := rdb.Set(ctx, "mykey", "hoge", 5 * time.Second).Err()
  if err != nil {
      panic(err)
  }

  return c.String(http.StatusOK, "OK")
}
