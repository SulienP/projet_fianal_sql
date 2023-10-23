package sql

import (
	"database/sql"
	"fmt"
	"log"
)

func getDataBase() *sql.DB {
	db, errSQLOpen := sql.Open("sqlite3", "public/compagny.db")
	if errSQLOpen != nil {
		fmt.Println(errSQLOpen)
	}

	return db
}

func getEmails() []string {
	db := getDataBase()

	rows, errQuery := db.Query("SELECT email FROM employees")
	if errQuery != nil {
		log.Fatalln(errQuery)
	}

	defer rows.Close() // Assurez-vous de fermer les lignes après les avoir utilisées
	var emails []string

	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			log.Fatal(err)
		}
		emails = append(emails, email) // Ajoutez chaque email à la liste des emails
	}

	return emails
}
