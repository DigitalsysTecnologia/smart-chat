package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202307141048_chat",
		Migrate: func(tx *gorm.DB) error {
			sql := `CREATE TABLE "Chat" (
    					"ID" SERIAL PRIMARY KEY,
    					"CreatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    					"UpdatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					);`

			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("Chat")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
