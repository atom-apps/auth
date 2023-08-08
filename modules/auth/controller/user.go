package controller

import (
	"errors"

	"github.com/atom-apps/auth/modules/auth/dto"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
)

// @provider
type UserController struct {
	*Common
	door *casdoorsdk.Client
}

// Profile
//
//	@Summary		Profile
//	@Description	Profile
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	""
//	@Router			/auth/user/profile [get]
func (c *UserController) Profile(ctx *fiber.Ctx) (*casdoorsdk.User, error) {
	claim, _, err := c.Claim(ctx, c.door)
	if err != nil {
		return nil, err
	}
	return c.door.GetUser(claim.Name)
}

// ResetPassword
//
//	@Summary		ResetPassword
//	@Description	ResetPassword
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	""
//	@Router			/auth/user/reset-password [post]
func (c *UserController) ResetPassword(ctx *fiber.Ctx, body *dto.ResetPasswordForm) error {
	claim, _, err := c.Claim(ctx, c.door)
	if err != nil {
		return err
	}

	ok, err := c.door.SetPassword(claim.Owner, claim.Name, body.Old, body.New)
	if !ok {
		return errors.New("reset password failed")
	}

	return err
}
