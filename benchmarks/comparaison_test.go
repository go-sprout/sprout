package benchmarks_test

import (
	"bytes"
	"log/slog"
	"sync"
	"testing"
	"text/template"

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
	"github.com/go-sprout/sprout/registry/time"
	"github.com/go-sprout/sprout/registry/uniqueid"
	"github.com/go-sprout/sprout/sprigin"
	"github.com/stretchr/testify/assert"
)

/**
 * BenchmarkSprig are the benchmarks for Sprig.
 * It is the same as SproutBench but with Sprig.
 */
func BenchmarkSprig(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tmpl, err := template.New("allFunctions").Funcs(sprig.FuncMap()).ParseGlob("*.sprig.tmpl")
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		for _, t := range tmpl.Templates() {
			err := tmpl.ExecuteTemplate(&buf, t.Name(), nil)
			if err != nil {
				panic(err)
			}
		}

		buf.Reset()
	}
}

/**
 * BenchmarkSprout are the benchmarks for Sprout.
 */
func BenchmarkSprout(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fnHandler := sprout.New(
			sprout.WithLogger(slog.New(&slog.TextHandler{})),
		)

		_ = fnHandler.AddRegistries(
			std.NewRegistry(),
			uniqueid.NewRegistry(),
			semver.NewRegistry(),
			backward.NewRegistry(),
			reflect.NewRegistry(),
			time.NewRegistry(),
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

		tmpl, err := template.New("allFunctions").Funcs(fnHandler.Build()).ParseGlob("*.sprout.tmpl")

		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		for _, t := range tmpl.Templates() {
			err := tmpl.ExecuteTemplate(&buf, t.Name(), fnHandler)
			if err != nil {
				panic(err)
			}
		}
		buf.Reset()
	}
}

func TestCompareSprigAndSprout(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(2)

	var bufSprig, bufSprout bytes.Buffer

	go func() {
		defer wg.Done()

		tmplSprig, err := template.New("compare").Funcs(sprig.FuncMap()).ParseFiles("compare.tmpl")
		assert.NoError(t, err)

		err = tmplSprig.ExecuteTemplate(&bufSprig, "compare.tmpl", nil)
		assert.NoError(t, err)
	}()

	go func() {
		defer wg.Done()

		tmplSprout, err := template.New("compare").Funcs(sprigin.FuncMap()).ParseGlob("compare.tmpl")
		assert.NoError(t, err)

		err = tmplSprout.ExecuteTemplate(&bufSprout, "compare.tmpl", nil)
		assert.NoError(t, err)
	}()

	wg.Wait()
	// sprig is expected and sprout is actual
	assert.Equal(t, bufSprig.String(), bufSprout.String())
}
