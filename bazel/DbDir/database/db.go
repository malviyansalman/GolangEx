package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func GetDBClient() (*sqlx.DB, error) {
	userName := "root"
	password := "Emancipation@1"
	host := "127.0.0.1"
	port := 3306
	dbName := "student"
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", userName, password, host, port, dbName)
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return db, nil
}
