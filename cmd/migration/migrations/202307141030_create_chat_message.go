package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202307141030_create_chat_message",
		Migrate: func(tx *gorm.DB) error {
			sql := `CREATE TABLE "ChatMessage" (
    					"ID" SERIAL PRIMARY KEY,
   						"UserID" VARCHAR(50),
   						"Question" VARCHAR(50),
   						"ResponseID" VARCHAR(50),
   						"Response" VARCHAR(50),
    					"QuestionDate" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    					"ResponseDate" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    					"ChatID" BIGINT,
    					FOREIGN KEY ("ChatID")REFERENCES "CHAT"("ID")
					);`

			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("ChatMessage")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
