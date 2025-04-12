package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	
	"strconv"
)

type Task struct {
	Description string `json: "description"`
	Done bool `json:"done"`
}

const dataFile = "tasks.json"

func loadTasks() []Task {
	var tasks []Task

	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return tasks
	}

	json.Unmarshal(file, &tasks)
	return tasks
}

// save tasks
func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", " ")
	ioutil.WriteFile(dataFile, data, 0644)
}

// add task
func Add(args []string) {
	if(len(args) == 0) {
		fmt.Println("no task provided")
		return
	}

	task := Task{Description: args[0], Done: false}
	tasks := loadTasks()
	tasks = append(tasks, task)
	saveTasks(tasks)

	// confirm print
	fmt.Println("Added task:", args[0])
}

// list out items
func List() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("no tasks yet")
		return
	}

	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i + 1, status, task.Description)
	}
}

func MarkDone(args []string) {
	if len(args) == 0 {
		fmt.Println("Provide task number to mark as done.")
		return
	}

	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		fmt.Println("invalid task number")
		return
	}

	tasks := loadTasks()
	if index > len(tasks) {
		fmt.Println("task number out of range.")
		return
	}

	tasks[index-1].Done = true
	saveTasks(tasks)
	fmt.Println("Task marked as done")
}

func Delete(args []string) {
	if len(args)==0 {
		fmt.Println("provide task number to delete")
		return
	}

	index, err := strconv.Atoi(args[0])
    if err != nil || index < 1 {
        fmt.Println("Invalid task number.")
        return
    }

	tasks := loadTasks()
    if index > len(tasks) {
        fmt.Println("Task number out of range.")
        return
    }

	deleted := tasks[index-1].Description
    tasks = append(tasks[:index-1], tasks[index:]...)
    saveTasks(tasks)

    fmt.Println("Deleted task:", deleted)
}