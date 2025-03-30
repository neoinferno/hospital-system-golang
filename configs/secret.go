package configs

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Secret struct {
	Postgres Postgres
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Db       string `env:"POSTGRES_DB"`
}

func GetSecret() (secret Secret) {
	fmt.Println("Loading .env file")
	if err := godotenv.Load(strings.Join([]string{"./configs", ".env"}, "/")); err != nil {
		panic(fmt.Errorf("unable to load .env file: %v", err))
	}
	if err := env.Parse(&secret); err != nil {
		panic(fmt.Errorf("failed to parse env vars: %v", err))
	}

	return secret
}
