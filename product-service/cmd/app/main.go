package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id          int64  `json:"id"`
	ProductName string `json:"name"`
	ProductCode string `json:"code"`
}

func main() {
	http.HandleFunc("/product", handleProduct)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	data := Product{Id: 101, ProductName: "Shampoo", ProductCode: "CDL-123"}

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
