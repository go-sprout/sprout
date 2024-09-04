package numeric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	}

	for _, test := range tests {
		result, err := operateNumeric(test.values, test.op, test.initial)
		require.NoError(t, err)
		assert.Equal(t, test.expected, result, "operateNumeric(%v, %v, %v) returned %v, expected %v", test.values, test.op, test.initial, result, test.expected)
	}
}
