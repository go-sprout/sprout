package filesystem

import "github.com/go-sprout/sprout"

type FileSystemRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of filesystem registry.
func NewRegistry() *FileSystemRegistry {
	return &FileSystemRegistry{}
}

// Uid returns the unique identifier of the registry.
func (fsr *FileSystemRegistry) Uid() string {
	return "filesystem"
}

// LinkHandler links the handler to the registry at runtime.
func (fsr *FileSystemRegistry) LinkHandler(fh sprout.Handler) {
	fsr.handler = &fh
}
