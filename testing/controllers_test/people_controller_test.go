package main

import (
	"fmt"
	"golang-rest-api-best-practice/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSavePerson(t *testing.T) {
	req, _ := http.NewRequest("POST", "/people", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.SavePerson)
	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Code)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}