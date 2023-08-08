package controller

import (
	"errors"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
)

type Common struct{}

func (c *Common) Claim(ctx *fiber.Ctx, door *casdoorsdk.Client) (*casdoorsdk.Claims, string, error) {
	token := string(ctx.Request().Header.Peek("Authorization"))
	if token == "" {
		return nil, token, errors.New("Invalid token")
	}
	token = token[len("Bearer "):]

	claim, err := door.ParseJwtToken(token)
	if err != nil {
		return nil, "", err
	}

	return claim, token, nil
}
