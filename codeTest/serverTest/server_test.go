package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMyhandler(t *testing.T){
	requestBody := `{"name":"rainbowbear"}`
	
	req := httptest.NewRequest("GET","/test", strings.NewReader(requestBody))
	req.Header.Set("Content_type", "application/json")
	
	rr := httptest.NewRecorder()

	testHandler(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
    }

	expectedResponseBody := "rainbowbear님 반갑습니다."
    if rr.Body.String() != expectedResponseBody {
        t.Errorf("Expected body %s, but got %s", expectedResponseBody, rr.Body.String())
    }
}