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

	return tasks, nil
}

func ListTasks(tasks []Task, status string) ([]Task, error) {
	if status == "" {
		return tasks, nil
	}
	
	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		} 
	}
	return filteredTasks, nil
}

func UpdateTaskDescription(tasks []Task, id int, description string) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		} 
	}
	return tasks, fmt.Errorf("No tasks found with that ID")
}