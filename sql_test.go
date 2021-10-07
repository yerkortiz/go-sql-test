package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// a successful case
func TestInsertUserPass(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	fmt.Println(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO user").WithArgs("bolas", "secas").WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = InsertUser(db, "user", "123")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertUserFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO user").WithArgs("bolas", "secas").WillReturnError(errors.New("XD"))

	_, err = InsertUser(db, "user", "123")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
