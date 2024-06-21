package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sergiman94/task-api/internal/models"
)

type allTasks []models.Task

var tasks = allTasks{
	{
		ID:      uuid.Must(uuid.NewRandom()).String(),
		Name:    "Sergio",
		Content: "Programming",
	},
}

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (t *TaskService) CreateTask(ctx context.Context, taskRequest models.Task) (models.Task, error) {
	task := models.Task{
		ID:      uuid.Must(uuid.NewRandom()).String(),
		Name:    taskRequest.Name,
		Content: taskRequest.Content,
	}

	tasks = append(tasks, task)

	return task, nil
}
