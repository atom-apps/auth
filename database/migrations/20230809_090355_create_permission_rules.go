package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_090355CreatePermissionRules) table() interface{} {
	type PermissionRule struct { // arbitrary number of variable fields
		ID        uint   `gorm:"primaryKey;autoIncrement"`
		Ptype     string `gorm:"size:16"`
		V0        string `gorm:"size:128"`
		V1        string `gorm:"size:128"`
		V2        string `gorm:"size:256"`
		V3        string `gorm:"size:256"`
		V4        string `gorm:"size:256"`
		V5        string `gorm:"size:256"`
		DeletedAt gorm.DeletedAt
	}

	return &PermissionRule{}
}

func (m *Migration20230809_090355CreatePermissionRules) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_090355CreatePermissionRules) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_090355CreatePermissionRulesMigration)
}

type Migration20230809_090355CreatePermissionRules struct {
	id string
}

func New20230809_090355CreatePermissionRulesMigration() contracts.Migration {
	return &Migration20230809_090355CreatePermissionRules{id: "20230809_090355_create_permission_rules"}
}

func (m *Migration20230809_090355CreatePermissionRules) ID() string {
	return m.id
}
