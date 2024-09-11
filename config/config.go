package config

func NewConfig() *Config {
	newConf := &Config{
		Network: NetWork{
			Ip:   "0.0.0.0", // если локально  проверять - 127.0.0.1  Если в  dcoker-compose - 0.0.0.0,
			Port: "8080",
		},
		Postgres: DataBase{
			Host:         "postgres", //localhost если локально. postgres - если в контенере docker compose
			Port:         "5432",
			Username:     "postgres",
			Password:     "0000",
			DatabaseName: "testDB",
			Sslmode:      "disable",
		},
	}

	return newConf
}
