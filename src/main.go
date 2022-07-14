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

var Status map[Server]int = make(map[Server]int, 0)

func main() {
	// Status =  make(map[Server]int, 0)
	router := gin.Default()
	router.GET("/", AlphaServer)
	router.POST("/SSH", PostServerDetails )
	router.Run(":8080")
}

func AlphaServer(c *gin.Context) {

	var Response[] string
	fmt.Print(Status)
	for CurrentServer, LoginCount := range Status {
		fmt.Printf("Server : %v was logged in by %v for %v number of times \n", CurrentServer.NAME, CurrentServer.PERSON, LoginCount)
		CurrentOutput := fmt.Sprintf("Server : %v was logged in by %v for %v number of times \n", CurrentServer.NAME, CurrentServer.PERSON, LoginCount)
		Response = append(Response, string(CurrentOutput))
	}
	
	if Response != nil {
		c.IndentedJSON(http.StatusOK, Response)
	} else {
		// Response = append(Response, "OOPS !! No Data yet. Please SSH into some machines to get the data")
		c.IndentedJSON(http.StatusOK, "OOPS !! No Data yet. Please SSH into some machines to get the data")
	}
}

func PostServerDetails(c *gin.Context) {
	var NewServer Server
	c.BindJSON(&NewServer)
	Status[NewServer]++
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}