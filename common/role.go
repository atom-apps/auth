package common

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/samber/lo"
)

const (
	RoleAdmin       = "Admin"
	RoleTenantAdmin = "TenantAdmin"
)

func AnyAdmin(claim *casdoorsdk.Claims) bool {
	return claim.IsAdmin || claim.IsGlobalAdmin || IsAdmin(claim.Roles) || IsTenantAdmin(claim.Roles)
}

func IsAdmin(roles []*casdoorsdk.Role) bool {
	names := lo.Map(roles, func(role *casdoorsdk.Role, _ int) string {
		return role.Name
	})

	return lo.Contains(names, RoleAdmin)
}

func IsTenantAdmin(roles []*casdoorsdk.Role) bool {
	names := lo.Map(roles, func(role *casdoorsdk.Role, _ int) string {
		return role.Name
	})

	return lo.Contains(names, RoleTenantAdmin)
}
