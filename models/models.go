package models

import (
	"time"
)

type Movie struct {
	Id       string        `json:"id" gorm:"primaryKey,column:id"`
	Name     string        `json:"name" binding:"required"`
	Language string        `json:"language" binding:"required"`
	Length   time.Duration `json:"length" binding:"required"`
}

type MovieUpdate struct {
	Id       string        `json:"id" binding:"required"`
	Name     string        `json:"name"`
	Language string        `json:"language"`
	Length   time.Duration `json:"length"`
}

type Error struct {
	Message string `json:"message"`
}

type Movies struct {
	Movies []Movie `json:"movies"`
}
