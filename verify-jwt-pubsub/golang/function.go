package function

import (
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
}

func top(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// App is exposed as a function
func App(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}
