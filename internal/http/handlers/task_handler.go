package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sergiman94/task-api/internal/models"
)

type TaskServiceI interface {
	CreateTask(ctx context.Context, taskRequest models.Task) (models.Task, error)
}

type TaskHandler struct {
	Service TaskServiceI
}

func NewTaskHandler(service TaskServiceI) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskRequest, err := models.FromRequest(r)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	taskResponse, err := h.Service.CreateTask(r.Context(), taskRequest)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(taskResponse)
}
