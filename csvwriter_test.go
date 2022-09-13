package gotools

import (
	"fmt"
	"testing"
	"time"
)

type S2 struct {
	A string // 小写的不输出，大写的才行
	B time.Time
}

type ZX struct {
	Name string
	Age  int64
	S    S2
}

type ZXI interface {
	f1()
}

func (z *ZX) f1() {
	fmt.Println("abcd")
}

func (z *ZX) String() string {
	return fmt.Sprintf("%s,%d", z.Name, z.Age)
}

func TestCSVWriter(t *testing.T) {
	s2 := &S2{"abc", time.Now()}
	jl, _ := NewCSVWriter("13-test.csv")
	zx := &ZX{
		Name: "zhengxu",
		Age:  18,
		S:    *s2,
	}
	jl.writeStr("Name, Age")
	jl.write(zx)
	jl.Close()
}
