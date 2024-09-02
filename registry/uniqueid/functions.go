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
// Example:
//
//	{{ uuidv4 }} // Output: "3f0c463e-53f5-4f05-a2ec-3c083aa8f937"
func (ur *UniqueIDRegistry) Uuidv4() string {
	return uuid.New().String()
}
