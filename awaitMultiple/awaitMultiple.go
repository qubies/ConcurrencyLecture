package main

import (
	"github.com/qubies/ConcurrencyLecture/helper"
)

// What if we need to wait for multiple things to finish before proceeding?

// 1. Make a database query
// 2. Make a connection to google
// 3. Foreach row in the query result perform a google search
func main() {
	// Method one (BAD)
	// Wait for each previous step to complete before continuing
	conn := <-helper.ConnectToGoole()
	data := <-helper.QueryDatabase("SELECT * FROM FOO;")
	results := make([]string, 0)

	for datum := range data {
		results = append(results, <-conn.AskGoogle(datum))
	}

	// What's wrong with method one?

	// Method two (Better, but still BAD)
	// Query database and connect to google at the same time
	results = make([]string, 0)
	tasks := 2
	for tasks > 0 {
		select {
		case conn = <-helper.ConnectToGoole():
			tasks--
		case data = <-helper.QueryDatabase("SELECT * FROM FOO;"):
			tasks--
		}
	}

	for datum := range conn.AskGoogle(datum) {
		results = append(results, <-)
	}
}
