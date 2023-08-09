package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_100328CreateUserMapping) table() interface{} {
	type UserMapping struct {
		gorm.Model
		Uuid string `gorm:"size:128"`
		Name string `gorm:"size:128"`
	}

	return &UserMapping{}
}

func (m *Migration20230809_100328CreateUserMapping) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_100328CreateUserMapping) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_100328CreateUserMappingMigration)
}

type Migration20230809_100328CreateUserMapping struct {
	id string
}

func New20230809_100328CreateUserMappingMigration() contracts.Migration {
	return &Migration20230809_100328CreateUserMapping{id: "20230809_100328_create_user_mapping"}
}

func (m *Migration20230809_100328CreateUserMapping) ID() string {
	return m.id
}
