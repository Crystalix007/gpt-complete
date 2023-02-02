package main

import (
	"fmt"
	"io"
	"os"

	complete "github.com/Crystalix007/gpt-complete/lib"
)

func main() {
	err := doCompletion()
	if err != nil {
		fmt.Printf("Encountered error: %v", err)
		os.Exit(1)
	}
}

func doCompletion() error {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	completion, err := complete.Complete(string(input))
	if err != nil {
		return err
	}

	os.Stdout.WriteString(completion)
	return nil
}
