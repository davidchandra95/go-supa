package main

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	Task      string    `json:"task"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewTask struct {
	Task string `json:"task"`
}
