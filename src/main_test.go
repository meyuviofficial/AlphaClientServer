package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	// "github.com/stretchr/testify"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}
func TestHomepageHandler(t *testing.T) {
    mockResponse := `"OOPS !! No Data yet. Please SSH into some machines to get the data"`
    r := SetUpRouter()
    r.GET("/", AlphaServer)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    responseData, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
}