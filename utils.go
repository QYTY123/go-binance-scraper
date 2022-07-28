package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func GetOSEnv(key, defaultValue string) string {
	log.SetFormatter(&log.JSONFormatter{})

	//err := godotenv.Load()
	value := os.Getenv(key)

	//if err != nil {
	//	log.Warn(err)
	//}

	if value == "" {
		return defaultValue
	}
	return value
}
