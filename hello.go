package main

import (
	"fmt"

	"github.com/go_introduction/go_introduction/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
