package service

import (
	"github.com/atom-apps/auth/modules/auth/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		tenantUserDao *dao.TenantUserDao,
	) (*TenantUserService, error) {
		obj := &TenantUserService{
			tenantUserDao: tenantUserDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		tenantDao *dao.TenantDao,
	) (*TenantService, error) {
		obj := &TenantService{
			tenantDao: tenantDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		userMappingDao *dao.UserMappingDao,
	) (*UserMappingService, error) {
		obj := &UserMappingService{
			userMappingDao: userMappingDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
