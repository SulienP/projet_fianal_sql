package sql

import (
	"log"
)

func getDepartementNames(employeeID int) (string) {
	db := getDataBase()
	defer db.Close()

	query := `SELECT posts.postName, departements.serviceName FROM posts, departements  INNER JOIN employees ON posts.postId = employees.postId WHERE employees.employeeId = ?`

	var postName, serviceName string

	err := db.QueryRow(query, employeeID).Scan(&postName, &serviceName)
	if err != nil {
		log.Printf("Erreur lors de la récupération des données de département : %v", err)
		
	}

	return postName
}