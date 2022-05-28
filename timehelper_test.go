package gotools

import "time"

const CommonTimeStr string = "2006-1-2 15:4:5"

func CommonFormat(t time.Time) string {
	return t.Format(CommonTimeStr)
}
