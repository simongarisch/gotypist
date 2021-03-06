package main

import (
	"log"
)

// Quit implements error
// Throw this when quitting the app.
type Quit struct{}

func (quit Quit) Error() string {
	return "Quit program."
}

func errorHandling(err error) {
	log.Fatal(err)
}
