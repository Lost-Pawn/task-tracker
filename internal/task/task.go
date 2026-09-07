package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID 			int 		`json:"id"`;
	Description string 		`json:"description"`;
	Status 		string 		`json:"status"`;
	CreatedAt 	time.Time 	`json:"createdAt"`;
	UpdatedAt 	time.Time 	`json:"updatedAt"`;
}

func AddTask(tasks []Task, description string) ([]Task, error) {
	newID := 1
	
	for _, task := range tasks {
		if task.ID >= newID {
			newID = task.ID + 1
		}
	}

	if description == "" {
		return tasks, fmt.Errorf("task description cannot be empty")

	}

	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	storageErr := storage.SaveTasks(tasks)
	if storageErr != nil {
		return tasks, fmt.Errorf("failed to save tasks: %w", storageErr)
	}

	return tasks, nil
}