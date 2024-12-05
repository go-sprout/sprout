package filesystem

import (
	"path"
	"path/filepath"
)

// PathBase returns the last element of the path.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	string - the base element of the path.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: pathBase].
//
// [Sprout Documentation: pathBase]: https://docs.atom.codes/sprout/registries/filesystem#pathbase
func (fsr *FileSystemRegistry) PathBase(value string) string {
	return path.Base(value)
}

// PathDir returns all but the last element of the path, effectively the path's
// directory.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	string - the directory part of the path.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: pathDir].
//
// [Sprout Documentation: pathDir]: https://docs.atom.codes/sprout/registries/filesystem#pathdir
func (fsr *FileSystemRegistry) PathDir(value string) string {
	return path.Dir(value)
}

// PathExt returns the file extension of the path.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	string - the extension of the file in the path.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: pathExt].
//
// [Sprout Documentation: pathExt]: https://docs.atom.codes/sprout/registries/filesystem#pathext
func (fsr *FileSystemRegistry) PathExt(value string) string {
	return path.Ext(value)
}

// PathClean cleans up the path, simplifying any redundancies like double slashes.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	string - the cleaned path.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: pathClean].
//
// [Sprout Documentation: pathClean]: https://docs.atom.codes/sprout/registries/filesystem#pathclean
func (fsr *FileSystemRegistry) PathClean(value string) string {
	return path.Clean(value)
}

// PathIsAbs checks if the path is absolute.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	bool - true if the path is absolute, otherwise false.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: pathIsAbs].
//
// [Sprout Documentation: pathIsAbs]: https://docs.atom.codes/sprout/registries/filesystem#pathisabs
func (fsr *FileSystemRegistry) PathIsAbs(value string) bool {
	return path.IsAbs(value)
}

// OsBase returns the last element of the path, using the OS-specific path
// separator.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	string - the base element of the path.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: osBase].
//
// [Sprout Documentation: osBase]: https://docs.atom.codes/sprout/registries/filesystem#osbase
func (fsr *FileSystemRegistry) OsBase(value string) string {
	return filepath.Base(value)
}

// OsDir returns all but the last element of the path, using the OS-specific
// path separator.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	string - the directory part of the path.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: osDir].
//
// [Sprout Documentation: osDir]: https://docs.atom.codes/sprout/registries/filesystem#osdir
func (fsr *FileSystemRegistry) OsDir(value string) string {
	return filepath.Dir(value)
}

// OsExt returns the file extension of the path, using the OS-specific path
// separator.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	string - the extension of the file in the path.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: osExt].
//
// [Sprout Documentation: osExt]: https://docs.atom.codes/sprout/registries/filesystem#osext
func (fsr *FileSystemRegistry) OsExt(value string) string {
	return filepath.Ext(value)
}

// OsClean cleans up the path, using the OS-specific path separator and
// simplifying redundancies.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	string - the cleaned path.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: osClean].
//
// [Sprout Documentation: osClean]: https://docs.atom.codes/sprout/registries/filesystem#osclean
func (fsr *FileSystemRegistry) OsClean(value string) string {
	return filepath.Clean(value)
}

// OsIsAbs checks if the path is absolute, using the OS-specific path separator.
//
// Parameters:
//
//	value string - the path string.
//
// Returns:
//
//	bool - true if the path is absolute, otherwise false.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: osIsAbs].
//
// [Sprout Documentation: osIsAbs]: https://docs.atom.codes/sprout/registries/filesystem#osisabs
func (fsr *FileSystemRegistry) OsIsAbs(value string) bool {
	return filepath.IsAbs(value)
}
