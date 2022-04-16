package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestApplication_AllUsers(t *testing.T) {
	// create mock rows
	var mockedRows = mockDB.NewRows([]string{"id", "email", "first_name", "last_name", "password", "active", "created_at", "updated_at", "has_token"})
	mockedRows.AddRow("1", "me@there.com", "Anuwat", "Kh.", "xxxxxx", "1", time.Now(), time.Now(), "0")

	// tell mock
	mockDB.ExpectQuery("select \\\\* ").WillReturnRows(mockedRows)

	// create test recorder
	rr := httptest.NewRecorder()
	// create req
	req, _ := http.NewRequest("POST", "/admin/users", nil)
	// call the handler
	handler := http.HandlerFunc(testApp.AllUsers)
	handler.ServeHTTP(rr, req)

	// check expected
	if rr.Code != http.StatusOK {
		t.Error("AllUsers returned wrong status code of", rr.Code)
	}

}
