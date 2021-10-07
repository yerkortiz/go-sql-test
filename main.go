package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE TABLE user(
	id int NOT NULL AUTO_INCREMENT,
	name varchar(32) NOT NULL,
	email varchar(32) NOT NULL,
	PRIMARY KEY(id)
)
*/
const insertUser = `
	INSERT INTO user(name, email)
	VALUES (?, ?)
	`

func InsertUser(db *sql.DB, name string, email string) (lastInsertID int64, err error) {
	result, err := db.Exec(insertUser, name, email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/test1")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	InsertUser(db, "user", "123")

}
