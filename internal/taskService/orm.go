package taskService // Или ваш пакет, если он не main

import (
 "gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task string `json:"task"`
	IsDone bool `json:"is_done"`
}