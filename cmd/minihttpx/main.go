package main

import (
    "fmt"
    "os"

    "github.com/cristophercervantes/minihttpx/internal/runner"
)

var version = "dev"
var buildTime = ""

func main() {
    options := runner.ParseOptions()
    r := runner.New(options)
    if err := r.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
