package main

import (
	"os"
	"path/filepath"
)

// FileSystem is an interface that abstracts file system operations.
// This allows us to inject mock implementations for testing.
type FileSystem interface {
	Glob(pattern string) ([]string, error)
	ReadFile(filename string) ([]byte, error)
}

// OSFileSystem is a concrete implementation of FileSystem using the os and filepath packages.
type OSFileSystem struct{}

func (fs *OSFileSystem) Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

func (fs *OSFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
