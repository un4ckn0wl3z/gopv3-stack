package main

import (
	"log"
	"os"
	"testing"
	"vue-api/internal/data"

	"github.com/DATA-DOG/go-sqlmock"
)

var testApp application
var mockDB sqlmock.Sqlmock

func TestMain(m *testing.M) {
	testDB, myMock, _ := sqlmock.New()
	mockDB = myMock
	defer testDB.Close()

	testApp = application{
		config:      config{},
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog:    log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		models:      data.New(testDB),
		environment: "development",
	}

	os.Exit(m.Run())

}
