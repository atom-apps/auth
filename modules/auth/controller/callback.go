package controller

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// @provider
type CallbackController struct {
	door *casdoorsdk.Client
}

// Signin
//
//	@Summary		signin callback
//	@Description	signin callback
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	oauth2.Token
//	@Router			/auth/callback/signin [get]
func (c *CallbackController) Signin(ctx *fiber.Ctx) (*oauth2.Token, error) {
	// TODO: Write to local storage
	return c.door.GetOAuthToken(ctx.Query("code"), ctx.Query("state"))
}

// Signup
//
//	@Summary		signup
//	@Description	signup
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		301	{string}	string	""
//	@Router			/auth/callback/signup [get]
func (c *CallbackController) Signup(ctx *fiber.Ctx) (*oauth2.Token, error) {
	return c.Signin(ctx)
}
