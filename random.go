package gotools

import "fmt"
import "testing"
import "time"
import "math/rand"

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

func TestRandNumber(t *testing.T) {

	for i := 0; i < 10; i++ {
		fmt.Println(RandomTimeRange(time.Second, time.Second*3))
	}
}