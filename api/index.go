package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stellaraf/stellar-ip-ranges/lib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	app := echo.New()

	app.Use(middleware.Recover())
	app.Use(middleware.Gzip())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	app.GET("/*", lib.BaseHandler)

	app.ServeHTTP(w, r)
}
