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

// RegisterFunctions registers all functions of the registry.
func (fsr *FileSystemRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
	sprout.AddFunction(funcsMap, "pathBase", fsr.PathBase)
	sprout.AddFunction(funcsMap, "pathDir", fsr.PathDir)
	sprout.AddFunction(funcsMap, "pathExt", fsr.PathExt)
	sprout.AddFunction(funcsMap, "pathClean", fsr.PathClean)
	sprout.AddFunction(funcsMap, "pathIsAbs", fsr.PathIsAbs)
	sprout.AddFunction(funcsMap, "osBase", fsr.OsBase)
	sprout.AddFunction(funcsMap, "osDir", fsr.OsDir)
	sprout.AddFunction(funcsMap, "osExt", fsr.OsExt)
	sprout.AddFunction(funcsMap, "osClean", fsr.OsClean)
	sprout.AddFunction(funcsMap, "osIsAbs", fsr.OsIsAbs)
}
