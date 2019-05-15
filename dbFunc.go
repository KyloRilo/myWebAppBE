package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

type userObj struct {
	gorm.Model
	LoginName    string
	PasswordHash string
	FirstName    string
	LastName     string
	email        string
}
type dbInfo struct {
	Database string
	Host     string
	Port     string
	User     string
	Pass     string
}

func GetAllUsers() interface{} {
	dbConnect := dbInfo{
		os.Getenv("DB"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
	}
	var users userObj
	db, err := gorm.Open("mssql", "sqlserver://%s:%s@%s:%s?database=%s",
		dbConnect.User,
		dbConnect.Pass,
		dbConnect.Host,
		dbConnect.Port,
		dbConnect.Database,
	)

	if err != nil {
		log.Fatalf("Failed to connnect to database")
	}
	r := db.Find(&users)
	rJSON, _ := json.Marshal(r)
	defer db.Close()
	return rJSON
}
