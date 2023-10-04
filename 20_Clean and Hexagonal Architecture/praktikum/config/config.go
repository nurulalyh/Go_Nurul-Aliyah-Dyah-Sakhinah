package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

var (
	JWT_SECRET string
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     uint16
	DBName     string
}

func InitConfig() *Config {
	var config = readConfig()
	if config == nil {
		return nil
	}

	return config
}

func readConfig() *Config {
	var res = new(Config)

	res.DBUser = os.Getenv("DB_USER")
	res.DBPassword = os.Getenv("DB_PASSWORD")
	res.DBHost = os.Getenv("DB_HOST")
	cnvPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		logrus.Error("Cannot convert database port, ", err.Error())
	}
	res.DBPort = uint16(cnvPort)
	res.DBName = os.Getenv("DB_NAME")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	
	return res
}