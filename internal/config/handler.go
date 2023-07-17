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

	config := &model.Config{
		Database: model.Database{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
		},
		DeepAi: model.DeepAi{
			URL:    os.Getenv("DEEP_AI_URL"),
			ApiKey: os.Getenv("DEEP_AI_KEY"),
		},
	}

	config.Database.DbConnString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Database)

	restPortString := os.Getenv("REST_PORT")
	if restPortString == "" {
		restPortString = "8888"
	}

	fmt.Println("pass port: ", restPortString)

	restPort, err := strconv.Atoi(restPortString)
	if err != nil {
		panic(err.Error())
	}

	config.RestPort = restPort
	c.Config = config
}

func (c *configService) GetConfig() *model.Config {
	return c.Config
}
