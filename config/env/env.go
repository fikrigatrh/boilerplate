package env

import (
	"boilerplate/models"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
)

var Config models.ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		log.Fatalf("cannot open file .env")
	}
}

func loadConfig() error {
	err := godotenv.Load() //library untuk manggil data dari env
	if err != nil {
		logrus.Fatal(err, "config/env: load file env")
	}

	// ini lib dari "github.com/caarlos0/env"
	err = env.Parse(&Config) //parsing isi dari struct server config
	if err != nil {
		return err
	}

	return err
}