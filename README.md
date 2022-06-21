# Introduction
This is a simple module to extend the default go inner libary.

## Time Format
a simple function, ` CommonTimeFormat(date) ` will result a YYYY-MM-dd HH:mm:ss style timestring.

## data structure
There are stack and queue implementation.   
Here is an example:

```go
import zx "github.com/notfresh/gotools"

func TestStack() {
	stack := zx.NewStack()
	stack.Push(1)
	stack.Push(2)
	res1 := stack.Pop()
	res2 := stack.Pop()
	res3 := stack.Pop()
	fmt.Println(res1, res2, res3)

}

func TestQueue() {
    q := zx.NewQueue()
    q.Push(12)
    q.Push(15)
    q.Pop()
    q.Print()
}


```



