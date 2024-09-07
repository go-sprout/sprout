package sprout

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrRecoverPanic(t *testing.T) {
	err := errors.New("test error")
	errMessage := "panic occurred"

	require.NotPanics(t, func() {
		defer ErrRecoverPanic(&err, errMessage)

		panic("test panic")
	})

	require.Error(t, err)
	assert.ErrorContains(t, err, errMessage)
}

func TestErrRecoverPanic_NoPanic(t *testing.T) {
	var err error
	errMessage := "panic occurred"

	defer ErrRecoverPanic(&err, errMessage)

	require.NoError(t, err)
}

func TestErrConvertFailed(t *testing.T) {
	baseErr := errors.New("test error")
	typ := "string"
	value := 1

	err := NewErrConvertFailed(typ, value, baseErr)

	require.Error(t, err)
	require.ErrorContains(t, err, "failed to convert: 1 to string")
	require.ErrorContains(t, err, "test error")
	require.ErrorIs(t, err, ErrConvertFailed)
	require.ErrorIs(t, err, baseErr)
}
