package main

import (
	"log"
	"net/http"
	"rest-api-mysql/controllers"
	"rest-api-mysql/database"
	"rest-api-mysql/model"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Process running in port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initializeHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initDB(){
	databaseConfig := database.Config{
		ServerName: "localhost:3306",
		User: "tuannha",
		Password: "admin@123",
		DB: "atrobotics",
	}

	connectionString := database.GetConnectionString(databaseConfig)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&model.User{})
}

func initializeHandlers(router *mux.Router){
	router.PathPrefix("/api/user")
	router.HandleFunc("/getAll", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/create", controllers.CreateUser).Methods("POST")
}