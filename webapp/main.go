package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var appID string

func main() {
	rand.Seed(time.Now().UnixNano())
	appID = fmt.Sprintf("%08x", rand.Uint32())

	router := gin.Default()
	router.GET("/convert/:fahrenheit", convertTemp)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func convertTemp(c *gin.Context) {
	var fahrenheit float64
	fmt.Sscanf(c.Param("fahrenheit"), "%f", &fahrenheit)

	celsius := (fahrenheit - 32) / 1.8

	c.JSON(http.StatusOK, gin.H{
		"celsius": celsius,
		"app_id":  appID,
	})
}
