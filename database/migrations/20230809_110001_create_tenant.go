package migrations

import (
	"github.com/atom-apps/auth/common"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func (m *Migration20230809_110001CreateTenant) table() interface{} {
	type Tenant struct {
		gorm.Model
		Name        string `gorm:"size:64"`
		Description string `gorm:"size:256"`
		Meta        datatypes.JSONType[common.TenantMeta]
	}

	return &Tenant{}
}

func (m *Migration20230809_110001CreateTenant) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_110001CreateTenant) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
	// return tx.Migrator().DropColumn(m.table(), "input_column_name")
}

// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
func init() {
	Migrations = append(Migrations, New20230809_110001CreateTenantMigration)
}

type Migration20230809_110001CreateTenant struct {
	id string
}

func New20230809_110001CreateTenantMigration() contracts.Migration {
	return &Migration20230809_110001CreateTenant{id: "20230809_110001_create_tenant"}
}

func (m *Migration20230809_110001CreateTenant) ID() string {
	return m.id
}
