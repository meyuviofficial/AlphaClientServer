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

var Status map[Server]int

func main() {
	Status =  make(map[Server]int, 0)
	router := gin.Default()
	router.GET("/", ApacheServer)
	router.POST("/SSH", PostServerDetails )
	router.Run(":8080")
}

func ApacheServer(c *gin.Context) {

	var Response[] string
	fmt.Print(Status)
	for CurrentServer, LoginCount := range Status {
		fmt.Printf("Server : %v was logged in by %v for %v number of times \n", CurrentServer.NAME, CurrentServer.PERSON, LoginCount)
		CurrentOutput := fmt.Sprintf("Server : %v was logged in by %v for %v number of times \n", CurrentServer.NAME, CurrentServer.PERSON, LoginCount)
		Response = append(Response, CurrentOutput)
	}
	c.IndentedJSON(http.StatusOK, Response)
}

func PostServerDetails(c *gin.Context) {
	var NewServer Server
	c.BindJSON(&NewServer)
	Status[NewServer]++

	
	// if Status == nil {
	// 	Status = map[Server]int{NewServer: 1}
	// }
	// // fmt.Print(Status[NewServer], "is the value of login attempts")
	q := url.Values{}
	location := url.URL{Path: "/", RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())
}