// main_test.go
package anscombe

import (
	"math"
	"testing"
)

func TestLinearRegression(t *testing.T) {
	datasets := GetAnscombeQuartet()

	// Expected coefficients (approximately)
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

func TestRSquared(t *testing.T) {
	datasets := GetAnscombeQuartet()
	expectedR2 := 0.67 // Approximately for all datasets
	tolerance := 0.1

	for i, dataset := range datasets {
		slope, intercept := LinearRegression(dataset.X, dataset.Y)
		r2 := RSquared(dataset.X, dataset.Y, slope, intercept)

		if math.Abs(r2-expectedR2) > tolerance {
			t.Errorf("Dataset %d: Expected R² ≈ %.2f, got %.2f", i+1, expectedR2, r2)
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

func BenchmarkRSquared(b *testing.B) {
	datasets := GetAnscombeQuartet()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, dataset := range datasets {
			slope, intercept := LinearRegression(dataset.X, dataset.Y)
			RSquared(dataset.X, dataset.Y, slope, intercept)
		}
	}
}
