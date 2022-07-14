package main

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/stretchr/testify"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}
func TestDefaultPage(t *testing.T) {
    mockResponse := "OOPS !! No Data yet. Please SSH into some machines to get the data"
    r := SetUpRouter()
    r.GET("/", AlphaServer)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    responseData, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostMethod(t *testing.T) {
	r := SetUpRouter()
    r.POST("/SSH", PostServerDetails)
	
	newServer := Server{
		PERSON: "Person A",
		NAME: "Server A",
		IP: "192.168.0.1",
	}

	jsonBody, _ := json.Marshal(newServer)
	// fmt.Print(jsonBody)
	req, _ := http.NewRequest("POST", "/SSH", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusFound, w.Code)

// func TestBadRequest(t *testing.T) {
// 	r := SetUpRouter()
// 	r.POST("/SSH", PostServerDetails)
	
// 	newServer := Server{
// 		PERSON: "Person A",
// 		NAME: "Server A",
// 		IP: "192.168.0.1",
// 	}

// 	jsonBody, _ := json.Marshal(newServer)
// 	// fmt.Print(jsonBody)
// 	req, _ := http.NewRequest("POST", "/SSH", bytes.NewBuffer(jsonBody))
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusFound, w.Code)
}