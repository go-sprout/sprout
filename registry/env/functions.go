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
//
// For an example of this function in a Go template, refer to [Sprout Documentation: env].
//
// [Sprout Documentation: env]: https://docs.atom.codes/sprout/registries/env#env
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
// For an example of this function in a Go template, refer to [Sprout Documentation: expandEnv].
//
// [Sprout Documentation: expandEnv]: https://docs.atom.codes/sprout/registries/env#expandenv
func (er *EnvironmentRegistry) ExpandEnv(str string) string {
	return os.ExpandEnv(str)
}
