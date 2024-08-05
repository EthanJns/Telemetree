package main

import (
	"fmt"
	"promise/promise"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	p := promise.Create[int]()
	j := 9
	fmt.Println(p)
	p.OnComplete(testFunc)

	time.Sleep(1 * time.Second)
	p.Complete(&j)
	fmt.Println(p)
}

func testFunc(i *int) {
	fmt.Println("This is from the testFunctionSent to the the OnComplete Method")
	fmt.Println(*i)
}
