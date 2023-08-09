package service

import (
	"context"

	"github.com/atom-apps/auth/database/models"
	"github.com/atom-apps/auth/modules/auth/dao"
	"github.com/atom-apps/auth/modules/auth/dto"

	"github.com/jinzhu/copier"
)

// @provider
type UserMappingService struct {
	userMappingDao *dao.UserMappingDao
}

func (svc *UserMappingService) GetByID(ctx context.Context, id int64) (*models.UserMapping, error) {
	return svc.userMappingDao.GetByID(ctx, id)
}

func (svc *UserMappingService) GetByName(ctx context.Context, name string) (*models.UserMapping, error) {
	return svc.userMappingDao.GetByName(ctx, name)
}

// CreateFromModel
func (svc *UserMappingService) CreateFromModel(ctx context.Context, model *models.UserMapping) error {
	return svc.userMappingDao.Create(ctx, model)
}

// Create
func (svc *UserMappingService) Create(ctx context.Context, body *dto.UserMappingForm) error {
	model := &models.UserMapping{}
	_ = copier.Copy(model, body)
	return svc.userMappingDao.Create(ctx, model)
}

// Update
func (svc *UserMappingService) Update(ctx context.Context, id int64, body *dto.UserMappingForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.userMappingDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *UserMappingService) UpdateFromModel(ctx context.Context, model *models.UserMapping) error {
	return svc.userMappingDao.Update(ctx, model)
}

// Delete
func (svc *UserMappingService) Delete(ctx context.Context, id int64) error {
	return svc.userMappingDao.Delete(ctx, id)
}
