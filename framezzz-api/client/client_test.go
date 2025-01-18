package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestServer(statusCode int, responseBody string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, responseBody)
	}))
}

func TestMakeAPICall_Success(t *testing.T) {
	server := setupTestServer(http.StatusOK, "Success")
	defer server.Close()

	response, err := MakeAPICall("GET", server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response != "Success\n" {
		t.Errorf("Expected response 'Success', got %v", response)
	}
}

func TestMakeAPICall_BadRequest(t *testing.T) {
	server := setupTestServer(http.StatusBadRequest, "Bad Request")
	defer server.Close()

	_, err := MakeAPICall("GET", server.URL)
	if err == nil {
		t.Fatal("Expected error, got none")
	}
	if err.Error() != "Bad Request: 400\nBody: Bad Request\n" {
		t.Errorf("Expected 400 error, got %v", err)
	}
}

func TestMakeAPICall_NotFound(t *testing.T) {
	server := setupTestServer(http.StatusNotFound, "Not Found")
	defer server.Close()

	_, err := MakeAPICall("GET", server.URL)
	if err == nil {
		t.Fatal("Expected error, got none")
	}
	if err.Error() != "Not Found: 404\nBody: Not Found\n" {
		t.Errorf("Expected 404 error, got %v", err)
	}
}

func TestMakeAPICall_TooManyRequests(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintln(w, "Too Many Requests")
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Success")
		}
	}))
	defer server.Close()

	response, err := MakeAPICall("GET", server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response != "Success\n" {
		t.Errorf("Expected response 'Success', got %v", response)
	}
	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

func TestMakeAPICall_InternalServerError(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Internal Server Error")
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Success")
		}
	}))
	defer server.Close()

	response, err := MakeAPICall("GET", server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response != "Success\n" {
		t.Errorf("Expected response 'Success', got %v", response)
	}
	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}
