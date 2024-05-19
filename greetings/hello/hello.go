package main

import (
	"os"

	"example.com/greetings"
	"github.com/charmbracelet/log"
	"github.com/erikgeiser/promptkit/textinput"
)

func main() {
    // Get a greeting message and print it.
    input := textinput.New("What is your name?")
    input.InitialValue = os.Getenv("USER")
    input.Placeholder = "Your name cannot be empty"

    name, err := input.RunPrompt()

    if err != nil {
        log.Fatal("Error: failed to retrieve user input", "err", err)
    }

    message, err2 := greetings.Hello(name)
    if err2 != nil {
        log.Fatal("Error: empty name", "err2", err2)
    }

    log.Print(message)
}
