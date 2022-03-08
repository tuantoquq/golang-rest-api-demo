package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-api-mysql/database"
	"rest-api-mysql/model"

	"github.com/gorilla/mux"
)

var svResponse model.Response

//get all users from database
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var user []model.User
	err := database.Connector.Find(&user).Error
	if err != nil {
		svResponse = model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		w.WriteHeader(http.StatusOK)
		svResponse = model.Response{
			Code:    http.StatusOK,
			Message: "get all user successfully!",
			Data:    user,
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(svResponse)
}

//get user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var user model.User
	err := database.Connector.First(&user, key).Error
	if err != nil {
		svResponse = model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		svResponse = model.Response{
			Code:    http.StatusOK,
			Message: "get user by id = " + fmt.Sprint(key) + " successfully!",
			Data:    user,
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(svResponse)
}

//create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user model.User
	json.Unmarshal(requestBody, &user)
	err := database.Connector.Create(&user).Error
	if err != nil {
		svResponse = model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		svResponse = model.Response{
			Code:    http.StatusOK,
			Message: "create user successfully!",
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(svResponse)
}

//Update user by id
func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	key := vars["id"]
	var user model.User
	json.Unmarshal(requestBody, &user)
	err := database.Connector.Save(&user).Error
	if err != nil {
		svResponse = model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		svResponse = model.Response{
			Code:    http.StatusOK,
			Message: "update user by id = " + fmt.Sprint(key) + " successfully!",
			Data:    user,
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(svResponse)
}

//Delete User by id
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var user model.User

	err := database.Connector.First(&user, key).Error
	if err != nil {
		svResponse = model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		database.Connector.Where("id = ?", key).Delete(&user)
		svResponse = model.Response{
			Code:    http.StatusOK,
			Message: "Delete user by id = " + fmt.Sprint(key) + " successfully!",
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(svResponse)
}
