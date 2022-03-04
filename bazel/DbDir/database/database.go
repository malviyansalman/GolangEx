package database

import (
	"db/Models"
	"github.com/jmoiron/sqlx"
)

func CreateTables(db *sqlx.DB) {

	student := Models.CreateStudentTable("v3")
	//ss := fmt.Sprintf(student, "v2")
	db.MustExec(student)
}
