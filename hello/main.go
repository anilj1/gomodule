package main

import (
	"errors"
	"log"
	"math"

	"github.com/anilj1/gomodule/greetings"
	"rsc.io/quote"
)

/*
*

	Replace the modules from local code:
	  - go mod edit -replace example.com/greetings=../greetings

	Synchronize dependencies that are not yet tracked in the module
	  - go mod tidy
*/

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	log.Println(math.Max(2, 4))
	log.Println(errors.New("New error message"))

	// Greet a user
	message, err := greetings.Hello("Peter!")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	log.Println(message)
	log.Println(quote.Go())
}
