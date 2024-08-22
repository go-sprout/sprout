package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"sort"
)

func main() {
	// Simulate parsed functions from YAML
	functions := []string{"toBool", "toLower", "toUpper", "olla", "toto"} // Example functions

	// Sort functions alphabetically
	sort.Strings(functions)

	// File path
	filePath := "functions.go"

	// Parse existing file or create a new AST if the file doesn't exist
	fset := token.NewFileSet()
	var file *ast.File
	var err error

	if _, err = os.Stat(filePath); err == nil {
		// If the file exists, parse it
		file, err = parser.ParseFile(fset, filePath, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}
	} else {
		// Create a new file structure
		file = &ast.File{
			Name:  ast.NewIdent("main"),
			Decls: []ast.Decl{},
		}
	}

	// Create a map to track existing functions
	existingFuncs := make(map[string]*ast.FuncDecl)
	var updatedDecls []ast.Decl

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			existingFuncs[fn.Name.Name] = fn
		} else {
			updatedDecls = append(updatedDecls, decl)
		}
	}

	// Add or update functions in the AST
	for _, funcName := range functions {
		if fn, exists := existingFuncs[funcName]; exists {
			// Keep existing function as is
			updatedDecls = append(updatedDecls, fn)
		} else {
			// Create a new placeholder function
			funcDecl := createPlaceholderFunction(funcName)
			updatedDecls = append(updatedDecls, funcDecl)
		}
	}

	// Sort the functions alphabetically in the AST
	sort.Slice(updatedDecls, func(i, j int) bool {
		fn1, ok1 := updatedDecls[i].(*ast.FuncDecl)
		fn2, ok2 := updatedDecls[j].(*ast.FuncDecl)
		if ok1 && ok2 {
			return fn1.Name.Name < fn2.Name.Name
		}
		return false
	})

	// Assign the sorted decls with empty lines back to the file
	file.Decls = updatedDecls

	// Write the updated AST back to the file
	output, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	err = format.Node(output, fset, file)
	if err != nil {
		panic(err)
	}
}

// createPlaceholderFunction creates an AST for a placeholder function.
func createPlaceholderFunction(name string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("v")},
						Type:  ast.NewIdent("interface{}"),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("interface{}"),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				// Placeholder comment
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun:  ast.NewIdent("panic"),
						Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"Function not implemented"`}},
					},
				},
			},
		},
	}
}
