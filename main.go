package main

import (
	"fmt"
	"os"
)

// ----
// MAIN
// ----

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
