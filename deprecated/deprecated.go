// This package are used to store all helpers functions used to validate and log
// deprecated functions and signatures in the project codebase. This is really
// useful to ensure no breaking changes are introduced from the migration of
// Sprig to Sprout.
//
// This package will be removed when no deprecated functions are left in the
// project codebase (scheduled for v1.1).
package deprecated

import (
	"fmt"
	"log/slog"
)

// SignatureWarn logs a warning message about a deprecated function signature.
//
// Parameters:
//
//	l *slog.Logger - the logger to use.
//	functionName string - the name of the function.
//	oldSign string - the old signature of the function.
//	newSign string - the new signature of the function.
//
// Example:
//
//	deprecated.SignatureWarn(l, "get", "{{ get dict key }}", "{{ dict | get key }}")
func SignatureWarn(l *slog.Logger, functionName, oldSign, newSign string) {
	msg := fmt.Sprintf("The signature of `%s` has changed from `%s` to `%s`, please update your code before next upgrade. This change will simplify the usage of the function and respect go/template conventions and allow usage of pipe (`|`).", functionName, oldSign, newSign)

	l.With("function", functionName, "notice", "deprecated").Warn(fmt.Sprintf("Template function `%s` is deprecated: %s", functionName, msg))
}

// ErrArgsCount returns an error message about the number of arguments expected and received.
//
// Parameters:
//
//	expected int - the number of arguments expected.
//	got int - the number of arguments received.
//
// Returns:
//
//	error - the error message.
//
// Example:
//
//	deprecated.ErrArgsCount(2, 3) // Output: "expected 2 arguments, got 3"
func ErrArgsCount(expected, got int) error {
	return fmt.Errorf("expected %d arguments, got %d", expected, got)
}
