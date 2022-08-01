package gotools

import "time"
import "math/rand"

func RandomBetween(t1, t2 int64) int64 {
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)

	n1, n2 := t1, t2
	if n1 > n2 {
		n2, n1 = n1, n2
	}
	offset := n2 - n1
	ret := int64(r2.Float64()*float64(offset)) + n1
	return ret
}
