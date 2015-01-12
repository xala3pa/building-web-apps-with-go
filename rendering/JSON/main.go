package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	Name        string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"-"`    // Field is ignored
	Country     string `json:"pais"` // Keyname changed
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	app := App{"GymGo", "Alvaro Salazar", "Track your fitness goals", "Spain"}

	content, err := json.Marshal(app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}
