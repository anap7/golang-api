package controllers

import "net/http"

// Create a new user on Database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User successfully created!"))
}

// It'll search for a user on database and return all information about it
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User successfully found!"))
}

// It'll get all users from database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User list successfully found"))
}

// It'll update some users information on database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User successfully updated"))
}

// It'll delete a user on database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User successfully deleted"))
}