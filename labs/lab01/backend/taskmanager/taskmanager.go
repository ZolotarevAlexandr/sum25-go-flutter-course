package taskmanager

import (
	"errors"
	"time"
)

// Predefined errors
var (
	ErrTaskNotFound = errors.New("task not found")
	ErrEmptyTitle   = errors.New("title cannot be empty")
)

// Task represents a single task
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

// TaskManager manages a collection of tasks
type TaskManager struct {
	tasks  map[int]Task
	nextID int
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

// AddTask adds a new task to the manager, returns an error if the title is empty, and increments the nextID
func (tm *TaskManager) AddTask(title, description string) (Task, error) {
	if title == "" {
		return Task{}, ErrEmptyTitle
	}
	newTask := Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}
	tm.tasks[newTask.ID] = newTask
	tm.nextID++
	return newTask, nil
}

// UpdateTask updates an existing task, returns an error if the title is empty or the task is not found
func (tm *TaskManager) UpdateTask(id int, title, description string, done bool) error {
	if title == "" {
		return ErrEmptyTitle
	}
	originalTask, ok := tm.tasks[id]
	if !ok {
		return ErrTaskNotFound
	}
	newTask := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Done:        done,
		CreatedAt:   originalTask.CreatedAt,
	}
	tm.tasks[id] = newTask
	return nil
}

// DeleteTask removes a task from the manager, returns an error if the task is not found
func (tm *TaskManager) DeleteTask(id int) error {
	_, ok := tm.tasks[id]
	if !ok {
		return ErrTaskNotFound
	}
	delete(tm.tasks, id)
	return nil
}

// GetTask retrieves a task by ID, returns an error if the task is not found
func (tm *TaskManager) GetTask(id int) (Task, error) {
	task, ok := tm.tasks[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

// ListTasks returns all tasks, optionally filtered by done status, returns an empty slice if no tasks are found
func (tm *TaskManager) ListTasks(filterDone *bool) []Task {
	result := make([]Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		if (filterDone == nil) || (*filterDone && task.Done) || (!*filterDone && !task.Done) {
			result = append(result, task)
		}
	}
	return result
}
