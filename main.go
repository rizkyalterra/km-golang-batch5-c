package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id int 			`json:"id"`
	Name string 	`json:"name"`
	Email string 	`json:"email"`
}

type BaseResponse struct {
	Status bool 		`json:"status"`
	Message string   	`json:"message"`
	Data interface{} 	`json:"data"`
}

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("Sekarang server REST API dijalankan di port :8000")
	http.ListenAndServe(":8000", nil)
}

func generateDataUser() []User {
	result := []User{}
	result = append(result, User{1, "A", "A"})
	result = append(result, User{2, "B", "B"})
	result = append(result, User{3, "C", "C"})
	return result
}

func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		response := BaseResponse {
			Status: true,
			Message: "Success Get Data",
			Data: generateDataUser(),
		}
		resultJson, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(resultJson)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}