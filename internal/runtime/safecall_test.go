package runtime

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSafeCall(t *testing.T) {
	// Test a function that returns a string.
	fn := func() (string, error) { return "cheese", nil }
	out, err := SafeCall(fn)
	require.NoError(t, err)
	assert.Equal(t, "cheese", out)

	// Test a function that returns a string and an error.
	fn2 := func() (string, error) { return "cheese", fmt.Errorf("oh no") }
	out, err = SafeCall(fn2)
	require.Error(t, err)
	assert.Equal(t, "cheese", out)

	// Test a function that returns a string and an error.
	fn3 := func() (string, error) { return "", fmt.Errorf("oh no") }
	out, err = SafeCall(fn3)
	require.Error(t, err)
	assert.Empty(t, out)

	// Test a function that returns a string and an error.
	fn4 := func() (string, error) { return "", nil }
	out, err = SafeCall(fn4)
	require.NoError(t, err)
	assert.Empty(t, out)

	// Test a function that returns nothing.
	fn5 := func() {}
	out, err = SafeCall(fn5)
	require.NoError(t, err)
	assert.Nil(t, out)

	// Test a function that returns 3 values.
	a, b, c := "a", "b", "c"
	fn6 := func(a, b, c string) (string, string, string) { return a, b, c }
	out, err = SafeCall(fn6, a, b, c)
	require.NoError(t, err)
	assert.Equal(t, out, a, "the return should be the first argument")

	// Test a case where the function panics.
	fn7 := func() { panic("oh no") }
	out, err = SafeCall(fn7)
	require.ErrorContains(t, err, "recovered from panic: oh no")
	assert.Nil(t, out)

	// Test when fn is not a function.
	fn8 := "cheese"
	out, err = SafeCall(fn8)
	require.ErrorContains(t, err, "fn is not a function")
	assert.Nil(t, out)

	// Test a function taking nil arguments
	fn9 := func(a, b, c any) string { return "crap" }
	out, err = SafeCall(fn9, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, "crap", out)

	// Test a variadic function taking nil arguments
	fn10 := func(a ...any) string { return "crap" }
	out, err = SafeCall(fn10, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, "crap", out)
}
