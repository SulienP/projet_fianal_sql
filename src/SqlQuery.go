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

	return db
}
