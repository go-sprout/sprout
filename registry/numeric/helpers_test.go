package numeric

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCleanFloatPrecision(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		// Floating-point precision issue cases
		{"0.1+0.2 noise", 0.1 + 0.2, 0.3},
		{"0.7+0.1 noise", 0.7 + 0.1, 0.8},
		{"1.1+2.2 noise", 1.1 + 2.2, 3.3},
		// Edge cases
		{"zero", 0.0, 0.0},
		{"negative zero", math.Copysign(0, -1), 0.0},
		{"positive infinity", math.Inf(1), math.Inf(1)},
		{"negative infinity", math.Inf(-1), math.Inf(-1)},
		// Normal values should stay unchanged
		{"integer one", 1.0, 1.0},
		{"half", 0.5, 0.5},
		{"negative", -5.5, -5.5},
		{"pi approximation", 3.14159, 3.14159},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := cleanFloatPrecision(tc.input)
			if math.IsNaN(tc.expected) {
				assert.True(t, math.IsNaN(result))
			} else {
				assert.InDelta(t, tc.expected, result, 0.01)
			}
		})
	}

	// Special test for NaN (can't compare with ==)
	t.Run("NaN", func(t *testing.T) {
		result := cleanFloatPrecision(math.NaN())
		assert.True(t, math.IsNaN(result))
	})
}

func TestOperateNumeric(t *testing.T) {
	tests := []struct {
		values   []any
		op       numericOperation
		initial  any
		expected any
	}{
		{[]any{1, 2}, func(a, b float64) float64 { return a + b }, 0, 3},
		{[]any{1.5, 2.5}, func(a, b float64) float64 { return a - b }, 0, -1.0},
		{[]any{1.5, 2.5}, func(a, b float64) float64 { return a * b }, 1, 3.75},
		{[]any{1.5, 2.5}, func(a, b float64) float64 { return a / b }, 1, 0.6},
		{[]any{uint(1), uint(2)}, func(a, b float64) float64 { return a + b }, uint(0), uint(3)},
		// Test floating-point precision fix
		{[]any{0.1, 0.2}, func(a, b float64) float64 { return a + b }, 0.0, 0.3},
		// Test subtraction with zero and negative
		{[]any{0.0, -100.5}, func(a, b float64) float64 { return a - b }, 0.0, 100.5},
	}

	for _, test := range tests {
		result, err := operateNumeric(test.values, test.op, test.initial)
		require.NoError(t, err)
		assert.Equal(t, test.expected, result, "operateNumeric(%v, %v, %v) returned %v, expected %v", test.values, test.op, test.initial, result, test.expected)
	}
}
