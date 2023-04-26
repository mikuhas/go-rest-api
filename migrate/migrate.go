package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConnection := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConnection)
	dbConnection.AutoMigrate(&model.User{}, &model.Task{})
}
