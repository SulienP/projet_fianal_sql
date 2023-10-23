package sql

import (
	"database/sql"
	"fmt"
)

func getDataBase() *sql.DB {
	db, errSQLOpen := sql.Open("sqlite3", "compagny.db")
	if errSQLOpen != nil {
		fmt.Println(errSQLOpen)
	}
	defer db.Close()

	return db
}

func getAllEmployees() {
	rows, err := db.Query("SELECT id, nom, age FROM ma_table")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var id int
	var nom string
	var age int
	for rows.Next() {
		err := rows.Scan(&id, &nom, &age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, nom, age)
	}
}
