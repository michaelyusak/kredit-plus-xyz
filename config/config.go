package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type ServiceConfig struct {
	Port           string   `json:"port"`
	GracefulPeriod int64    `json:"graceful_perion"`
	Postgres       DBConfig `json:"postgres"`
}

func Init(log *logrus.Logger) ServiceConfig {
	configPath := os.Getenv("KREDIT_PLUS_USERS_SERVICE_CONFIG")

	var config ServiceConfig

	configFile, err := os.OpenFile(configPath, os.O_RDONLY, 0644)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": fmt.Sprintf("[config][Init][OpenFile] error: %s", err.Error()),
		}).Fatal("error initiating config file")

		return config
	}
	defer configFile.Close()

	var configData []byte

	_, err = configFile.Read(configData)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": fmt.Sprintf("[config][Init][Read] error: %s", err.Error()),
		}).Fatal("error initiating config file")

		return config
	}

	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": fmt.Sprintf("[config][Init][Unmarshal] error: %s", err.Error()),
		}).Fatal("error initiating config file")

		return config
	}

	return config
}
