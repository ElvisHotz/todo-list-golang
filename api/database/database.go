package database

import (
	"fmt"
	"todo-list/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	return db
}

func connect() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", config.ConnectionString)

	if err != nil {
		fmt.Println("Status: ", err)
	}

	if err = db.DB().Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
