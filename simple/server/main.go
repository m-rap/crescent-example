package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/m-rap/crescent"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Writing cors headers\n")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func func1(c *gin.Context) {
	crescent.SendEncRestResponse("{\"data\": \"data rahasia balasan lho\"}", c)
}

func main() {
	crescent.Init()

	route := gin.New()
	route.Use(CORSMiddleware())

	apiGroup := route.Group("/api")
	apiGroup.Use(crescent.Negotiate)
	apiGroup.POST("/func1", func1)

	//route.Static("/app", "../client")
	route.Run(":8080")
}
