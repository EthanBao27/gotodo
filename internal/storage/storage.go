package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// defination of a basic task
type Task struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
}

var filePath string

// default save path is ~/.gotodo/tasks.json
func init() {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".gotodo")
	_ = os.MkdirAll(dir, 0755)
	filePath = filepath.Join(dir, "tasks.json")
}

func SetPath(p string) {
	filePath = p
}

func GetCurrentPath() string {
	return filePath
}

// load all tasks from json file
func load() ([]Task, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}
	if len(b) == 0 {
		return []Task{}, nil
	}
	var tasks []Task
	if err := json.Unmarshal(b, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// save tasks to file
func save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

// list all tasks
func List() ([]Task, error) {
	return load()
}

// add a new task
func Add(content string) (Task, error) {
	tasks, err := load()
	if err != nil {
		return Task{}, err
	}
	nextID := 1
	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
	nt := Task{
		ID:        nextID,
		Content:   content,
		Done:      false,
		CreatedAt: time.Now().Local().String(),
	}
	tasks = append(tasks, nt)
	return nt, save(tasks)
}

// set the status of a task
func SetDone(id int, done bool) error {
	tasks, err := load()
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = done
			return save(tasks)
		}
	}
	return fmt.Errorf("task %d not found", id)
}

// delete a task
func Delete(id int) error {
	tasks, err := load()
	if err != nil {
		return err
	}
	idx := -1
	for i, t := range tasks {
		if t.ID == id {
			idx = i
			break
		}
	}
	if idx < 0 {
		return fmt.Errorf("task %d not found", id)
	}
	tasks = append(tasks[:idx], tasks[idx+1:]...)
	return save(tasks)
}

// overwrite file by an empty list
func Clear() error {
	return save([]Task{})
}

