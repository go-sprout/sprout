package env

import (
	"os"

	"github.com/go-sprout/sprout"
)

// RegisterFunctions registers all functions of the registry.
func (er *EnvironmentRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
	sprout.AddFunction(funcsMap, "env", er.Env)
	sprout.AddFunction(funcsMap, "expandEnv", er.ExpandEnv)
}

// Env retrieves the value of an environment variable.
//
// Parameters:
//
//	key string - the name of the environment variable.
//
// Returns:
//
//	string - the value of the environment variable.
//
// Example:
//
//	{{ "PATH" | env }} // Output: "/usr/bin:/bin:/usr/sbin:/sbin"
func (er *EnvironmentRegistry) Env(key string) string {
	return os.Getenv(key)
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
//
// Example:
//
//	{{ "Path is $PATH" | expandEnv }} // Output: "Path is /usr/bin:/bin:/usr/sbin:/sbin"
func (er *EnvironmentRegistry) ExpandEnv(str string) string {
	return os.ExpandEnv(str)
}
