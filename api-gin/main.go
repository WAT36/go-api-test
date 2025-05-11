// main.go
package main

import "github.com/gin-gonic/gin"

type response struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default() // ロギング + リカバリ済み
	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "world")
		c.JSON(200, response{Message: "Hello, " + name + "!"})
	})

	r.Run(":8080") // デフォルトは 0.0.0.0:8080
}
