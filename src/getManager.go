package sql

import (
    "log"
)

func getManagers(id int) string {
    db := getDataBase()
    defer db.Close()

    query := `
        SELECT
            CASE
                WHEN employees.employeeId < 8 THEN 'Manager 1'
                WHEN employees.employeeId BETWEEN 10 AND 12 THEN 'Manager 2'
                ELSE 'Manager 3'
            END AS Manager
        FROM employees
        WHERE employeeId = ?;
    `

    var manager string

    err := db.QueryRow(query, id).Scan(&manager)
    if err != nil {
        log.Fatal(err)
    }

    return manager
}
