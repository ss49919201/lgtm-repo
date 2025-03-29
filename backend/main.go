package main

import (
	"fmt"
	"os"
)

func runServer() error {
	return nil
}

func main() {
	if err := runServer(); err != nil {
		fmt.Printf("failed to run server: %s\n", err)
		os.Exit(1)
	}
}
