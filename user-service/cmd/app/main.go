package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func main() {
	http.HandleFunc("/user", handleUser)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	data := User{Id: 101, Name: "Trevor", Gender: "male"}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), 405)
	}
	fmt.Println("handle health endpoint triggered")
	w.Header().Set("Content-Type", "application/json")

	//w.WriteHeader(http.StatusCreated)
	// writeResponse(w, "Ok !")
	err := json.NewEncoder(w).Encode(data)
	fmt.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), 400)
	}

}
