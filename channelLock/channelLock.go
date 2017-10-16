package main

import (
	"fmt"
)

func addToInt(A *int) {
	*A++

}

func main() {

	IGo := make(chan int)
	YouGo := make(chan int)
	IDone := make(chan int)
	YouDone := make(chan int)

	A := 0

	go func() {
		for i := 0; i < 10000000; i++ {
			<-IGo
			addToInt(&A)
			YouGo <- 1
		}
		IDone <- 1
	}()

	go func() {
		for i := 0; i < 10000000; i++ {
			<-YouGo
			addToInt(&A)
			IGo <- 1
		}
		YouDone <- 1
	}()
	IGo <- 1
	fmt.Println("here")
	<-IDone
	//<-IGo
	fmt.Println("there")
	<-YouDone
	fmt.Println("everywhere")
	fmt.Println(A)

}
