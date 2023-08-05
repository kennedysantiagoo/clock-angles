package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
)

func GetConnection() (*sql.DB, error) {
	psqlconn := getConnectionInfo()
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Println("Error to open connection ", err)
		return nil, err
	}

	return db, err
}

func getConnectionInfo() string {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)
	return connInfo
}
