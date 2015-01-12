package main

import (
	"html/template"
	"net/http"
	"path"
)

type App struct {
	Name        string
	Author      string
	Description string
}

func main() {
	http.HandleFunc("/", ShowApps)
	http.ListenAndServe(":8080", nil)
}

func ShowApps(w http.ResponseWriter, r *http.Request) {
	app := App{"Gymgo", "Álvaro Salazar", "Track your fitness goals"}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, app); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
