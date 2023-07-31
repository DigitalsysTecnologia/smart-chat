package model

type Config struct {
	DeepAi                DeepAi
	Database              Database
	RestPort              int
	AuthorizerApiEndpoint string
}

type DeepAi struct {
	URL    string
	ApiKey string
}

type Database struct {
	DbConnString string
	Username     string
	Password     string
	Host         string
	Port         string
	Database     string
	ConnString   string
}
