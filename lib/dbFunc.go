package lib

import (
	"encoding/json"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
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
	/*
		dbConnect := dbInfo{
			os.Getenv("DB"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
		}*/
	var users userObj
	db, err := gorm.Open("mssql", "")
	defer db.Close()

	if err != nil {
		log.Fatalf("Failed to connnect to database: %s", err)
	}
	r := db.First(&users)
	rJSON, _ := json.Marshal(r)
	n := len(rJSON)
	obj := string(rJSON[:n])
	defer db.Close()
	return obj
}
