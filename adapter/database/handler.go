package database

import "smart-chat/internal/model"

type databaseProvider struct {
	config *model.Config
}

func NewDatabaseProvider(cfg *model.Config) *databaseProvider {
	return &databaseProvider{
		config: cfg,
	}
}

func (t *databaseProvider) GetConnection() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(t.config.Database.DbConnString), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})

	return db, err
}
