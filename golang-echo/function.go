package function

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

func init() {
	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", top)
	e.GET("/a", path)
	e.GET("/b", path)
	e.GET("/*", path)
}

func top(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func path(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("GET %s", c.Request().URL.EscapedPath()))
}

// App is exposed as a function
func App(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}
