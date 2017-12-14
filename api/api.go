package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teamNoble/uniqueId/backend"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/id", func(c *gin.Context) {
		c.String(200, backend.Prase())
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":6666")
}