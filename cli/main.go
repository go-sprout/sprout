package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"log/slog"
	"os"
	"path"
	gostrings "strings"
	"text/template"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/conversion"
	"github.com/go-sprout/sprout/registry/numeric"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/strings"
	"github.com/go-sprout/sprout/registry/time"
	"gopkg.in/yaml.v3"
)

const srcPath = "registry/conversion/manual_functions.go"
const configPath = "registry/conversion/registry.yaml"

var logger *slog.Logger
var tmpl *template.Template

type RegistryConfig struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string `yaml:"name"`
		Author      string `yaml:"author"`
		Description string `yaml:"description"`
	} `yaml:"metadata"`
	Spec struct {
		Notices []struct {
			Functions []string `yaml:"functions"`
			Kind      string   `yaml:"kind"`
			Message   string   `yaml:"message"`
		} `yaml:"notices"`
		Functions []struct {
			Name        string   `yaml:"name"`
			Description string   `yaml:"description"`
			Aliases     []string `yaml:"aliases"`
			Parameters  []struct {
				Name        string `yaml:"name"`
				Type        string `yaml:"type"`
				Description string `yaml:"description"`
			} `yaml:"parameters"`
			Returns struct {
				Type        string `yaml:"type"`
				Description string `yaml:"description"`
			} `yaml:"returns"`
			Examples []struct {
				Template string `yaml:"template"`
				Result   string `yaml:"result"`
			} `yaml:"examples"`
		} `yaml:"functions"`
	} `yaml:"spec"`
}

type Package struct {
	Name      string
	Imports   []string
	Functions []Function

	SrcFile string
}

type Function struct {
	Name          string
	Body          string
	Documentation []string
	Receiver      NameType
	Parameters    NameTypeSlice
	Results       NameTypeSlice
}

type NameTypeSlice []NameType

type NameType struct {
	Name string
	Type string
}

func (nt NameType) String() string {
	if nt.Name == "" {
		return nt.Type
	}

	return nt.Name + " " + nt.Type
}

func (nts NameTypeSlice) String() string {
	var res []string
	for _, nt := range nts {
		res = append(res, nt.String())
	}

	return gostrings.Join(res, ", ")
}

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	handler := sprout.New(sprout.WithLogger(logger))

	handler.AddRegistries(
		conversion.NewRegistry(),
		numeric.NewRegistry(),
		strings.NewRegistry(),
		slices.NewRegistry(),
		time.NewRegistry(),
	)

	tmpl = template.Must(template.New("").Funcs(handler.Build()).ParseGlob("*.tmpl"))
}

func main() {

	if len(os.Args) < 2 {
		os.Args = append(os.Args, "generate")
	}

	switch os.Args[1] {
	case "scaffold":
		scaffold()
		return
	case "generate":
		generateFunctions()
		return
	default:
		generateFunctions()
	}
}

func scaffold() {
	// Read the configuration file
	configFile, err := os.Open(configPath)
	if err != nil {
		logger.Error("Failed to open config file", "err", err)
	}

	config := &RegistryConfig{}
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(config)
	if err != nil {
		logger.Error("Failed to decode config file", "err", err)
	}

	generateRegistryFile(config)
}

func generateRegistryFile(config *RegistryConfig) {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "registry.go.tmpl", config)
	if err != nil {
		logger.Error("Failed to execute template", "err", err)
		return
	}

	// Write to a file in the folder of the source file named generated_functions.go
	file, err := os.Create(fmt.Sprintf("%s/%s.go", path.Dir(configPath), config.Metadata.Name))
	if err != nil {
		logger.Error("Failed to create file", "err", err)
		return
	}

	_, err = file.Write(buf.Bytes())
	if err != nil {
		logger.Error("Failed to write to file", "err", err)
		return
	}

	logger.Info("Generated code written to file", "file", file.Name())
}

func generateRegistryFunctions() {

}

func generateFunctions() {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, srcPath, nil, parser.ParseComments)
	if err != nil {
		logger.Error("Failed to parse file", "err", err.Error())
	}

	if node.Name == nil {
		logger.Error("No package name found")
		return
	}

	// Extract package name
	d := &Package{
		Name:    node.Name.Name,
		SrcFile: srcPath,
	}

	// Extract imports
	for _, imprt := range node.Imports {
		d.Imports = append(d.Imports, gostrings.ReplaceAll(imprt.Path.Value, "\"", ""))
	}

	for _, decl := range node.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Extract function name
		functionData := &Function{
			Name: fn.Name.Name,
			Receiver: NameType{
				Name: fn.Recv.List[0].Names[0].Name,
				Type: extractType(fn.Recv.List[0].Type),
			},
		}

		// Extract comments
		for _, comment := range fn.Doc.List {
			functionData.Documentation = append(functionData.Documentation, comment.Text)
		}

		// Extract parameters
		for _, param := range fn.Type.Params.List {
			for _, name := range param.Names {
				functionData.Parameters = append(functionData.Parameters, NameType{
					Name: name.Name,
					Type: extractType(param.Type),
				})
			}
		}

		// Extract return type
		for i, result := range fn.Type.Results.List {
			nt := NameType{
				Type: extractType(result.Type),
			}

			if len(result.Names) > i {
				nt.Name = result.Names[i].Name
			}

			functionData.Results = append(functionData.Results, nt)

		}

		// Extract function body
		functionData.Body += blockStmtToString(fset, fn.Body)

		d.Functions = append(d.Functions, *functionData)
	}

	generateCode(d)
}

func generateCode(data *Package) {

	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "functions.go.tmpl", data)
	if err != nil {
		logger.Error("Failed to execute template", "err", err)
		return
	}

	// Write to a file in the folder of the source file named generated_functions.go
	file, err := os.Create(fmt.Sprintf("%s/generated_functions.go", path.Dir(data.SrcFile)))
	if err != nil {
		logger.Error("Failed to create file", "err", err)
		return
	}

	_, err = file.Write(buf.Bytes())
	if err != nil {
		logger.Error("Failed to write to file", "err", err)
		return
	}

	logger.Info("Generated code written to file", "file", file.Name())

}

func extractType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + extractType(t.X)
	case *ast.ArrayType:
		return "[]" + extractType(t.Elt)
	default:
		return "unknown"
	}
}

func blockStmtToString(fset *token.FileSet, block *ast.BlockStmt) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, block)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
