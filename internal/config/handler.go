package config

import (
	"fmt"
	"os"
	"smart-chat/internal/model"
	"strconv"
)

type configService struct {
	Config *model.Config
}

func NewConfigService() *configService {
	cfgService := &configService{}
	cfgService.loadConfig()

	return cfgService
}

func (c *configService) loadConfig() {

	c.Config.Database = model.Database{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}

	c.Config.Database.DbConnString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Config.Database.Username, c.Config.Database.Password, c.Config.Database.Host, c.Config.Database.Port, c.Config.Database.Database)

	restPortString := os.Getenv("PORT")
	if restPortString == "" {
		restPortString = "8030"
	}

	fmt.Println("pass port: ", restPortString)

	restPort, err := strconv.Atoi(restPortString)
	if err != nil {
		panic(err.Error())
	}

	c.Config.Database.RestPort = restPort
}

func (c *configService) GetConfig() *model.Config {
	return c.Config
}
