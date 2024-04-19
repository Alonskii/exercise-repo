package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	// Create a new HTTP request with method GET
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function directly and pass in the request and response recorder
	HelloServer(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHelloServerWithDifferentMethods(t *testing.T) {
	testMethods := []struct {
		method string
		status int
	}{
		{"GET", http.StatusOK},
		{"POST", http.StatusMethodNotAllowed},
		{"PUT", http.StatusMethodNotAllowed},
		{"DELETE", http.StatusMethodNotAllowed},
	}

	// Loop through each test method
	for _, tm := range testMethods {
		// Create a new HTTP request with the current method
		req, err := http.NewRequest(tm.method, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the handler function directly and pass in the request and response recorder
		HelloServer(rr, req)

		// Check the response status code
		if status := rr.Code; status != tm.status {
			t.Errorf("handler returned wrong status code for method %s: got %v want %v",
				tm.method, status, tm.status)
		}
	}
}
