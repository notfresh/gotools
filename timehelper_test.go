package gotools

import (
	"fmt"
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

func TestCommonTimeParse(t *testing.T) {
	result := "2020-05-28 15:24:00"
	origin_date := time.Date(2020, 5, 28, 15, 24, 0, 0, time.Local)
	if date, _ := CommonTimeParse(result); date != origin_date {
		fmt.Println("parsed date ", date)
		fmt.Println("origin date ", origin_date)
		t.Fatal("CommonTimeFormat Error")
	}
}
