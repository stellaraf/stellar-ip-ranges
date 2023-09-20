package lib

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/stellaraf/stellar-ip-ranges/constants"
)

var ErrUnknownTypeOrScope = errors.New("unknown type and/or scope")

func MatchText(scope, _type string) ([]byte, error) {
	for _, c := range constants.TextConditions {
		if c.Match(scope, _type) {
			return c.Data.Text(), nil
		}
	}
	return nil, ErrUnknownTypeOrScope
}

func MatchJSON(scope, _type string) (fiber.Map, error) {
	for _, c := range constants.JSONConditions {
		if c.Match(scope, _type) {
			return c.Response, nil
		}
	}
	return nil, ErrUnknownTypeOrScope
}

func TextHandler(ctx *fiber.Ctx) error {
	scope := ctx.Params("scope", "global")
	_type := ctx.Params("type", "dual")
	ctx.Set("content-type", "text/plain")
	ctx.Status(200)

	res, err := MatchText(scope, _type)
	if err != nil {
		return ctx.Status(400).Send([]byte(err.Error()))
	}
	return ctx.Send(res)

}

func JSONHandler(ctx *fiber.Ctx) error {
	scope := ctx.Params("scope", "global")
	_type := ctx.Params("type", "dual")
	ctx.Status(200)
	res, err := MatchJSON(scope, _type)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}
