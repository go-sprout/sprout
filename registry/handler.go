package registry

import "log/slog"

// Handler is the interface that wraps the basic methods of a handler to manage
// all registries and functions.
// The Handler brick is the main brick of sprout. It is used to configure and
// manage a cross-registry configuration and function management like a global
// logging system, error handling, and more.
// ! This interface is not meant to be implemented by the user but by the
// ! library itself. An user could implement it but it is not recommended.
type Handler interface {
	AddRegistry(registry Registry) error
	AddRegistries(registries ...Registry) error
	Registry() FunctionMap
	Logger() *slog.Logger
}
