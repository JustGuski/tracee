package root

import (
	"os"
	"fmt"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tracee test")
		return
	}
	args := os.Args[1:]
	command := args[0]

	switch command {
	case "test":
		fmt.Println("Running tests...")
		return
	default:
		fmt.Println("Unknown command!")
		return
	}
}
