package lib

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"go.stellar.af/go-utils/slice"
	"go.stellar.af/stellar-ip-ranges/constants"
)

func BaseHandler(ctx echo.Context) error {
	ctx.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	params := strings.Split(strings.Join(ctx.ParamValues(), "/"), "/")
	if slice.Contains(params, "json") {
		return JSONHandler(ctx)
	}
	return TextHandler(ctx)
}

func TextHandler(ctx echo.Context) error {
	params := strings.Split(strings.Join(ctx.ParamValues(), "/"), "/")
	res := MatchText(params)
	if res == "" {
		return ctx.String(400, "no matching parameters")
	}
	return ctx.String(200, res)
}

func JSONHandler(ctx echo.Context) error {
	params := strings.Split(strings.Join(ctx.ParamValues(), "/"), "/")
	res := MatchJSON(params)
	if res == nil {
		return ctx.JSON(400, map[string]string{"error": "no matching parameters"})
	}
	return ctx.JSON(200, res)
}

func GeofeedHandler(ctx echo.Context) error {
	title := "AS14525 Geofeed"
	sha := os.Getenv("VERCEL_GIT_COMMIT_SHA")
	if sha != "" {
		title += " " + sha
	}
	title += "\n# Append .txt or .csv to download as a file in the respective format."
	csv := constants.GEOFEED.CSV(title)
	path := ctx.Request().URL.Path
	if strings.HasSuffix(path, ".csv") {
		ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=stellar-geofeed.csv")
		return ctx.Blob(200, "text/csv", []byte(csv))
	}
	if strings.HasSuffix(path, ".txt") {
		ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=stellar-geofeed.txt")
		return ctx.Blob(200, "text/plain", []byte(csv))
	}
	return ctx.String(200, csv)
}
