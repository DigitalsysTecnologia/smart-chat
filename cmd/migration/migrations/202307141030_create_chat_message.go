package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"smart-chat/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202307141030_create_chat_message",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.ChatMessage{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("ChatMessage")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
