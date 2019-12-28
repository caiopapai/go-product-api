package connection

import (
	"database/sql"
	"fmt"

	//Postgrees driver
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Pooble@01"
	dbname   = "pooble"
)

//GetConnection creates a new Database connection
func GetConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Successfully connected!")
	return db
}

//CloseConnection close the Database connection
func CloseConnection(db *sql.DB) {
	defer db.Close()
}
