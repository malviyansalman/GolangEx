package Models

import "fmt"

var StudentSchema = `create table if not exists Student(
	first_name text,
	last_name text,
	email text,
	version int
	);`

var TeacherSchema = `create table if not exists Teacher(
	first_name text,
	last_name text,
	email text,
	version int
	);`

func CreateStudentTable(version string) string {
	var StudentSchema = `create table if not exists Student_%s(
	first_name text,
	last_name text,
	email text,
	version int
	);`
	StudentSchema = fmt.Sprintf(StudentSchema, version)
	return StudentSchema
}
