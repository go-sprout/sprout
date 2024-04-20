package errors_test

import (
	"testing"

	"github.com/42atomys/sprout/errors"
)

// Benchmark the creation of new errorStruct instances.
func BenchmarkNewError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errors.New("test error")
	}
}

// Benchmark the string output of an error chain.
func BenchmarkErrorOutput(b *testing.B) {
	root := errors.New("root error")
	mid := errors.New("middle error", root)
	end := errors.New("end error", mid)

	b.ResetTimer() // Start timing after setup.
	for i := 0; i < b.N; i++ {
		_ = end.Error()
	}
}

// Benchmark the Cause method.
func BenchmarkCause(b *testing.B) {
	root := errors.New("root error")
	mid := errors.New("middle error", root)
	end := errors.New("end error", mid)

	b.ResetTimer() // Start timing after setup.
	for i := 0; i < b.N; i++ {
		_ = end.(errors.ErrorChain).Cause()
	}
}

// Benchmark Stack computation.
func BenchmarkStack(b *testing.B) {
	root := errors.New("root error")
	mid := errors.New("middle error", root)
	end := errors.New("end error", mid)

	b.ResetTimer() // Start timing after setup.
	for i := 0; i < b.N; i++ {
		_ = end.Stack()
	}
}

func BenchmarkNestingTest(b *testing.B) {
	b.Run("NewError", func(b *testing.B) {
		b.ResetTimer()

		var err errors.Error
		for i := 0; i < b.N; i++ {
			err = errors.New("test error", err)
		}

		b.StopTimer()

		if len(err.Stack()) != b.N+1 {
			b.Logf("stack: %+v", err.Stack())
			b.Fatalf("expected stack length to be %d, got %d", b.N+1, len(err.Stack()))
		}
	})
}
