package controller

import (
	"errors"

	"github.com/atom-apps/auth/common"
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

// ForceResetPassword
//
//	@Summary		ForceResetPassword
//	@Description	ForceResetPassword
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	""
//	@Router			/auth/user/force-reset-password/{uuid} [put]
func (c *UserController) ForceResetPassword(ctx *fiber.Ctx, uuid string, body *dto.ResetPasswordForm) error {
	claim, _, err := c.Claim(ctx, c.door)
	if err != nil {
		return err
	}

	if !common.AnyAdmin(claim) {
		return errors.New("invalid user")
	}

	u, err := c.door.GetUserByUserId(uuid)
	if err != nil {
		return err
	}
	ok, err := c.door.SetPassword(u.Owner, u.Name, "", body.New)
	if !ok {
		return errors.New("reset password failed")
	}
	return err
}

// Update user info
//
//	@Summary		update user info
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	""
//	@Router			/auth/user/{uuid} [put]
func (c *UserController) Update(ctx *fiber.Ctx, uuid string, user *dto.UserForm) error {
	claim, _, err := c.Claim(ctx, c.door)
	if err != nil {
		return err
	}

	if claim.ID != uuid || !common.AnyAdmin(claim) {
		return errors.New("invalid user")
	}

	u, err := c.door.GetUserByUserId(uuid)
	if err != nil {
		return err
	}

	u.DisplayName = user.DisplayName

	ok, err := c.door.UpdateUser(u)
	if !ok {
		return errors.New("update user failed")
	}
	return err
}
