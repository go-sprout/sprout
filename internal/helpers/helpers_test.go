package helpers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStringer struct{}

func (ts testStringer) String() string {
	return "stringer"
}

func TestToString(t *testing.T) {
	assert.Equal(t, "string", ToString("string"))
	assert.Equal(t, "5", ToString(5))
	assert.Equal(t, "true", ToString(true))
	assert.Equal(t, "bytes", ToString([]byte("bytes")))
	assert.Equal(t, "error", ToString(errors.New("error")))
	assert.Equal(t, "stringer", ToString(testStringer{}))
}

func TestEmpty(t *testing.T) {
	// Testing for nil pointers
	assert.True(t, Empty(nil))

	// Testing for zero-length collections
	assert.True(t, Empty([]int{}))
	assert.True(t, Empty(map[string]int{}))
	assert.True(t, Empty(""))

	// Testing for zero values of numeric types
	assert.True(t, Empty(0))
	assert.True(t, Empty(0.0))
	assert.True(t, Empty(complex(0, 0)))
	assert.True(t, Empty(uint(0)))
	assert.True(t, Empty(uint8(0)))
	assert.True(t, Empty(uint16(0)))
	assert.True(t, Empty(uint32(0)))
	assert.True(t, Empty(uint64(0)))
	assert.True(t, Empty(uintptr(0)))

	// Testing for false booleans
	assert.True(t, Empty(false))

	// Testing for non-empty structs
	assert.False(t, Empty(struct{}{}))

	// Testing for other types
	assert.True(t, Empty([]byte{}))
	assert.False(t, Empty(func() {}))

	// Testing for non-empty strings
	assert.False(t, Empty("string"))

	// Testing for non-empty values
	assert.False(t, Empty(map[string]string{"key": "value"}))
	assert.False(t, Empty([]string{"string"}))
	assert.False(t, Empty(" "))
	assert.False(t, Empty(testStringer{}))
	assert.False(t, Empty(func() {}))
	assert.False(t, Empty(&struct{}{}))
	assert.False(t, Empty(errors.New("error")))
	assert.False(t, Empty(1))
	assert.False(t, Empty(1.0))
	assert.False(t, Empty(1i))
	assert.False(t, Empty(true))
	assert.False(t, Empty(uint(1)))
	assert.False(t, Empty(uint8(1)))

}

func TestUntilStep(t *testing.T) {
	assert.Equal(t, []int{0, 2, 4, 6, 8}, UntilStep(0, 10, 2), "UntilStep: positive step")
	assert.Equal(t, []int{10, 8, 6, 4, 2}, UntilStep(10, 0, -2), "UntilStep: negative step")
	assert.Empty(t, UntilStep(0, 10, 0), "UntilStep: invalid step")
	assert.Empty(t, UntilStep(10, 0, 2), "UntilStep: start > stop and positive step")
}

func TestStrSlice(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  []string
	}{
		{"nil input", nil, []string{}},
		{"[]string input", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"[]any input", []any{"a", 1, true}, []string{"a", "1", "true"}},
		{"[]int input", []int{1, 2, 3}, []string{"1", "2", "3"}},
		{"[]float64 input", []float64{1.2, 3.4, 5.6}, []string{"1.2", "3.4", "5.6"}},
		{"non-slice input", "hello", []string{"hello"}},
		{"empty slice input", []string{}, []string{}},
		{"float64 input with integer value", 7., []string{"7"}},
		{"float64 input with decimal value", .8, []string{"0.8"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrSlice(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
