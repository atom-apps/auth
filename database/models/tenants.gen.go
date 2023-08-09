// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"github.com/atom-apps/auth/common"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const TableNameTenant = "tenants"

// Tenant mapped from table <tenants>
type Tenant struct {
	ID          int64                                 `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time                             `gorm:"column:created_at;type:timestamp with time zone" json:"created_at"`
	UpdatedAt   time.Time                             `gorm:"column:updated_at;type:timestamp with time zone" json:"updated_at"`
	DeletedAt   gorm.DeletedAt                        `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
	Name        string                                `gorm:"column:name;type:character varying(64)" json:"name"`
	Description string                                `gorm:"column:description;type:character varying(256)" json:"description"`
	Meta        datatypes.JSONType[common.TenantMeta] `gorm:"column:meta;type:jsonb" json:"meta"`
}

func (*Tenant) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameTenant
	}
	return namer.TableName(TableNameTenant)
}