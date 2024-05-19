package main

import (
    "fmt"
    "os"
    "rsc.io/quote"
    "github.com/erikgeiser/promptkit/confirmation"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    input := confirmation.New("Ready to go?", confirmation.Undecided)

    ready, err := input.RunPrompt()
    if err != nil {
        fmt.Printf("Error: %v\n", err)

        os.Exit(1)
    }

    _ = ready
}
