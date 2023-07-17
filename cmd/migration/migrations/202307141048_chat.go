package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"smart-chat/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202307141048_chat",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.Chat{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("Chat")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
