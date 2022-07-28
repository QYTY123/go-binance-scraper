package utils

import (
	"strconv"
	"time"
)

func convertUnixTime(unixTime string) time.Time {
	i, err := strconv.ParseInt(unixTime, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm
}
