package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func main() {
	http.HandleFunc("/task", taskHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req requestBody

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
			return
		}

		task = req.Task

		fmt.Fprintf(w, "Задача сохранена: %s", task)
	} else if r.Method == http.MethodGet {

		fmt.Fprintf(w, "hello, %s", task)
	} else {
		
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
