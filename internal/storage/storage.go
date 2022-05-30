package storage

import (
	"context"
	"otus/internal/models"
	"strconv"
	"sync"
)

type storage struct {
	mu     sync.RWMutex
	tasks  models.TaskList
	lastId int64
}

func (s *storage) GetTaskList(ctx context.Context) models.TaskList {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var tasks models.TaskList
	copy(tasks.Tasks, s.tasks.Tasks)
	return tasks
}

func (s *storage) CreateTask(ctx context.Context, task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastId++
	task.Id = strconv.FormatInt(s.lastId, 16)
	s.tasks.Tasks = append(s.tasks.Tasks, task)
	return task
}

func (s *storage) DeleteTask(ctx context.Context, id string) {

}

type Storage interface {
	GetTaskList(ctx context.Context) models.TaskList
	CreateTask(ctx context.Context, task models.Task) models.Task
	DeleteTask(ctx context.Context, id string)
}

func New() *storage {
	return &storage{}
}
