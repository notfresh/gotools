package gotools

import (
	"testing"
	"time"
)

func TestCommonTimeFormat(t *testing.T) {
	result := "2020-05-28 15:24:00"
	date := time.Date(2020, 5, 28, 15, 24, 0, 0, time.Local)
	if result != CommonTimeFormat(date) {
		print(CommonTimeFormat(date))
		t.Fatal("CommonTimeFormat Error")
	}
}
