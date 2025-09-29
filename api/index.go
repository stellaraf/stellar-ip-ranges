package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.stellar.af/stellar-ip-ranges/lib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	app := echo.New()

	app.Use(middleware.Recover())
	app.Use(middleware.Gzip())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin},
	}))

	app.GET("/rapid7*", lib.Rapid7Handler)

	app.GET("/crowdstrike*", lib.CrowdstrikeHandler)

	app.GET("/geofeed*", lib.GeofeedHandler)

	app.GET("/*", lib.BaseHandler)

	app.ServeHTTP(w, r)
}
