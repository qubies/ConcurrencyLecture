package helper

import (
	"time"
)

// GoogleConnection ...
type GoogleConnection struct {
}

// ConnectToGoole ...
func ConnectToGoole() chan GoogleConnection {
	ret := make(chan GoogleConnection)
	go func() {
		<-time.NewTimer(3 * time.Second).C
		ret <- GoogleConnection{}
	}()
	return ret
}

// AskGoogle ...
func (*GoogleConnection) AskGoogle(question int) chan string {
	ret := make(chan string)
	go func() {
		<-time.NewTimer(3 * time.Second).C
		ret <- "true"
	}()
	return ret
}

// QueryDatabase ...
func QueryDatabase(query string) chan []int {
	ret := make(chan []int)
	go func() {
		<-time.NewTimer(3 * time.Second).C
		ints := make([]int, 0, 10)
		for i := 0; i < 10; i++ {
			ints = append(ints, i)
		}
		ret <- ints
	}()
	return ret
}
