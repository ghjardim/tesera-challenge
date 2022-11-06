package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

	// API routes
	router := mux.NewRouter()
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomerById).Methods("GET")
	router.HandleFunc("/customers", createNewCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", alterCustomer).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// API Controllers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	var customer []Customer
	db.Find(&customer)
	json.NewEncoder(w).Enconde(&customer)
}
func getCustomerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer Customer
	db.First(&customer, params["id"])
	json.NewEncoder(w).Enconde(customer)
}
func createNewCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)

	createdCustomer := db.Create(&customer)
	err := createdCustomer.Error
	if  err != nil {
		json.NewEncoder(w).Encode(err)
	} else  {
		json.NewEncoder(w).Encode(&person)
	}
}
func alterCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer Customer

}
