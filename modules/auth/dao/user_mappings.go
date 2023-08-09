package dao

import (
	"context"

	"github.com/atom-apps/auth/common"
	"github.com/atom-apps/auth/database/models"
	"github.com/atom-apps/auth/database/query"
	"github.com/atom-apps/auth/modules/auth/dto"

	"gorm.io/gen/field"
)

// @provider
type UserMappingDao struct {
	query *query.Query
}

func (dao *UserMappingDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *UserMappingDao) Context(ctx context.Context) query.IUserMappingDo {
	return dao.query.UserMapping.WithContext(ctx)
}

func (dao *UserMappingDao) decorateSortQueryFilter(query query.IUserMappingDo, sortFilter *common.SortQueryFilter) query.IUserMappingDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.UserMapping.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.UserMapping.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *UserMappingDao) decorateQueryFilter(query query.IUserMappingDo, queryFilter *dto.UserMappingListQueryFilter) query.IUserMappingDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.UUID != nil {
		query = query.Where(dao.query.UserMapping.UUID.Eq(*queryFilter.UUID))
	}

	return query
}

func (dao *UserMappingDao) UpdateColumn(ctx context.Context, id int64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.UserMapping.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *UserMappingDao) Update(ctx context.Context, model *models.UserMapping) error {
	_, err := dao.Context(ctx).Where(dao.query.UserMapping.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *UserMappingDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Where(dao.query.UserMapping.ID.Eq(id)).Delete()
	return err
}

func (dao *UserMappingDao) DeletePermanently(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.UserMapping.ID.Eq(id)).Delete()
	return err
}

func (dao *UserMappingDao) Restore(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.UserMapping.ID.Eq(id)).UpdateSimple(dao.query.UserMapping.DeletedAt.Null())
	return err
}

func (dao *UserMappingDao) Create(ctx context.Context, model *models.UserMapping) error {
	return dao.Context(ctx).Create(model)
}

func (dao *UserMappingDao) GetByID(ctx context.Context, id int64) (*models.UserMapping, error) {
	return dao.Context(ctx).Where(dao.query.UserMapping.ID.Eq(id)).First()
}

func (dao *UserMappingDao) GetByName(ctx context.Context, name string) (*models.UserMapping, error) {
	return dao.Context(ctx).Where(dao.query.UserMapping.Name.Eq(name)).First()
}

func (dao *UserMappingDao) GetByIDs(ctx context.Context, ids []int64) ([]*models.UserMapping, error) {
	return dao.Context(ctx).Where(dao.query.UserMapping.ID.In(ids...)).Find()
}

func (dao *UserMappingDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserMappingListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.UserMapping, int64, error) {
	query := dao.query.UserMapping
	userMappingQuery := query.WithContext(ctx)
	userMappingQuery = dao.decorateQueryFilter(userMappingQuery, queryFilter)
	userMappingQuery = dao.decorateSortQueryFilter(userMappingQuery, sortFilter)
	return userMappingQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *UserMappingDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserMappingListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.UserMapping, error) {
	query := dao.query.UserMapping
	userMappingQuery := query.WithContext(ctx)
	userMappingQuery = dao.decorateQueryFilter(userMappingQuery, queryFilter)
	userMappingQuery = dao.decorateSortQueryFilter(userMappingQuery, sortFilter)
	return userMappingQuery.Find()
}

func (dao *UserMappingDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserMappingListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.UserMapping, error) {
	query := dao.query.UserMapping
	userMappingQuery := query.WithContext(ctx)
	userMappingQuery = dao.decorateQueryFilter(userMappingQuery, queryFilter)
	userMappingQuery = dao.decorateSortQueryFilter(userMappingQuery, sortFilter)
	return userMappingQuery.First()
}
