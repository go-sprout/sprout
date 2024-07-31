package pesticide

import (
	"testing"

	"github.com/go-sprout/sprout"
	"github.com/stretchr/testify/assert"
)

type RegexpTestCase struct {
	Name     string
	Template string
	Regexp   string
	Length   int
}

func RunRegexpTestCases(t *testing.T, registry sprout.Registry, tcs []RegexpTestCase) {
	t.Helper()
	handler := sprout.New()
	_ = handler.AddRegistry(registry)

	for _, test := range tcs {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()

			result, err := runTemplate(t, handler, test.Template, nil)
			assert.NoError(t, err)

			assert.Regexp(t, test.Regexp, result)
			if test.Length != -1 {
				assert.Len(t, result, test.Length)
			}
		})
	}
}
