package lib

import (
	"strings"

	"github.com/labstack/echo/v4"
	"go.stellar.af/go-utils/slice"
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
