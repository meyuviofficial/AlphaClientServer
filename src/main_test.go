package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

function TestHandler(t, *testing.T) {
	req, err := http.NewRequest("Get", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)

	
}

