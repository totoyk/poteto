package config

import "github.com/caarlos0/env/v10"

type Env struct {
	// DB
	DBType string `env:"DB_TYPE,required" envDefault:"mysql"`
	DBHost string `env:"DB_HOST,required" envDefault:"db"`
	DBPort string `env:"DB_PORT,required" envDefault:"3306"`
	DBUser string `env:"DB_USER,required" envDefault:"root"`
	DBPass string `env:"DB_PASS,required" envDefault:""`
	DBName string `env:"DB_NAME,required" envDefault:"prototype"`
	TZ     string `env:"TZ" envDefault:"Asia/Tokyo"`
}

func LoadEnv() (*Env, error) {
	var conf Env
	if err := env.Parse(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
