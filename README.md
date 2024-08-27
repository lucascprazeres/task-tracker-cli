# Task Tracker

Track your tasks via terminal. 
This is my solution for the [first challenge](https://roadmap.sh/projects/task-tracker) of the backend developer roadmap!

## How to use?

```bash
# download the repo
git clone https://github.com/lucascprazeres/task-tracker-cli

# go to the project dir
cd task-tracker-cli

# download dependencies
go mod download

# build the project
go build -o task-tracker cmd/main.go

# run it!
./task-tracker list
```

## Commands

```bash
# Adding a new task
./task-tracker add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./task-tracker update 1 "Buy groceries and cook dinner"
./task-tracker delete 1

# Marking a task as in progress or done
./task-tracker mark-in-progress 1
./task-tracker mark-done 1

# Listing all tasks
./task-tracker list

# Listing tasks by status
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```