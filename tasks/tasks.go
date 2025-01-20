package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"
)

var StoreFileName string = "tasks.json"

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tasks = []Task

func GetTasks() Tasks {
	var data []byte
	var err error
	data, err = os.ReadFile(StoreFileName)
	var tasks Tasks

	if err != nil {
		fmt.Println(map[string]any{"err": err})
		if _, err := os.Create(StoreFileName); err != nil {
			data, _ = os.ReadFile(StoreFileName)
		}
	}

	json.Unmarshal(data, &tasks)

	return tasks
}

func PrintTasks(tasks *Tasks) {
	greenBG := "\033[42m"
	greyBG := "\033[47m"
	yellowBG := "\033[43m"
	reset := "\033[0m"

	for _, task := range *tasks {
		BG := map[string]string{
			"TODO":  greyBG,
			"DOING": yellowBG,
			"DONE":  greenBG,
		}[task.Status]

		fmt.Printf("%v %s%v: %s %v\n", task.ID, BG, task.Status, reset, task.Description)
		fmt.Printf("Updated: %v\n", task.UpdatedAt)
		fmt.Print("-------------------\n")
	}
}

func AddTask(description string) error {
	tasks := GetTasks()
	task := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "TODO",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	err := SaveTasks(&tasks)

	return err
}

func GetTaskById(id int) (*Task, error) {
	tasks := GetTasks()
	x := sort.Search(len(tasks), func(i int) bool {
		return tasks[i].ID == id
	})

	if x < 0 {
		return nil, errors.New("Task Not Found")
	}

	task := &tasks[x]

	return task, nil
}

func UpdateTask(id int, task Task) (*Task, error) {
	tasks := GetTasks()
	tsk, err := GetTaskById(id)

	if err != nil {
		return nil, err
	}

	desc := task.Description
	if desc == "" {
		desc = tsk.Description
	}
	status := task.Status
	if status == "" {
		status = tsk.Status
	}

	tsk = &Task{
		ID:          tsk.ID,
		Description: desc,
		Status:      status,
		CreatedAt:   tsk.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = *tsk
		}
	}

	err = SaveTasks(&tasks)

	return tsk, err
}

func MarkInProgress(id int) (*Task, error) {
	task, err := UpdateTask(id, Task{
		Status: "DOING",
	})

	return task, err
}

func MarkDone(id int) (*Task, error) {
	task, err := UpdateTask(id, Task{
		Status: "DONE",
	})

	return task, err
}

func SaveTasks(tasks *Tasks) error {
	var err error
	var data []byte
	data, _ = json.MarshalIndent(tasks, "", "	")
	err = os.WriteFile(StoreFileName, data, 0644)

	return err
}
