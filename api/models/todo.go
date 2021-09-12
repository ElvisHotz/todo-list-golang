package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Todos struct {
	gorm.Model
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	DateForFinished time.Time `json:"dateForFinished"`
	Finished        bool      `json:"finished"`
}

func (t *Todos) TableName() string {
	return "todos"
}
