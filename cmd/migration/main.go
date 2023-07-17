package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"smart-chat/cmd/migration/migrations"
	"smart-chat/internal/config"
)

func main() {
	cfg := config.NewConfigService()

	db, err := gorm.Open(mysql.Open(cfg.GetConfig().Database.DbConnString), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})

	migrationsToExec := migrations.GetMigrationsToExec()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsToExec)
	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")

}
