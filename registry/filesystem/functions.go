package filesystem

import (
	"path"
	"path/filepath"
)

// PathBase returns the last element of the path.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	string - the base element of the path.
//
// Example:
//
//	{{ "/path/to/file.txt" | pathBase }} // Output: "file.txt"
func (fsr *FileSystemRegistry) PathBase(str string) string {
	return path.Base(str)
}

// PathDir returns all but the last element of the path, effectively the path's
// directory.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	string - the directory part of the path.
//
// Example:
//
//	{{ "/path/to/file.txt" | pathDir }} // Output: "/path/to"
func (fsr *FileSystemRegistry) PathDir(str string) string {
	return path.Dir(str)
}

// PathExt returns the file extension of the path.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	string - the extension of the file in the path.
//
// Example:
//
//	{{ "/path/to/file.txt" | pathExt }} // Output: ".txt"
func (fsr *FileSystemRegistry) PathExt(str string) string {
	return path.Ext(str)
}

// PathClean cleans up the path, simplifying any redundancies like double slashes.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	string - the cleaned path.
//
// Example:
//
//	{{ "/path//to/file.txt" | pathClean }} // Output: "/path/to/file.txt"
func (fsr *FileSystemRegistry) PathClean(str string) string {
	return path.Clean(str)
}

// PathIsAbs checks if the path is absolute.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	bool - true if the path is absolute, otherwise false.
//
// Example:
//
//	{{ "/path/to/file.txt" | pathIsAbs }} // Output: true
func (fsr *FileSystemRegistry) PathIsAbs(str string) bool {
	return path.IsAbs(str)
}

// OsBase returns the last element of the path, using the OS-specific path
// separator.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	string - the base element of the path.
//
// Example:
//
//	{{ "C:\\path\\to\\file.txt" | osBase }} // Output: "file.txt"
func (fsr *FileSystemRegistry) OsBase(str string) string {
	return filepath.Base(str)
}

// OsDir returns all but the last element of the path, using the OS-specific
// path separator.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	string - the directory part of the path.
//
// Example:
//
//	{{ "C:\\path\\to\\file.txt" | osDir }} // Output: "C:\\path\\to"
func (fsr *FileSystemRegistry) OsDir(str string) string {
	return filepath.Dir(str)
}

// OsExt returns the file extension of the path, using the OS-specific path
// separator.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	string - the extension of the file in the path.
//
// Example:
//
//	{{ "C:\\path\\to\\file.txt" | osExt }} // Output: ".txt"
func (fsr *FileSystemRegistry) OsExt(str string) string {
	return filepath.Ext(str)
}

// OsClean cleans up the path, using the OS-specific path separator and
// simplifying redundancies.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	string - the cleaned path.
//
// Example:
//
//	{{ "C:\\path\\\\to\\file.txt" | osClean }} // Output: "C:\\path\\to\\file.txt"
func (fsr *FileSystemRegistry) OsClean(str string) string {
	return filepath.Clean(str)
}

// OsIsAbs checks if the path is absolute, using the OS-specific path separator.
//
// Parameters:
//
//	str string - the path string.
//
// Returns:
//
//	bool - true if the path is absolute, otherwise false.
//
// Example:
//
//	{{ "C:\\path\\to\\file.txt" | osIsAbs }} // Output: true
func (fsr *FileSystemRegistry) OsIsAbs(str string) bool {
	return filepath.IsAbs(str)
}
