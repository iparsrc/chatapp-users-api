package utils

import "time"

const (
	timeLayout = "2006-01-02"
)

func NowStr() string {
	return time.Now().UTC().Format(timeLayout)
}
