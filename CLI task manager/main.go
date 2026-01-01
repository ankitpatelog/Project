package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const fileName = "tasks.json"

// ---------------- MAIN ----------------

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: add | list | done | delete | edit-title")
		return
	}

	command := os.Args[1]
	tasks := loadTasks()

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task title")
			return
		}
		addTask(os.Args[2], &tasks)

	case "list":
		listTasks(tasks)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID")
			return
		}
		id := toInt(os.Args[2])
		markDone(id, &tasks)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID")
			return
		}
		id := toInt(os.Args[2])
		deleteTask(id, &tasks)

	case "edit-title":
	if len(os.Args) < 4 {
		fmt.Println("Usage: edit-title <new-title> <id>")
		return
	}

	newTitle := os.Args[2]
	id := toInt(os.Args[3])
	editTitle(&tasks, newTitle, id)


	default:
		fmt.Println("Unknown command")
	}

	saveTasks(tasks)
}

// ---------------- TASK FUNCTIONS ----------------

func addTask(title string, tasks *[]Task) {
	id := 1
	if len(*tasks) > 0 {
		id = (*tasks)[len(*tasks)-1].ID + 1
	}

	task := Task{
		ID:    id,
		Title: title,
		Done:  false,
	}

	*tasks = append(*tasks, task)
	fmt.Println("Task added:", title)
}

func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
	}
}

func markDone(id int, tasks *[]Task) {
	for i, task := range *tasks {
		if task.ID == id {
			(*tasks)[i].Done = true
			fmt.Println("Task marked as done")
			return
		}
	}
	fmt.Println("Task not found")
}

func deleteTask(id int, tasks *[]Task) {
	for i, task := range *tasks {
		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			fmt.Println("Task deleted")
			return
		}
	}
	fmt.Println("Task not found")
}

// ---------------- FILE HANDLING ----------------

func loadTasks() []Task {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	_ = json.Unmarshal(file, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	_ = os.WriteFile(fileName, data, 0644)
}

// ---------------- HELPERS ----------------

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Invalid number")
		os.Exit(1)
	}
	return n
}

func editTitle(tasks *[]Task, newTitle string, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			(*tasks)[i].Title = newTitle
			fmt.Println("Title updated successfully")
			return
		}
	}
	fmt.Println("Invalid task ID")
}

