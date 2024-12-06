package uniqueid

import (
	"github.com/google/uuid"
)

// Uuidv4 generates a new random UUID (Universally Unique Identifier) version 4.
// This function does not take parameters and returns a string representation
// of a UUID.
//
// Returns:
//
//	string - a new UUID string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uuidv4].
//
// [Sprout Documentation: uuidv4]: https://docs.atom.codes/sprout/registries/uniqueid#uuidv4
func (ur *UniqueIDRegistry) Uuidv4() string {
	return uuid.New().String()
}
