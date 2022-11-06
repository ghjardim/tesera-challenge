package main

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Customer struct {
	gorm.Model
	Id int `gorm:"unique_index"`
	FirstName string
	LastName string
	Email string
	Password string
}

var db *gorm.DB
var err error

func main(){
	// Loading environment variables
	dialect :=  os.Getenv("DIALECT")
	host :=  os.Getenv("HOST")
	dbPort :=  os.Getenv("DBPORT")
	user :=  os.Getenv("USER")
	dbName :=  os.Getenv("DBNAME")
	password :=  os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=% password=%s port=%s sslmode=disable", host, user, dbName, password, dbPort)

	// Opening connection to database
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to db!")
	}

	// Set to close connection to db when main finishes
	defer db.Close()

	// Make migrations to the db if they have not already been created
	db.AutoMigrate(&Customer{})
}
