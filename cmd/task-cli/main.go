package main

import (
	"fmt"
	"os"
	"task-tracker/internal/storage"
	"task-tracker/internal/task"
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
		description := os.Args[2]
		if len(os.Args) < 3 {
			fmt.Println("Error: Task description is required for 'add' command.")
			return
		}

		tasks, err = task.AddTask(tasks, description)
		if err != nil {
			fmt.Println("Error adding task:", err)
		}

		saveErr := storage.SaveTasks(tasks)
		if saveErr != nil {
			fmt.Println("Error saving tasks:", saveErr)
		} else {
			fmt.Printf("Successfully added task: %d", tasks[len(tasks)-1].ID)
		}

	case "list":
		status := ""
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		if status == "" {
			fmt.Println("You might wanna filter through status: todo, done, in-process")
		}
		filteredTasks, err := task.ListTasks(tasks, status)
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}

		if len(filteredTasks) == 0 {
			fmt.Println("No tasks found.")
		} else {
			fmt.Println("Tasks:")
			for _, t := range filteredTasks {
				fmt.Printf("ID: %d, Description: %s, Status: %s\n", t.ID, t.Description, t.Status)
			}
		}
		
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