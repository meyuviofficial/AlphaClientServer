package main

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	User       string `json:"user" binding:"required"`
	ServerName string `json:"servername" binding:"required"`
	Ip         string `json:"ip" binding:"required"`
}

var Status map[Server]int = make(map[Server]int, 0)

func main() {
	Status = make(map[Server]int, 0)
	router := gin.Default()
	router.GET("/", AlphaServer)
	router.POST("/SSH", PostServerDetails)
	router.Run(":8080")
}
