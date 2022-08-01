package gotools

import "time"

const CommonTimeStr string = "2006-01-02 15:04:05"

func CommonTimeFormat(t time.Time) string {
	return t.Format(CommonTimeStr)
}

func CommonTimeParse(s string) (t time.Time, e error) {
	t, e = time.Parse(CommonTimeStr, s)
	return
}


func TimeNearEqual(time1 time.Time, time2 time.Time) bool {
	if time1.After(time2) {
		return time1.Sub(time2) <= time.Millisecond
	}
	return time2.Sub(time1) <= time.Millisecond
}
