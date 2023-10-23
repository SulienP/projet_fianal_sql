package sql

import (
	"database/sql"
	"fmt"
)

func getDataBase() *sql.DB {
	db, errSQLOpen := sql.Open("sqlite3", "public/compagny.db")
	if errSQLOpen != nil {
		fmt.Println(errSQLOpen)
	}

	return db
}

func getAllEmployees() {
	db := getDataBase()

	defer db.Close()

	rows, err := db.Query("SELECT email FROM employees")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var email string
	for rows.Next() {
		err := rows.Scan(&email)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(email)
	}
}
