package controller

import (
	"github.com/atom-apps/auth/common"
	"github.com/atom-apps/auth/modules/auth/dto"
	"github.com/atom-apps/auth/modules/auth/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type TenantController struct {
	tenantSvc *service.TenantService
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"TenantID"
//	@Success		200	{object}	dto.TenantItem
//	@Router			/tenants/{id} [get]
func (c *TenantController) Show(ctx *fiber.Ctx, id int64) (*dto.TenantItem, error) {
	item, err := c.tenantSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.tenantSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.TenantListQueryFilter	true	"TenantListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.TenantItem}
//	@Router			/tenants [get]
func (c *TenantController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.TenantListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.tenantSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.tenantSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.TenantForm	true	"TenantForm"
//	@Success		200		{string}	TenantID
//	@Router			/tenants [post]
func (c *TenantController) Create(ctx *fiber.Ctx, body *dto.TenantForm) error {
	return c.tenantSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"TenantID"
//	@Param			body	body		dto.TenantForm	true	"TenantForm"
//	@Success		200		{string}	TenantID
//	@Failure		500		{string}	TenantID
//	@Router			/tenants/{id} [put]
func (c *TenantController) Update(ctx *fiber.Ctx, id int64, body *dto.TenantForm) error {
	return c.tenantSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"TenantID"
//	@Success		200	{string}	TenantID
//	@Failure		500	{string}	TenantID
//	@Router			/tenants/{id} [delete]
func (c *TenantController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.tenantSvc.Delete(ctx.Context(), id)
}
