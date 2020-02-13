package test_web_frame

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

//https://github.com/gin-gonic/gin

func ping(c *gin.Context) {
	fmt.Println("--- full path:", c.FullPath()) // --- full path: /ping
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ginHello(c *gin.Context) {
	//c.Request.Body.Read()
}

func Test_gin01(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", ping)
	router.POST("/gin_hello", ginHello)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
