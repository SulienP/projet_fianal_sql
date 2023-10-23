// package requetSql

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func recuperationUser() {
// 	database, _ := sql.Open("sqlite3", "../../public/compagny.db")
// 	defer database.Close()
// 	rows, err := database.Query("SELECT * FROM EMPLOYEES")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var employeeId int
// 		var postId string
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var password string
// 		var isPresent bool
// 		var salary int
// 		var schedule string
// 		var breaksTimes string
// 		var dateHire string
// 		var endContact string
// 		err = rows.Scan(&id, &username, &password, &profilDescription, &mail, &xp)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		userIntoStruc := recuperationUserFromDb{}
// 		userIntoStruc = recuperationUserFromDb{
// 			Id:                id,
// 			Username:          username,
// 			ProfilDescription: profilDescription,
// 		}
// 		allUsers = append(allUsers, userIntoStruc)

// 	}
// 	return allUsers
// }
package sql
func test(){

}