package gotools

import (
	"math/rand"
	"time"
)

const CommonTimeStr string = "2006-01-02 15:04:05"
const DenseTimeStr string = "20060102150405"

func CommonTimeFormat(t time.Time) string {
	return t.Format(CommonTimeStr)
}

func DenseTimeFormat(t time.Time) string {
	return t.Format(DenseTimeStr)
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

func RandomTimeRange(t1, t2 time.Duration) time.Duration {
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)

	n1, n2 := int64(t1), int64(t2)
	if n1 > n2 {
		n2, n1 = n1, n2
	}
	offset := n2 - n1
	ret := int64(r2.Float64()*float64(offset)) + n1
	return time.Duration(ret)
}

func WaitRandomTimeRange(t1, t2 time.Duration) {
	ret := RandomTimeRange(t1, t2)
	//fmt.Println("Random Sleep for ", ret)
	time.Sleep(ret)
}
