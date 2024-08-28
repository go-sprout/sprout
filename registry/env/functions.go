package env

import (
	"os"
)

// Env retrieves the value of an environment variable.
//
// Parameters:
//
//	key string - the name of the environment variable.
//
// Returns:
//
//	string - the value of the environment variable.
//	error - a placeholder for future error handling.
//
// Example:
//
//	{{ "PATH" | env }} // Output: "/usr/bin:/bin:/usr/sbin:/sbin"
func (er *EnvironmentRegistry) Env(key string) (string, error) {
	return os.Getenv(key), nil
}

// ExpandEnv replaces ${var} or $var in the string based on the values of the
// current environment variables.
//
// Parameters:
//
//	str string - the string with environment variables to expand.
//
// Returns:
//
//	string - the expanded string.
//	error - a placeholder for future error handling.
//
// Example:
//
//	{{ "Path is $PATH" | expandEnv }} // Output: "Path is /usr/bin:/bin:/usr/sbin:/sbin"
func (er *EnvironmentRegistry) ExpandEnv(str string) (string, error) {
	return os.ExpandEnv(str), nil
}
