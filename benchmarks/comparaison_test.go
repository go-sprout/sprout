package main

import (
	"bytes"
	"log/slog"
	"testing"
	"text/template"

	"github.com/42atomys/sprout"
	sprig "github.com/Masterminds/sprig/v3"
)

type ErrHandling int

const (
	ErrHandlingReturnDefaultValue ErrHandling = iota + 1
	ErrHandlingPanic
	ErrHandlingErrorChannel
)

type parser struct {
	ErrHandling ErrHandling
	errChan     chan error
	Logger      *slog.Logger
}

type Option func(*parser)

func NewParser(opts ...Option) *parser {
	parser := &parser{
		ErrHandling: ErrHandlingReturnDefaultValue,
		errChan:     make(chan error),
		Logger:      slog.New(&slog.TextHandler{}),
	}

	for _, opt := range opts {
		opt(parser)
	}

	return parser
}

func WithErrHandling(eh ErrHandling) Option {
	return func(p *parser) {
		p.ErrHandling = eh
	}
}

func WithLogger(l *slog.Logger) Option {
	return func(p *parser) {
		p.Logger = l
	}
}

func WithErrorChannel(ec chan error) Option {
	return func(p *parser) {
		p.errChan = ec
	}
}

func WithParser(newParser *parser) Option {
	return func(p *parser) {
		*p = *newParser
	}
}

func (p *parser) Hello() string {
	return "Hello, World!"
}

// BACKWARD COMPATIBILITY
func FuncsMap(opts ...Option) template.FuncMap {
	parser := NewParser(opts...)

	funcmap := sprout.FuncMap()

	funcmap["Hello"] = parser.Hello

	return funcmap
}

func main() {
	SproutBench()
}

// SpringBench est le bloc de code que vous souhaitez benchmark.
func SproutBench() {
	errChan := make(chan error)
	defer close(errChan)

	// Créer une instance de Context
	parser := NewParser(
		WithErrHandling(ErrHandlingPanic),
		WithLogger(slog.New(&slog.TextHandler{})),
		WithErrorChannel(errChan),
	)

	go func() {
		for err := range parser.errChan {
			parser.Logger.Error(err.Error())
		}
	}()

	// Parse le template en ajoutant les méthodes de Context via Funcs.
	tmpl, err := template.New("allFunctions").Funcs(FuncsMap(WithParser(parser))).ParseGlob("*.sprout.tmpl")

	if err != nil {
		panic(err)
	}

	// Exécuter le template en passant l'instance de Context comme donnée.
	// Notez que dans ce cas, le contexte lui-même n'est pas utilisé directement dans le template.
	var buf bytes.Buffer
	for _, t := range tmpl.Templates() {
		err := tmpl.ExecuteTemplate(&buf, t.Name(), parser)
		if err != nil {
			panic(err)
		}
	}
	buf.Reset()
}

func SprigBench() {
	// Parse le template en ajoutant les méthodes de Context via Funcs.
	tmpl, err := template.New("allFunctions").Funcs(sprig.FuncMap()).ParseGlob("*.sprig.tmpl")
	if err != nil {
		panic(err)
	}

	// Exécuter le template en passant l'instance de Context comme donnée.
	// Notez que dans ce cas, le contexte lui-même n'est pas utilisé directement dans le template.
	var buf bytes.Buffer
	for _, t := range tmpl.Templates() {
		err := tmpl.ExecuteTemplate(&buf, t.Name(), nil)
		if err != nil {
			panic(err)
		}
	}

	buf.Reset()
}

func BenchmarkSprig(b *testing.B) {
	// b.ResetTimer() est utilisé pour s'assurer que toute initialisation préalable
	// ne soit pas prise en compte dans le temps du benchmark.
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SprigBench()
	}
}

func BenchmarkSprout(b *testing.B) {
	// b.ResetTimer() est utilisé pour s'assurer que toute initialisation préalable
	// ne soit pas prise en compte dans le temps du benchmark.
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SproutBench()
	}
}
