package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202307141048_chat",
		Migrate: func(tx *gorm.DB) error {
			sql := `CREATE TABLE IF NOT EXISTS chat(
					id INT AUTO_INCREMENT PRIMARY KEY,
					user_id VARCHAR(191),
					created_at VARCHAR(50)
					);`

			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("chat")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
