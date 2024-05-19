package main

import (
	"example.com/greetings"
	"github.com/charmbracelet/log"
)

func main() {
    // Get a greeting message and print it.
    names := []string{"Gladys", "Samantha", "Darrin"}

    messages, err := greetings.Hellos(names)

    if err != nil {
        log.Fatal("Error: empty name", "err", err)
    }

    log.Print(messages)
}
