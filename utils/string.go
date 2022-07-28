package utils

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

func StringToInt64(s string) (int64, error) {
	numero, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		log.Warn(err)
		return 0, err
	}
	return numero, err
}

func StringToFloat64(s string) (float64, error) {
	numero, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Warn(err)
		return 0, err
	}
	return numero, err
}
