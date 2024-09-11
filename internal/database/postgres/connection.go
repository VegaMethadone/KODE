package postgres

import (
	"fmt"
	"kode/config"
)

func getConnStr(conf *config.Config) string {
	if conf.Postgres.ConnStr == "" {
		conf.Postgres.ConnStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.DatabaseName, conf.Postgres.Sslmode, conf.Postgres.Host)
	}

	return conf.Postgres.ConnStr
}

func NewPostgres() *Postgres {
	newPostgres := &Postgres{}
	return newPostgres
}
