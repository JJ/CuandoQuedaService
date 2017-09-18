package main

// taken from https://github.com/tucnak/telebot
import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {

	req, err := http.NewRequest("GET", "/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(dame_hitos)
	recorder := httptest.NewRecorder()

	// Taken from https://elithrar.github.io/article/testing-http-handlers-go/
	handler.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"file":"0.Repositorio","title":"Git y GitHub","fecha":"21/09/2017"}`
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

