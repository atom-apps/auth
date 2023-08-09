package dao

import (
	"github.com/atom-apps/auth/database/query"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		query *query.Query,
	) (*TenantUserDao, error) {
		obj := &TenantUserDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*TenantDao, error) {
		obj := &TenantDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*UserMappingDao, error) {
		obj := &UserMappingDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
