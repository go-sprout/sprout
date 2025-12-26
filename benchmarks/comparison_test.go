package benchmarks_test

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"sync"
	"testing"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/backward"
	"github.com/go-sprout/sprout/registry/checksum"
	"github.com/go-sprout/sprout/registry/conversion"
	"github.com/go-sprout/sprout/registry/crypto"
	"github.com/go-sprout/sprout/registry/encoding"
	"github.com/go-sprout/sprout/registry/env"
	"github.com/go-sprout/sprout/registry/filesystem"
	"github.com/go-sprout/sprout/registry/maps"
	"github.com/go-sprout/sprout/registry/numeric"
	"github.com/go-sprout/sprout/registry/random"
	"github.com/go-sprout/sprout/registry/reflect"
	"github.com/go-sprout/sprout/registry/regexp"
	"github.com/go-sprout/sprout/registry/semver"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/std"
	"github.com/go-sprout/sprout/registry/strings"
	rtime "github.com/go-sprout/sprout/registry/time"
	"github.com/go-sprout/sprout/registry/uniqueid"
	"github.com/go-sprout/sprout/sprigin"
	"github.com/stretchr/testify/assert"
)

var data = map[string]any{
	"string":      "example string value",
	"intString":   "1",
	"url":         "https://example.com",
	"arrayCommas": "a,b,c",
	"int":         123,
	"float":       123.456,
	"bool":        true,
	"array":       []any{"a", "b", "c"},
	"map":         map[string]any{"foo": "bar", "nested": map[string]any{"far": "bee"}},
	"object":      struct{ Name string }{"example object"},
	"func":        func() string { return "example function" },
	"error":       fmt.Errorf("example error"),
	"time":        time.Now(),
	"duration":    5 * time.Second,
	"channel":     make(chan any),
	"json":        `{"foo": "bar"}`,
	"yaml":        "foo: bar",
	"nil":         nil,
}

type noopHandler struct{}

func (h *noopHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return false // Disable logging
}

func (h *noopHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil // Do nothing
}

func (h *noopHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *noopHandler) WithGroup(name string) slog.Handler {
	return h
}

func BenchmarkComparison(b *testing.B) {
	runtime.GOMAXPROCS(1)

	b.Run("Sprig", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sprigBench()
		}
	})

	b.Run("Sprout", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sproutBench("allFunctionsAsSprig.sprout.tmpl")
		}
	})
}

func BenchmarkAllSprout(b *testing.B) {
	runtime.GOMAXPROCS(1)

	b.Run("SproutWithNewFeatures", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sproutBench("allFunctions.sprout.tmpl")
		}
	})
}

/**
 * BenchmarkSprig are the benchmarks for Sprig.
 * It is the same as SproutBench but with Sprig.
 */
func sprigBench() {
	tmpl, err := template.New("allFunctions").Funcs(sprig.FuncMap()).ParseGlob("allFunctions.sprig.tmpl")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	for _, t := range tmpl.Templates() {
		err := tmpl.ExecuteTemplate(&buf, t.Name(), data)
		if err != nil {
			panic(err)
		}
	}

	buf.Reset()
}

/**
 * BenchmarkSprout are the benchmarks for Sprout.
 */
func sproutBench(templatePath string) {
	fnHandler := sprout.New(
		sprout.WithLogger(slog.New(&noopHandler{})),
		sprout.WithSafeFuncs(true),
	)

	_ = fnHandler.AddRegistries(
		std.NewRegistry(),
		uniqueid.NewRegistry(),
		semver.NewRegistry(),
		backward.NewRegistry(),
		reflect.NewRegistry(),
		rtime.NewRegistry(),
		strings.NewRegistry(),
		random.NewRegistry(),
		checksum.NewRegistry(),
		conversion.NewRegistry(),
		numeric.NewRegistry(),
		encoding.NewRegistry(),
		regexp.NewRegistry(),
		slices.NewRegistry(),
		maps.NewRegistry(),
		crypto.NewRegistry(),
		filesystem.NewRegistry(),
		env.NewRegistry(),
	)

	tmpl, err := template.New("allFunctions").Funcs(fnHandler.Build()).ParseGlob(templatePath)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	for _, t := range tmpl.Templates() {
		err := tmpl.ExecuteTemplate(&buf, t.Name(), data)
		if err != nil {
			panic(err)
		}
	}
	buf.Reset()
}

func TestCompareSprigAndSprout(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(2)

	var bufSprig, bufSprout bytes.Buffer

	go func() {
		defer wg.Done()

		tmplSprig, err := template.New("compare").Funcs(sprig.FuncMap()).ParseFiles("compare.tmpl")
		assert.NoError(t, err)

		err = tmplSprig.ExecuteTemplate(&bufSprig, "compare.tmpl", data)
		assert.NoError(t, err)
	}()

	go func() {
		defer wg.Done()

		tmplSprout, err := template.New("compare").Funcs(sprigin.FuncMap()).ParseGlob("compare.tmpl")
		assert.NoError(t, err)

		err = tmplSprout.ExecuteTemplate(&bufSprout, "compare.tmpl", data)
		assert.NoError(t, err)
	}()

	wg.Wait()
	// sprig is expected (---) and sprout is actual (+++)
	assert.Equal(t, bufSprig.String(), bufSprout.String())
}
