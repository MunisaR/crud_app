package models

import "time"

type Task struct {
	ID uint `json:"id" gorm:"primary_key"`
	AssignedTo string `json:"assigned_to"`
	Task string `json:"task"`
	Deadline   time.Time `json:"deadline"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}