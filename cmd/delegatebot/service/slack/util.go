package slack

import (
	"math"
	"strconv"
	"time"
)

func convertSlackTimestamp(timestamp string) time.Time {
	timeFloat, err := strconv.ParseFloat(timestamp, 10)
	if err != nil {
		panic(err) // TODO unpanic?
	}

	sec, dec := math.Modf(timeFloat)

	return time.Unix(int64(sec), int64(dec*(1e9)))
}