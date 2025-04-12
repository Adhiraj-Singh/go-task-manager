package main

import (
	"fmt"
	"os"
	"go-task-manager/tasks"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-task-manager [add}list]done|delete] [task]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		tasks.Add(os.Args[2:])
	case "list":
		tasks.List()
	case "done":
		tasks.MarkDone(os.Args[2:])
	case "delete":
		tasks.Delete(os.Args[2:])
	default:
		fmt.Println("unknown command", command)
	}
}