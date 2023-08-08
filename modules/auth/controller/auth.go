package controller

import (
	"net/url"
	"strings"

	"github.com/atom-apps/auth/modules/auth/dto"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// @provider
type AuthController struct {
	*Common
	door *casdoorsdk.Client
}

func (c *AuthController) buildURL(ctx *fiber.Ctx, path string) string {
	u := url.URL{}
	u.Scheme = string(ctx.Context().URI().Scheme())
	u.Host = ctx.Hostname()
	u.Path = "/" + strings.TrimLeft(path, "/")
	if ctx.Query("redirect") != "" {
		u.Query().Add("redirect", ctx.Query("redirect"))
	}
	return u.String()
}

// Signin
//
//	@Summary		signin
//	@Description	signin
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		301	{string}	string	""
//	@Router			/auth/signin [get]
func (c *AuthController) Signin(ctx *fiber.Ctx) error {
	signinUrl := c.door.GetSigninUrl(c.buildURL(ctx, "/auth/callback/signin"))
	return ctx.Redirect(signinUrl, fiber.StatusMovedPermanently)
}

// Signup
//
//	@Summary		signup
//	@Description	signup
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/signup [get]
func (c *AuthController) Signup(ctx *fiber.Ctx) error {
	signupUrl := c.door.GetSignupUrl(false, c.buildURL(ctx, "/auth/callback/signup"))
	return ctx.Redirect(signupUrl, fiber.StatusMovedPermanently)
}

// RefreshToken
//
//	@Summary		RefreshToken
//	@Description	RefreshToken
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/refresh-token [post]
func (c *AuthController) RefreshToken(ctx *fiber.Ctx, body *dto.TokenForm) (*oauth2.Token, error) {
	_, _, err := c.Claim(ctx, c.door)
	if err != nil {
		return nil, err
	}

	return c.door.RefreshOAuthToken(body.Token)
}

// Logout
//
//	@Summary		logout
//	@Description	logout
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/logout [post]
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	_, token, err := c.Claim(ctx, c.door)
	if err != nil {
		return err
	}

	_, err = c.door.DeleteToken(token)
	return err
}
