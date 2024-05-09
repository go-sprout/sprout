package benchmarks_test

import (
	"bytes"
	"log/slog"
	"sync"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/go-sprout/sprout"
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
		errChan := make(chan error)
		defer close(errChan)

		fnHandler := sprout.NewFunctionHandler(
			sprout.WithErrHandling(sprout.ErrHandlingPanic),
			sprout.WithLogger(slog.New(&slog.TextHandler{})),
			sprout.WithErrorChannel(errChan),
		)

		go func() {
			for err := range errChan {
				fnHandler.Logger.Error(err.Error())
			}
		}()

		tmpl, err := template.New("allFunctions").Funcs(sprout.FuncMap(sprout.WithFunctionHandler(fnHandler))).ParseGlob("*.sprout.tmpl")

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

		tmplSprout, err := template.New("compare").Funcs(sprout.FuncMap()).ParseGlob("compare.tmpl")
		assert.NoError(t, err)

		err = tmplSprout.ExecuteTemplate(&bufSprout, "compare.tmpl", nil)
		assert.NoError(t, err)
	}()

	wg.Wait()
	// sprig is expected and sprout is actual
	assert.Equal(t, bufSprig.String(), bufSprout.String())
}
