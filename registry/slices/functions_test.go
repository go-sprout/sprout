package slices_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/slices"
)

func TestList(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ list }}`, ExpectedOutput: "[]"},
		{Input: `{{ .V | list "ab" true 4 5 }}`, ExpectedOutput: "[ab true 4 5 <nil>]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestAppend(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ .V | append "a" }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ .V | append "a" }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string(nil)}},
		{Input: `{{ .V | append "a" }}`, ExpectedOutput: "[x a]", Data: map[string]any{"V": []string{"x"}}},
		{Input: `{{ .V | append "a" }}`, ExpectedOutput: "[x a]", Data: map[string]any{"V": [1]string{"x"}}},
		{Input: `{{ .V | append "a" }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot append to nil"},
		{Input: `{{ .V | append "a" }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot append on type int"},
		{Input: `{{ append }}`, ExpectedErr: " expected 2 arguments, got 0"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestPrepend(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ .V | prepend "a" }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ .V | prepend "a" }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string(nil)}},
		{Input: `{{ .V | prepend "a" }}`, ExpectedOutput: "[a x]", Data: map[string]any{"V": []string{"x"}}},
		{Input: `{{ .V | prepend "a" }}`, ExpectedOutput: "[a x]", Data: map[string]any{"V": [1]string{"x"}}},
		{Input: `{{ .V | prepend "a" }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot prepend to nil"},
		{Input: `{{ .V | prepend "a" }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot prepend on type int"},
		{Input: `{{ prepend }}`, ExpectedErr: " expected 2 arguments, got 0"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestConcat(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ concat .V (list 1 2 3) }}`, ExpectedOutput: "[a 1 2 3]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ list 4 5 | concat (list 1 2 3) }}`, ExpectedOutput: "[1 2 3 4 5]"},
		{Input: `{{ concat .V (list 1 2 3) }}`, ExpectedOutput: "[1 2 3]", Data: map[string]any{"V": nil}},
		{Input: `{{ concat .V (list 1 2 3) }}`, ExpectedOutput: "[1 2 3]", Data: map[string]any{"V": "a"}},
		{Input: `{{ concat .V (list 1 2 3) }}`, ExpectedOutput: "[x 1 2 3]", Data: map[string]any{"V": []string{"x"}}},
		{Input: `{{ concat .V (list 1 2 3) }}`, ExpectedOutput: "[[x] 1 2 3]", Data: map[string]any{"V": [][]string{{"x"}}}},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestChunk(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ chunk 2 .V }}`, ExpectedOutput: "[[a b] [c d] [e]]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ chunk 2 .V }}`, ExpectedOutput: "[[a b] [c d]]", Data: map[string]any{"V": []string{"a", "b", "c", "d"}}},
		{Input: `{{ chunk 2 .V }}`, ExpectedOutput: "[[a b]]", Data: map[string]any{"V": []string{"a", "b"}}},
		{Input: `{{ chunk 2 .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ chunk 2 .V }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot chunk nil"},
		{Input: `{{ chunk 2 .V }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot chunk type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestUniq(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ uniq .V }}`, ExpectedOutput: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c", "a", "b", "c"}}},
		{Input: `{{ uniq .V }}`, ExpectedOutput: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c"}}},
		{Input: `{{ uniq .V }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a", "a", "a"}}},
		{Input: `{{ uniq .V }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ uniq .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ uniq .V }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot uniq nil"},
		{Input: `{{ uniq .V }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot find uniq on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestCompact(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ compact .V }}`, ExpectedOutput: "[a b c]", Data: map[string]any{"V": []string{"a", "", "b", "", "c"}}},
		{Input: `{{ compact .V }}`, ExpectedOutput: "[a a]", Data: map[string]any{"V": []string{"a", "", "a"}}},
		{Input: `{{ compact .V }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ compact .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ list 1 0 "" "hello" | compact }}`, ExpectedOutput: "[1 hello]"},
		{Input: `{{ list "" "" | compact }}`, ExpectedOutput: "[]"},
		{Input: `{{ list | compact }}`, ExpectedOutput: "[]"},
		{Input: `{{ compact .V }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot compact nil"},
		{Input: `{{ compact .V }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot compact on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestSlice(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ .V | slice }}`, ExpectedOutput: "[a b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ .V | slice 1 }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ .V | slice 1 3 }}`, ExpectedOutput: "[b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ .V | slice 0 1 }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ .V | slice 0 1  }}`, ExpectedOutput: "[a]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ .V | slice 0 1 }}`, ExpectedOutput: "<no value>", Data: map[string]any{"V": []string{}}},
		{Input: `{{ .V | slice 0 1 }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot slice nil"},
		{Input: `{{ .V | slice 0 1 }}`, Data: map[string]any{"V": 1}, ExpectedErr: "last argument must be a slice but got int"},
		{Input: `{{ .V | slice -1 1 }}`, Data: map[string]any{"V": []string{"a"}}, ExpectedErr: "start index out of bounds"},
		{Input: `{{ .V | slice 0 52 }}`, Data: map[string]any{"V": []string{"a"}}, ExpectedErr: "end index out of bounds"},
		{Input: `{{ slice }}`, ExpectedErr: "expected 2 arguments, got 0"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestHas(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ .V | has "a" }}`, ExpectedOutput: "true", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ .V | has "a" }}`, ExpectedOutput: "false", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{Input: `{{ .V | has 1 }}`, ExpectedOutput: "true", Data: map[string]any{"V": []any{"b", 1, nil, struct{}{}}}},
		{Input: `{{ .V | has .Nil }}`, ExpectedOutput: "true", Data: map[string]any{"Nil": nil, "V": []any{"b", 1, nil, struct{}{}}}},
		{Input: `{{ .V | has "nope" }}`, ExpectedOutput: "false", Data: map[string]any{"V": []any{"b", 1, nil, struct{}{}}}},
		{Input: `{{ .V | has 1 }}`, ExpectedOutput: "true", Data: map[string]any{"V": []int{1}}},
		{Input: `{{ .V | has "a" }}`, ExpectedOutput: "false", Data: map[string]any{"V": nil}},
		{Input: `{{ .V | has "a" }}`, ExpectedErr: "cannot find has on type int", Data: map[string]any{"V": 1}},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestWithout(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ .V | without "a" }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ .V | without "a" }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{Input: `{{ .V | without "a" }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e", "a"}}},
		{Input: `{{ .V | without "a" }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ .V | without "a" }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ .V | without "a" }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot without nil"},
		{Input: `{{ .V | without "a" }}`, Data: map[string]any{"V": 1}, ExpectedErr: "last argument must be a slice but got int"},
		{Input: `{{ without  }}`, ExpectedErr: "expected 2 arguments, got 0"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestRest(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ rest .V }}`, ExpectedOutput: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ rest .V }}`, ExpectedOutput: "[c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{Input: `{{ rest .V }}`, ExpectedOutput: "[c d e a]", Data: map[string]any{"V": []string{"b", "c", "d", "e", "a"}}},
		{Input: `{{ rest .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ rest .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ rest .V }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot rest nil"},
		{Input: `{{ rest .V }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot find rest on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestInitial(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ initial .V }}`, ExpectedOutput: "[a b c d]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ initial .V }}`, ExpectedOutput: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d"}}},
		{Input: `{{ initial .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ initial .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ initial .V }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot initial nil"},
		{Input: `{{ initial .V }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot find initial on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestFirst(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ first .V }}`, ExpectedOutput: "a", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ first .V }}`, ExpectedOutput: "<no value>", Data: map[string]any{"V": []string{}}},
		{Input: `{{ first .V }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot first nil"},
		{Input: `{{ first .V }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot find first on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestLast(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ last .V }}`, ExpectedOutput: "e", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ last .V }}`, ExpectedOutput: "<no value>", Data: map[string]any{"V": []string{}}},
		{Input: `{{ last .V }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot last nil"},
		{Input: `{{ last .V }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot find last on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestReverse(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ reverse .V }}`, ExpectedOutput: "[e d c b a]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ reverse .V }}`, ExpectedOutput: "[a b c d e]", Data: map[string]any{"V": []string{"e", "d", "c", "b", "a"}}},
		{Input: `{{ reverse .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ reverse .V }}`, Data: map[string]any{"V": nil}, ExpectedErr: "cannot reverse nil"},
		{Input: `{{ reverse .V }}`, Data: map[string]any{"V": 1}, ExpectedErr: "cannot find reverse on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestSortAlpha(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ sortAlpha .V }}`, ExpectedOutput: "[a b c d e]", Data: map[string]any{"V": []string{"e", "d", "c", "b", "a"}}},
		{Input: `{{ sortAlpha .V }}`, ExpectedOutput: "[1 2 3 4 5 a]", Data: map[string]any{"V": []any{5, 4, 3, 2, 1, "a"}}},
		{Input: `{{ sortAlpha .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ sortAlpha .V }}`, ExpectedOutput: "[<nil>]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestSplitList(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ .V | splitList "," }}`, ExpectedOutput: "[a b c d e]", Data: map[string]any{"V": "a,b,c,d,e"}},
		{Input: `{{ .V | splitList "," }}`, ExpectedOutput: "[a b c d e ]", Data: map[string]any{"V": "a,b,c,d,e,"}},
		{Input: `{{ .V | splitList "," }}`, ExpectedOutput: "[ a b c d e]", Data: map[string]any{"V": ",a,b,c,d,e"}},
		{Input: `{{ .V | splitList "," }}`, ExpectedOutput: "[ a b c d e ]", Data: map[string]any{"V": ",a,b,c,d,e,"}},
		{Input: `{{ .V | splitList "," }}`, ExpectedOutput: "[]", Data: map[string]any{"V": ""}},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestStrSlice(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ strSlice .V }}`, ExpectedOutput: "[a b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ strSlice .V }}`, ExpectedOutput: "[5 4 3 2 1]", Data: map[string]any{"V": []int{5, 4, 3, 2, 1}}},
		{Input: `{{ strSlice .V }}`, ExpectedOutput: "[5 a true 1]", Data: map[string]any{"V": []any{5, "a", true, nil, 1}}},
		{Input: `{{ strSlice .V }}`, ExpectedOutput: "[]", Data: map[string]any{"V": ""}},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestUntil(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{range $i, $e := until 5}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: "(00)(11)(22)(33)(44)"},
		{Input: `{{range $i, $e := until -5}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: "(00)(1-1)(2-2)(3-3)(4-4)"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestUntilStep(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{range $i, $e := untilStep 0 5 1}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: "(00)(11)(22)(33)(44)"},
		{Input: `{{range $i, $e := untilStep 3 6 1}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: "(03)(14)(25)"},
		{Input: `{{range $i, $e := untilStep 0 -10 -2}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: "(00)(1-2)(2-4)(3-6)(4-8)"},
		{Input: `{{range $i, $e := untilStep 3 0 1}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: ""},
		{Input: `{{range $i, $e := untilStep 3 99 0}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: ""},
		{Input: `{{range $i, $e := untilStep 3 99 -1}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: ""},
		{Input: `{{range $i, $e := untilStep 3 0 0}}({{$i}}{{$e}}){{end}}`, ExpectedOutput: ""},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}
