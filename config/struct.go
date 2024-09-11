package config

type Config struct {
	Network  NetWork
	Postgres DataBase
}

type NetWork struct {
	Ip   string
	Port string
}

type DataBase struct {
	ConnStr      string
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	Sslmode      string
}
