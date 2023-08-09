// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/atom-apps/auth/database/models"
)

func newTenantUser(db *gorm.DB, opts ...gen.DOOption) tenantUser {
	_tenantUser := tenantUser{}

	_tenantUser.tenantUserDo.UseDB(db, opts...)
	_tenantUser.tenantUserDo.UseModel(&models.TenantUser{})

	tableName := _tenantUser.tenantUserDo.TableName()
	_tenantUser.ALL = field.NewAsterisk(tableName)
	_tenantUser.ID = field.NewInt64(tableName, "id")
	_tenantUser.CreatedAt = field.NewTime(tableName, "created_at")
	_tenantUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tenantUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_tenantUser.TenantID = field.NewInt64(tableName, "tenant_id")
	_tenantUser.UserID = field.NewInt64(tableName, "user_id")
	_tenantUser.IsAdmin = field.NewBool(tableName, "is_admin")

	_tenantUser.fillFieldMap()

	return _tenantUser
}

type tenantUser struct {
	tenantUserDo tenantUserDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	TenantID  field.Int64
	UserID    field.Int64
	IsAdmin   field.Bool

	fieldMap map[string]field.Expr
}

func (t tenantUser) Table(newTableName string) *tenantUser {
	t.tenantUserDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tenantUser) As(alias string) *tenantUser {
	t.tenantUserDo.DO = *(t.tenantUserDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tenantUser) updateTableName(table string) *tenantUser {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.TenantID = field.NewInt64(table, "tenant_id")
	t.UserID = field.NewInt64(table, "user_id")
	t.IsAdmin = field.NewBool(table, "is_admin")

	t.fillFieldMap()

	return t
}

func (t *tenantUser) WithContext(ctx context.Context) ITenantUserDo {
	return t.tenantUserDo.WithContext(ctx)
}

func (t tenantUser) TableName() string { return t.tenantUserDo.TableName() }

func (t tenantUser) Alias() string { return t.tenantUserDo.Alias() }

func (t tenantUser) Columns(cols ...field.Expr) gen.Columns { return t.tenantUserDo.Columns(cols...) }

func (t *tenantUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tenantUser) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["tenant_id"] = t.TenantID
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["is_admin"] = t.IsAdmin
}

func (t tenantUser) clone(db *gorm.DB) tenantUser {
	t.tenantUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tenantUser) replaceDB(db *gorm.DB) tenantUser {
	t.tenantUserDo.ReplaceDB(db)
	return t
}

type tenantUserDo struct{ gen.DO }

type ITenantUserDo interface {
	gen.SubQuery
	Debug() ITenantUserDo
	WithContext(ctx context.Context) ITenantUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITenantUserDo
	WriteDB() ITenantUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITenantUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITenantUserDo
	Not(conds ...gen.Condition) ITenantUserDo
	Or(conds ...gen.Condition) ITenantUserDo
	Select(conds ...field.Expr) ITenantUserDo
	Where(conds ...gen.Condition) ITenantUserDo
	Order(conds ...field.Expr) ITenantUserDo
	Distinct(cols ...field.Expr) ITenantUserDo
	Omit(cols ...field.Expr) ITenantUserDo
	Join(table schema.Tabler, on ...field.Expr) ITenantUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITenantUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITenantUserDo
	Group(cols ...field.Expr) ITenantUserDo
	Having(conds ...gen.Condition) ITenantUserDo
	Limit(limit int) ITenantUserDo
	Offset(offset int) ITenantUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITenantUserDo
	Unscoped() ITenantUserDo
	Create(values ...*models.TenantUser) error
	CreateInBatches(values []*models.TenantUser, batchSize int) error
	Save(values ...*models.TenantUser) error
	First() (*models.TenantUser, error)
	Take() (*models.TenantUser, error)
	Last() (*models.TenantUser, error)
	Find() ([]*models.TenantUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.TenantUser, err error)
	FindInBatches(result *[]*models.TenantUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.TenantUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITenantUserDo
	Assign(attrs ...field.AssignExpr) ITenantUserDo
	Joins(fields ...field.RelationField) ITenantUserDo
	Preload(fields ...field.RelationField) ITenantUserDo
	FirstOrInit() (*models.TenantUser, error)
	FirstOrCreate() (*models.TenantUser, error)
	FindByPage(offset int, limit int) (result []*models.TenantUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITenantUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tenantUserDo) Debug() ITenantUserDo {
	return t.withDO(t.DO.Debug())
}

func (t tenantUserDo) WithContext(ctx context.Context) ITenantUserDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tenantUserDo) ReadDB() ITenantUserDo {
	return t.Clauses(dbresolver.Read)
}

func (t tenantUserDo) WriteDB() ITenantUserDo {
	return t.Clauses(dbresolver.Write)
}

func (t tenantUserDo) Session(config *gorm.Session) ITenantUserDo {
	return t.withDO(t.DO.Session(config))
}

func (t tenantUserDo) Clauses(conds ...clause.Expression) ITenantUserDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tenantUserDo) Returning(value interface{}, columns ...string) ITenantUserDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tenantUserDo) Not(conds ...gen.Condition) ITenantUserDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tenantUserDo) Or(conds ...gen.Condition) ITenantUserDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tenantUserDo) Select(conds ...field.Expr) ITenantUserDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tenantUserDo) Where(conds ...gen.Condition) ITenantUserDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tenantUserDo) Order(conds ...field.Expr) ITenantUserDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tenantUserDo) Distinct(cols ...field.Expr) ITenantUserDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tenantUserDo) Omit(cols ...field.Expr) ITenantUserDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tenantUserDo) Join(table schema.Tabler, on ...field.Expr) ITenantUserDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tenantUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITenantUserDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tenantUserDo) RightJoin(table schema.Tabler, on ...field.Expr) ITenantUserDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tenantUserDo) Group(cols ...field.Expr) ITenantUserDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tenantUserDo) Having(conds ...gen.Condition) ITenantUserDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tenantUserDo) Limit(limit int) ITenantUserDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tenantUserDo) Offset(offset int) ITenantUserDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tenantUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITenantUserDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tenantUserDo) Unscoped() ITenantUserDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tenantUserDo) Create(values ...*models.TenantUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tenantUserDo) CreateInBatches(values []*models.TenantUser, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tenantUserDo) Save(values ...*models.TenantUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tenantUserDo) First() (*models.TenantUser, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.TenantUser), nil
	}
}

func (t tenantUserDo) Take() (*models.TenantUser, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.TenantUser), nil
	}
}

func (t tenantUserDo) Last() (*models.TenantUser, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.TenantUser), nil
	}
}

func (t tenantUserDo) Find() ([]*models.TenantUser, error) {
	result, err := t.DO.Find()
	return result.([]*models.TenantUser), err
}

func (t tenantUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.TenantUser, err error) {
	buf := make([]*models.TenantUser, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tenantUserDo) FindInBatches(result *[]*models.TenantUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tenantUserDo) Attrs(attrs ...field.AssignExpr) ITenantUserDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tenantUserDo) Assign(attrs ...field.AssignExpr) ITenantUserDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tenantUserDo) Joins(fields ...field.RelationField) ITenantUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tenantUserDo) Preload(fields ...field.RelationField) ITenantUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tenantUserDo) FirstOrInit() (*models.TenantUser, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.TenantUser), nil
	}
}

func (t tenantUserDo) FirstOrCreate() (*models.TenantUser, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.TenantUser), nil
	}
}

func (t tenantUserDo) FindByPage(offset int, limit int) (result []*models.TenantUser, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tenantUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tenantUserDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tenantUserDo) Delete(models ...*models.TenantUser) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tenantUserDo) withDO(do gen.Dao) *tenantUserDo {
	t.DO = *do.(*gen.DO)
	return t
}