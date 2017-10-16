package main

import (
	"fmt"
)

func addToInt(A *int) {
	*A++

}

func main() {

	Go := make(chan int)
	Done := make(chan int)

	A := 0

	go func() {
		for i := 0; i < 10000000; i++ {
			<-Go
			addToInt(&A)
			Go <- 1
		}
		fmt.Println(1)
		Done <- 1
	}()

	go func() {
		for i := 0; i < 10000000; i++ {
			<-Go
			addToInt(&A)
			Go <- 1
		}
		fmt.Println(2)
		Done <- 1
	}()
	Go <- 1
	fmt.Println("here")
	<-Done
	fmt.Println("there")
	//<-IGo
	fmt.Println("everywhere")
	<-Done
	fmt.Println(A)

}
