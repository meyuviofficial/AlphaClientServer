package main

import (
	"github.com/gin-gonic/gin"
)

type server struct {
	User        string `json:"user" binding:"required"`
	Server_Name string `json:"servername" binding:"required"`
	Ip          string `json:"ip" binding:"required"`
}

var status map[server]int = make(map[server]int, 0)

func main() {
	status = make(map[server]int, 0)
	router := gin.Default()
	router.GET("/", AlphaServer)
	router.POST("/SSH", PostServerDetails)
	router.Run(":8080")
}
