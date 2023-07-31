package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202307141030_create_chat_message",
		Migrate: func(tx *gorm.DB) error {

			sql := `CREATE TABLE IF NOT EXISTS chat_message(
					id INT AUTO_INCREMENT PRIMARY KEY,
					chat_id int,
					question LONGTEXT,
					response_id VARCHAR(191),
					Response LONGTEXT,
					question_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					response_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					FOREIGN KEY (chat_id) REFERENCES chat(id)
);`
			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("chat_message")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
