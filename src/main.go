package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Server struct {
	PERSON string `json:"person"`
	NAME string `json:"name"`
	IP string `json:"ip"`
}

var Status[] Server
func main() {
	router := gin.Default()
	router.GET("/", ApacheServer)
	router.POST("/SSH", PostServerDetails )
	router.Run(":8080")
}

func ApacheServer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Status)
	fmt.Print("Hi ! I'm running inside the server")
}

func PostServerDetails(c *gin.Context) {
	var NewServer Server
	c.BindJSON(&NewServer)
	Status = append(Status, NewServer)
	q := url.Values{}
	location := url.URL{Path: "/", RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())
}