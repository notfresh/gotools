package gotools

import "time"

const CommonTimeStr string = "2006-01-02 15:04:05"

func CommonTimeFormat(t time.Time) string {
	return t.Format(CommonTimeStr)
}
