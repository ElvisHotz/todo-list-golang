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

func (t *Todos) BeforeCreate(db *gorm.DB) (err error) {
	if t.Finished {
		t.DateForFinished = time.Now()
	} else {
		t.DateForFinished = time.Time{}
	}
	return
}

func (t *Todos) BeforeUpdate(db *gorm.DB) (err error) {
	if t.Finished {
		t.DateForFinished = time.Now()
	} else {
		t.DateForFinished = time.Time{}
	}
	return
}
