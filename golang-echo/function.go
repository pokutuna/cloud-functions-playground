package function

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo
var count = 0

func init() {
	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", top)
	e.GET("/counter", counter)
	e.GET("/*", path)
}

func top(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func counter(c echo.Context) error {
	count++
	return c.String(http.StatusOK, fmt.Sprintf("〜 あなたは %d 人目の function への訪問者です ~", count))
}

func path(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("GET %s", c.Request().URL.EscapedPath()))
}

// App is exposed as a function
func App(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}
