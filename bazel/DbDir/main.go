package main

import (
	"db/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	_ "log"
)

type Student struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func main() {
	db, err := database.GetDBClient()
	if err != nil {
		panic(err)
	}
	database.CreateTables(db)

	students := []Student{
		{
			FirstName: "Salman",
			LastName:  "Ansari",
			Email:     "salman@gmail.com",
		}, {
			FirstName: "Some",
			LastName:  "LastName",
			Email:     "abc@xyz.com",
		},
		{
			FirstName: "MyName",
			LastName:  "Ansari",
			Email:     "myname@yourname.com",
		},
	}
	q := `INSERT INTO student_v3 (first_name,last_name,email) VALUES (:first_name,:last_name,:email)`
	_, err = db.NamedExec(q, students)
	if err != nil {
		fmt.Println("Error", err)
	}
}
