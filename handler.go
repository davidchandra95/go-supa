package main

import (
	"encoding/json"
	"net/http"
)

type TaskService interface {
	CreateTask(task NewTask) error
	GetTasks() ([]Task, error)
}

type Handler struct {
	taskSvc TaskService
}

func NewHandler(taskSvc TaskService) *Handler {
	return &Handler{taskSvc: taskSvc}
}

func (h *Handler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask NewTask

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.taskSvc.CreateTask(newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskSvc.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
