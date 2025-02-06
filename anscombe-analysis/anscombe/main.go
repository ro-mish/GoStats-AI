package main

import (
	"fmt"
	"math"
	"testing"
)

// AnscombeData represents one of the Anscombe quartet datasets
type AnscombeData struct {
	X []float64
	Y []float64
}

// GetAnscombeQuartet returns all four Anscombe datasets
func GetAnscombeQuartet() []AnscombeData {
	return []AnscombeData{
		{ // Dataset I
			X: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		},
		{ // Dataset II
			X: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
		},
		{ // Dataset III
			X: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		},
		{ // Dataset IV
			X: []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
			Y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.50, 5.56, 7.91, 6.89},
		},
	}
}

// LinearRegression calculates the coefficients for y = mx + b
func LinearRegression(x, y []float64) (m, b float64) {
	n := float64(len(x))

	sumX, sumY := 0.0, 0.0
	for i := range x {
		sumX += x[i]
		sumY += y[i]
	}
	meanX := sumX / n
	meanY := sumY / n

	numerator, denominator := 0.0, 0.0
	for i := range x {
		numerator += (x[i] - meanX) * (y[i] - meanY)
		denominator += (x[i] - meanX) * (x[i] - meanX)
	}

	m = numerator / denominator
	b = meanY - m*meanX
	return
}

// RSquared calculates the coefficient of determination
func RSquared(x, y []float64, m, b float64) float64 {
	meanY := 0.0
	for _, yi := range y {
		meanY += yi
	}
	meanY /= float64(len(y))

	totalSS, residualSS := 0.0, 0.0
	for i := range y {
		predicted := m*x[i] + b
		residualSS += math.Pow(y[i]-predicted, 2)
		totalSS += math.Pow(y[i]-meanY, 2)
	}

	return 1 - (residualSS / totalSS)
}

func main() {
	datasets := GetAnscombeQuartet()

	for i, dataset := range datasets {
		slope, intercept := LinearRegression(dataset.X, dataset.Y)
		r2 := RSquared(dataset.X, dataset.Y, slope, intercept)

		fmt.Printf("Dataset %d:\n", i+1)
		fmt.Printf("Slope: %.4f\n", slope)
		fmt.Printf("Intercept: %.4f\n", intercept)
		fmt.Printf("R²: %.4f\n\n", r2)
	}
}

// Test functions
func TestLinearRegression(t *testing.T) {
	datasets := GetAnscombeQuartet()
	expectedSlope := 0.5
	expectedIntercept := 3.0
	tolerance := 0.1

	for i, dataset := range datasets {
		slope, intercept := LinearRegression(dataset.X, dataset.Y)

		if math.Abs(slope-expectedSlope) > tolerance {
			t.Errorf("Dataset %d: Expected slope ≈ %.2f, got %.2f", i+1, expectedSlope, slope)
		}
		if math.Abs(intercept-expectedIntercept) > tolerance {
			t.Errorf("Dataset %d: Expected intercept ≈ %.2f, got %.2f", i+1, expectedIntercept, intercept)
		}
	}
}

func BenchmarkLinearRegression(b *testing.B) {
	datasets := GetAnscombeQuartet()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, dataset := range datasets {
			LinearRegression(dataset.X, dataset.Y)
		}
	}
}
