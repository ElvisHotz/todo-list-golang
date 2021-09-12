package database

import (
	"todo-list/models"

	"github.com/jinzhu/gorm"
)

func IniTables(db *gorm.DB) {
	db.AutoMigrate(&models.Todos{})
}
