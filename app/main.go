// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	v1 := router.Group("api/v1/calendar") {
		
	}
}
