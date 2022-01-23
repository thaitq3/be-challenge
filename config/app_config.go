package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	MySQL MySQL
	Redis Redis
}

type MySQL struct {
	DSN string
}

type Redis struct {
	Address string
}

var (
	Config *AppConfig
)

func LoadConfig() {
	Config = &AppConfig{
		MySQL: MySQL{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC",
				os.Getenv("MYSQL_USER"),
				os.Getenv("MYSQL_PASSWORD"),
				"mysql",
				"3306",
				os.Getenv("MYSQL_DATABASE"),
			),
		},
		Redis: Redis{
			Address: "redis:6379",
		},
	}

	s, _ := json.Marshal(Config)
	fmt.Println(string(s))
}
