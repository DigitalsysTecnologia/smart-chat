package model

type Config struct {
	DeepAi   DeepAi
	Database Database
	RestPort int
	BasePath string
}

type DeepAi struct {
	URL    string
	ApiKey string
}

type Database struct {
	DbConnString string
	RestPort     int
	Username     string
	Password     string
	Host         string
	Port         string
	Database     string
	ConnString   string
}
