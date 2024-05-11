package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var appID string

func main() {
	rand.Seed(time.Now().UnixNano())
	appID = fmt.Sprintf("%08x", rand.Uint32())

	router := gin.Default()
	router.POST("/convert", ConvertTemp)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}
