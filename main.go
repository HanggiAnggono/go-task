package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"hanggi.com/go-task/tasks"
)

func main() {
	args := os.Args
	var command string = "list"

	if len(args) > 1 {
		_, command = args[0], args[1]
	}

	switch command {
	default:
	case "list":
		taskList := tasks.GetTasks()
		tasks.PrintTasks(&taskList)
	case "add":
		description := strings.Join(args[2:], " ")
		err := tasks.AddTask(description)
		if err != nil {
			fmt.Printf("Error adding task: %s \n", err)
		} else {
			fmt.Printf("Added task: %s \n", description)
		}
	case "update":
		id, _ := strconv.Atoi(args[2])
		description := strings.Join(args[3:], " ")
		task, _ := tasks.UpdateTask(id, tasks.Task{
			Description: description,
		})
		fmt.Printf("Task updated %v \n", task.Description)
	case "wip":
		id, _ := strconv.Atoi(args[2])
		task, _ := tasks.MarkInProgress(id)
		fmt.Printf("Task updated %v \n", task.Status)
	case "done":
		id, _ := strconv.Atoi(args[2])
		task, _ := tasks.MarkDone(id)
		fmt.Printf("Task updated %v \n", task.Status)
	}

}
