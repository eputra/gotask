package main

import (
	"fmt"
	"os"

	"github.com/eputra/gotask/gtslice"
	"github.com/eputra/gotask/gtstruct"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: gotask <task>")
		os.Exit(1)
	} else {
		switch os.Args[1] {
		case "struct":
			gtstruct.Print()
		case "slice":
			gtslice.Print()
		default:
			fmt.Println("Task not found")
		}
	}
}
