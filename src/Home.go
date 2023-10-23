package sql

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import du pilote SQLite
)

type homePage struct {
	Hello string
}

func Home(w http.ResponseWriter, r *http.Request) {
tmpl := template.Must(template.ParseFiles("templates/Home.html"))

	data := homePage{
		Hello: "Hello, World!", // Exemple de données pour votre modèle
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
