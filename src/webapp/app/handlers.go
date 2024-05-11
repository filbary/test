package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CelsiusRequest struct {
	Fahrenheit float64 `json:"fahrenheit"`
}

func ConvertTemp(c *gin.Context) {
	var req CelsiusRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	celsius := CalculateCelsius(req.Fahrenheit)

	c.JSON(http.StatusOK, gin.H{
		"celsius": celsius,
		"app_id":  appID,
	})
}
