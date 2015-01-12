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

var tmpl = template.Must(template.ParseFiles(path.Join("templates", "index.html")))

func main() {
	http.HandleFunc("/", ShowApps)
	http.ListenAndServe(":8080", nil)
}

func ShowApps(w http.ResponseWriter, r *http.Request) {
	app := App{"Gymgo", "Álvaro Salazar", "Track your fitness goals"}

	if err := tmpl.Execute(w, app); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
