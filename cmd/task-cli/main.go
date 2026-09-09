package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/internal/storage"
	"task-tracker/internal/task"
)

func main() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

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
		if len(os.Args) < 4 {
			fmt.Println("Error: Task ID and new description are required for 'update' command.")
			return
		}

		newDescription := os.Args[3]
		if newDescription == "" {
			fmt.Println("Error: New description is required for 'update' command.")
			return
		}

		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Please provide a valid Task ID")
			return
		}

		tasks, err = task.UpdateTaskDescription(tasks, taskID, newDescription)
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}

		saveErr := storage.SaveTasks(tasks)
		if saveErr != nil {
			fmt.Println("Error saving tasks:", saveErr)
		} else {
			fmt.Printf("Successfully updated task ID %d to description '%s'\n", taskID, newDescription)
		}	

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task ID is required for 'delete' command.")
			return
		}

		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Please provide a valid Task ID")
			return
		}
		
		tasks, err = task.DeleteTask(tasks, taskID)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}

		saveErr := storage.SaveTasks(tasks)
		if saveErr != nil {
			fmt.Println("Error saving tasks:", saveErr)
		} else {
			fmt.Printf("Successfully deleted task ID %d\n", taskID)
		}

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task ID is required for 'mark-in-progress' command.")
			return
		}

		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Please provide a valid Task ID")
			return
		}
		
		tasks, err = task.MarkInProgress(tasks, taskID)
		if err != nil {
			fmt.Println("Error marking task as in-progress:", err)
			return
		}

		saveErr := storage.SaveTasks(tasks)
		if saveErr != nil {
			fmt.Println("Error saving tasks:", saveErr)
		} else {
			fmt.Printf("Successfully marked task ID %d as in-progress\n", taskID)
		}

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task ID is required for 'mark-done' command.")
			return
		}

		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Please provide a valid Task ID")
			return
		}
		
		tasks, err = task.MarkDone(tasks, taskID)
		if err != nil {
			fmt.Println("Error marking task as done:", err)
			return
		}

		saveErr := storage.SaveTasks(tasks)
		if saveErr != nil {
			fmt.Println("Error saving tasks:", saveErr)
		} else {
			fmt.Printf("Successfully marked task ID %d as done\n", taskID)
		}
		
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