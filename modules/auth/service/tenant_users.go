package service

import (
	"context"

	"github.com/atom-apps/auth/common"
	"github.com/atom-apps/auth/database/models"
	"github.com/atom-apps/auth/modules/auth/dao"
	"github.com/atom-apps/auth/modules/auth/dto"

	"github.com/jinzhu/copier"
)

// @provider
type TenantUserService struct {
	tenantUserDao *dao.TenantUserDao
}

func (svc *TenantUserService) DecorateItem(model *models.TenantUser, id int) *dto.TenantUserItem {
	var dtoItem *dto.TenantUserItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *TenantUserService) GetByID(ctx context.Context, id int64) (*models.TenantUser, error) {
	return svc.tenantUserDao.GetByID(ctx, id)
}

func (svc *TenantUserService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantUserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.TenantUser, error) {
	return svc.tenantUserDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *TenantUserService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantUserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.TenantUser, int64, error) {
	return svc.tenantUserDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *TenantUserService) CreateFromModel(ctx context.Context, model *models.TenantUser) error {
	return svc.tenantUserDao.Create(ctx, model)
}

// Create
func (svc *TenantUserService) Create(ctx context.Context, body *dto.TenantUserForm) error {
	model := &models.TenantUser{}
	_ = copier.Copy(model, body)
	return svc.tenantUserDao.Create(ctx, model)
}

// Update
func (svc *TenantUserService) Update(ctx context.Context, id int64, body *dto.TenantUserForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.tenantUserDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *TenantUserService) UpdateFromModel(ctx context.Context, model *models.TenantUser) error {
	return svc.tenantUserDao.Update(ctx, model)
}

// Delete
func (svc *TenantUserService) Delete(ctx context.Context, id int64) error {
	return svc.tenantUserDao.Delete(ctx, id)
}

// GetByUserID
func (svc *TenantUserService) GetByUserID(ctx context.Context, userID int64) (*models.TenantUser, error) {
	return svc.tenantUserDao.GetByUserID(ctx, userID)
}
