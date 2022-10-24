package main

import "github.com/gin-gonic/gin"

func GetDataD(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SUCCESS",
	})
}

func main() {
	r := gin.Default()

	r.GET("/api", GetDataD)

	r.Run(":3001")
}
