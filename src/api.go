package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func AlphaServer(c *gin.Context) {

	var Response bytes.Buffer
	for CurrentServer, LoginCount := range Status {
		CurrentOutput := fmt.Sprintf("Server : %v was logged in by %v for %v number of times\n", CurrentServer.ServerName, CurrentServer.User, LoginCount)
		Response.WriteString(CurrentOutput)
	}

	if Response.String() != "" {
		c.String(http.StatusAccepted, Response.String())
	} else {
		// Response = append(Response, "OOPS !! No Data yet. Please SSH into some machines to get the data")
		c.String(http.StatusOK, "OOPS !! No Data yet. Please SSH into some machines to get the data")
	}
}

func PostServerDetails(c *gin.Context) {
	var NewServer Server

	if err := c.ShouldBindJSON(&NewServer); err == nil {
		fmt.Printf("Valid Body")
		Status[NewServer]++
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
