package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConvertTemp(c *gin.Context) {
	var fahrenheit float64
	fmt.Sscanf(c.Param("fahrenheit"), "%f", &fahrenheit)

	celsius := CalculateCelcius(fahrenheit)

	c.JSON(http.StatusOK, gin.H{
		"celsius": celsius,
		"app_id":  appID,
	})
}
