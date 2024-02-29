package main

import "fmt"

func main() {
	c := make(chan int) //future
	go testFunc(10, c)  //async
	cv := <-c           //await
	fmt.Println(cv)
}

func testFunc(input int, result chan<- int) {
	result <- input * 2
	fmt.Println(input)
}
