package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Server struct {
	PERSON string `json:"person"`
	SERVER string `json:"server"`
	IP string `json:"ip"`
}
func main() {
	router := gin.Default()
	router.GET("/", ApacheServer)
	router.Run(":8080")
}

func ApacheServer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Server{"Yuvi", "Server_A", "192.168.0.1"})
	fmt.Print("Hi ! I'm running inside the server")
}