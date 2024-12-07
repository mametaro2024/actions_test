package main

import (
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler := HelloHandler{}

	handler.ServeHTTP(rec, req)

	expected := "Hello world!"
	result := rec.Body.String()

	if result != expected {
		t.Errorf("expected result is %v but actually %v", expected, result)
	}

}
