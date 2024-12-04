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
// For an example of this function in a go template, refer to [Sprout Documentation: pathBase].
//
// [Sprout Documentation: pathBase]: https://docs.atom.codes/sprout/registries/filesystem#pathbase
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
// For an example of this function in a go template, refer to [Sprout Documentation: pathDir].
//
// [Sprout Documentation: pathDir]: https://docs.atom.codes/sprout/registries/filesystem#pathdir
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
// For an example of this function in a go template, refer to [Sprout Documentation: pathExt].
//
// [Sprout Documentation: pathExt]: https://docs.atom.codes/sprout/registries/filesystem#pathext
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
// For an example of this function in a go template, refer to [Sprout Documentation: pathClean].
//
// [Sprout Documentation: pathClean]: https://docs.atom.codes/sprout/registries/filesystem#pathclean
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
// For an example of this function in a go template, refer to [Sprout Documentation: pathIsAbs].
//
// [Sprout Documentation: pathIsAbs]: https://docs.atom.codes/sprout/registries/filesystem#pathisabs
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
// For an example of this function in a go template, refer to [Sprout Documentation: osBase].
//
// [Sprout Documentation: osBase]: https://docs.atom.codes/sprout/registries/filesystem#osbase
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
// For an example of this function in a go template, refer to [Sprout Documentation: osDir].
//
// [Sprout Documentation: osDir]: https://docs.atom.codes/sprout/registries/filesystem#osdir
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
// For an example of this function in a go template, refer to [Sprout Documentation: osExt].
//
// [Sprout Documentation: osExt]: https://docs.atom.codes/sprout/registries/filesystem#osext
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
// For an example of this function in a go template, refer to [Sprout Documentation: osClean].
//
// [Sprout Documentation: osClean]: https://docs.atom.codes/sprout/registries/filesystem#osclean
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
// For an example of this function in a go template, refer to [Sprout Documentation: osIsAbs].
//
// [Sprout Documentation: osIsAbs]: https://docs.atom.codes/sprout/registries/filesystem#osisabs
func (fsr *FileSystemRegistry) OsIsAbs(str string) bool {
	return filepath.IsAbs(str)
}
