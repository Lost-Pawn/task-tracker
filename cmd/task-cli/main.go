package main

import (
	"fmt"
	"os"
	"task-tracker/internal/storage"
)

func main() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	fmt.Println("Loaded tasks:", tasks)

	if len(os.Args) < 2 {
		fmt.Println("No command provided. Use 'help' to see available commands.")
		return
	}

	switch (os.Args[1]) {
	case "add":
		fmt.Println("Add command selected.")
	case "list":
		fmt.Println("List command selected.")
	case "update":
		fmt.Println("Update command selected.")
	case "delete":
		fmt.Println("Delete command selected.")
	case "mark-in-progress":
		fmt.Println("Mark in progress command selected.")
	case "mark-done":
		fmt.Println("Mark completed command selected.")	
	case "help":
		fmt.Println("Help command selected.")
		fmt.Println("Available commands:")
		fmt.Println("  add     - Add a new task")
		fmt.Println("  list    - List all tasks")
		fmt.Println("  update  - Update an existing task")
		fmt.Println("  delete  - Delete a task")
		fmt.Println("  help    - Show this help message")	
		fmt.Println("  mark-in-progress - Mark a task as in progress")
		fmt.Println("  mark-done    - Mark a task as completed")
	default:
		fmt.Println("Unknown command. Use 'help' to see available commands.")
	}
}