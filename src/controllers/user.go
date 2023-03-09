package controllers

import "net/http"

// Create a new user on Database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User successfully created!"))
}

// Search for a user on database and return all information about it
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User successfully found!"))
}

// Get all users from database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User list successfully found"))
}

// Update some users information on database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User successfully updated"))
}

// Delete a user on database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User successfully deleted"))
}