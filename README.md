# Task Tracker CLI

A simple command-line tool to track and manage your tasks, built in Go as part of the [roadmap.sh backend projects](https://roadmap.sh/projects/task-tracker).

Tasks are stored locally in a `tasks.json` file.

## Features

- Add, update, and delete tasks
- Mark tasks as in-progress or done
- List all tasks, or filter by status

## Installation

Clone the repo and build the binary:

```bash
git clone "https://github.com/Lost-Pawn/task-tracker.git"
cd task-tracker
go install ./cmd/task-cli
```

Make sure `$(go env GOPATH)/bin` is on your `PATH` so the `task-cli` command is available globally.

## Usage

```bash
# Add a new task
task-cli add "Buy groceries"

# Update a task's description
task-cli update 1 "Buy groceries and cook dinner"

# Delete a task
task-cli delete 1

# Mark a task as in progress / done
task-cli mark-in-progress 1
task-cli mark-done 1

# List tasks
task-cli list
task-cli list todo
task-cli list in-progress
task-cli list done
```

## Task Properties

Each task has the following fields:

| Field       | Description                          |
|-------------|---------------------------------------|
| `id`        | Unique identifier for the task        |
| `description` | Task description                    |
| `status`    | `todo`, `in-progress`, or `done`      |
| `createdAt` | Timestamp when the task was created   |
| `updatedAt` | Timestamp when the task was last updated |

## Project Structure

```
task-tracker/
├── cmd/task-cli/     # entry point, CLI argument parsing
├── internal/task/    # task data model and logic
└── internal/storage/ # JSON file persistence
```