package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"task-tracker/internal/task"
)

func LoadTasks() ([]task.Task, error) {
	taskpath := filepath.Join(".", "tasks.json")
	_, err := os.Stat(taskpath)
	if errors.Is(err, fs.ErrNotExist) {
		return []task.Task{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to stat tasks file: %w", err)
	}
	tasks, err := os.Open(taskpath)
	if err != nil {
		return nil, fmt.Errorf("failed to open tasks file: %w", err)
	}
	defer tasks.Close()

	var tasksList []task.Task
	err = json.NewDecoder(tasks).Decode(&tasksList)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tasks file: %w", err)
	}

	return tasksList, nil
}

func SaveTasks(tasks []task.Task) error {
	taskpath := filepath.Join(".", "tasks.json")
	tasksFile, err := os.Create(taskpath)
	if err != nil {
		return fmt.Errorf("failed to create tasks file: %w", err)
	}
	defer tasksFile.Close()

	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode tasks to file: %w", err)
	}

	_, err = tasksFile.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write tasks to file: %w", err)
	}

	return nil
}