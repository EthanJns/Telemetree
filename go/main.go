package main

import (
	"fmt"
	"promise/promise"
)

func main() {
	fmt.Println("Hello, World!")
	p := promise.Create[int]()
	j := 9
	fmt.Println(p)
	p.OnComplete(testFunc)
	p.Complete(&j)
}

func testFunc(i *int) {
	fmt.Println("This is from the testFunctionSent to the the OnComplete Method")
	fmt.Println(*i)
}
