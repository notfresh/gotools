package gotools

import "fmt"
import "testing"

func TestRandNumber(t *testing.T) {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(RandomTimeRange(time.Second, time.Second*3))
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(RandomBetween(10, 30))
	}
}
