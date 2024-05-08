package main

import (
	"testing"
)

func TestCalculateCelcius(t *testing.T) {
	tests := []struct {
		name       string
		fahrenheit float64
		want       float64
	}{
		{"zero point", 32, 0},
		{"freezing point", 212, 100},
		{"subzero", -40, -40},
		{"random value", 77, (77 - 32) / 1.8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateCelcius(tt.fahrenheit); got != tt.want {
				t.Errorf("CalculateCelcius(%f) = %f; want %f", tt.fahrenheit, got, tt.want)
			}
		})
	}
}
