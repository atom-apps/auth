package controller

import (
	"errors"
	"fmt"

	"github.com/atom-apps/auth/common"
	"github.com/atom-apps/auth/modules/auth/dto"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type RoleController struct {
	*Common
	door *casdoorsdk.Client
}

// AddUsers
//
//	@Summary		AddUsers
//	@Description	AddUsers
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	""
//	@Router			/auth/role/add-users [put]
func (c *RoleController) AddUsers(ctx *fiber.Ctx, body *dto.RoleUsersForm) error {
	claim, _, err := c.Claim(ctx, c.door)
	if err != nil {
		return err
	}
	if !common.AnyAdmin(claim) {
		return errors.New("invalid user")
	}

	r, err := c.door.GetRole(body.Role)
	if err != nil {
		return err
	}
	oldSize := len(r.Roles)

	r.Roles = append(r.Roles, lo.Map(body.Users, func(user string, _ int) string {
		return fmt.Sprintf("%s/%s", c.door.OrganizationName, user)
	})...)
	r.Roles = lo.Uniq(r.Roles)

	if oldSize == len(r.Roles) {
		return nil
	}

	ok, err := c.door.UpdateRole(r)
	if !ok {
		return errors.New("role add users failed")
	}

	return err
}

// RemoveUsers
//
//	@Summary		RemoveUsers
//	@Description	RemoveUsers
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	""
//	@Router			/auth/role/remove-users [put]
func (c *RoleController) RemoveUsers(ctx *fiber.Ctx, body *dto.RoleUsersForm) error {
	claim, _, err := c.Claim(ctx, c.door)
	if err != nil {
		return err
	}

	if !common.AnyAdmin(claim) {
		return errors.New("invalid user")
	}

	r, err := c.door.GetRole(body.Role)
	if err != nil {
		return err
	}
	oldSize := len(r.Roles)

	removeUsers := lo.Map(body.Users, func(user string, _ int) string {
		return fmt.Sprintf("%s/%s", c.door.OrganizationName, user)
	})

	r.Roles = lo.Filter(r.Roles, func(user string, _ int) bool {
		return !lo.Contains(removeUsers, user)
	})

	if oldSize == len(r.Roles) {
		return nil
	}

	ok, err := c.door.UpdateRole(r)
	if !ok {
		return errors.New("role add users failed")
	}

	return err
}
