package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

// Defining the route handler for the default homepage
func AlphaServer(c *gin.Context) {

	var response bytes.Buffer
	for current_server, login_count := range status {
		current_output := fmt.Sprintf("Server : %v was logged in by %v for %v number of times\n", current_server.Server_Name, current_server.User, login_count)
		response.WriteString(current_output)
	}

	if response.String() != "" {
		c.String(http.StatusAccepted, response.String())
	} else {
		// response = append(response, "OOPS !! No Data yet. Please SSH into some machines to get the data")
		c.String(http.StatusOK, "OOPS !! No Data yet. Please SSH into some machines to get the data")
	}
}

// Defining the default handler for POST method
func PostServerDetails(c *gin.Context) {
	var new_server server

	if err := c.ShouldBindJSON(&new_server); err == nil {
		fmt.Printf("Valid Body")
		status[new_server]++
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
