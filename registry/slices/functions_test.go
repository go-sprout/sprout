package slices_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/slices"
)

func TestList(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ list }}`, Expected: "[]"},
		{Input: `{{ .V | list "ab" true 4 5 }}`, Expected: "[ab true 4 5 <nil>]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestAppend(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ append .V "a" }}`, Expected: "[a]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ append .V "a" }}`, Expected: "[a]", Data: map[string]any{"V": []string(nil)}},
		{Input: `{{ append .V "a" }}`, Expected: "[]", Data: map[string]any{"V": nil}},
		{Input: `{{ append .V "a" }}`, Expected: "[x a]", Data: map[string]any{"V": []string{"x"}}},
		{Input: `{{ append .V "a" }}`, Expected: "[x a]", Data: map[string]any{"V": [1]string{"x"}}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestPrepend(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ prepend .V "a" }}`, Expected: "[a]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ prepend .V "a" }}`, Expected: "[a]", Data: map[string]any{"V": []string(nil)}},
		{Input: `{{ prepend .V "a" }}`, Expected: "[]", Data: map[string]any{"V": nil}},
		{Input: `{{ prepend .V "a" }}`, Expected: "[a x]", Data: map[string]any{"V": []string{"x"}}},
		{Input: `{{ prepend .V "a" }}`, Expected: "[a x]", Data: map[string]any{"V": [1]string{"x"}}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestConcat(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ concat .V (list 1 2 3) }}`, Expected: "[a 1 2 3]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ list 4 5 | concat (list 1 2 3) }}`, Expected: "[1 2 3 4 5]"},
		{Input: `{{ concat .V (list 1 2 3) }}`, Expected: "[1 2 3]", Data: map[string]any{"V": nil}},
		{Input: `{{ concat .V (list 1 2 3) }}`, Expected: "[1 2 3]", Data: map[string]any{"V": "a"}},
		{Input: `{{ concat .V (list 1 2 3) }}`, Expected: "[x 1 2 3]", Data: map[string]any{"V": []string{"x"}}},
		{Input: `{{ concat .V (list 1 2 3) }}`, Expected: "[[x] 1 2 3]", Data: map[string]any{"V": [][]string{{"x"}}}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestChunk(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ chunk 2 .V }}`, Expected: "[[a b] [c d] [e]]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ chunk 2 .V }}`, Expected: "[[a b] [c d]]", Data: map[string]any{"V": []string{"a", "b", "c", "d"}}},
		{Input: `{{ chunk 2 .V }}`, Expected: "[[a b]]", Data: map[string]any{"V": []string{"a", "b"}}},
		{Input: `{{ chunk 2 .V }}`, Expected: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ chunk 2 .V }}`, Expected: "[]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestUniq(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ uniq .V }}`, Expected: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c", "a", "b", "c"}}},
		{Input: `{{ uniq .V }}`, Expected: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c"}}},
		{Input: `{{ uniq .V }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a", "a", "a"}}},
		{Input: `{{ uniq .V }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ uniq .V }}`, Expected: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ uniq .V }}`, Expected: "[]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestCompact(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ compact .V }}`, Expected: "[a b c]", Data: map[string]any{"V": []string{"a", "", "b", "", "c"}}},
		{Input: `{{ compact .V }}`, Expected: "[a a]", Data: map[string]any{"V": []string{"a", "", "a"}}},
		{Input: `{{ compact .V }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ compact .V }}`, Expected: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ compact .V }}`, Expected: "[]", Data: map[string]any{"V": nil}},
		{Input: `{{ list 1 0 "" "hello" | compact }}`, Expected: "[1 hello]"},
		{Input: `{{ list "" "" | compact }}`, Expected: "[]"},
		{Input: `{{ list | compact }}`, Expected: "[]"},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestSlice(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ slice .V }}`, Expected: "[a b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ slice .V 1 }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ slice .V 1 3 }}`, Expected: "[b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ slice .V 0 1 }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ slice .V 0 1  }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ slice .V 0 1 }}`, Expected: "<no value>", Data: map[string]any{"V": []string{}}},
		{Input: `{{ slice .V 0 1 }}`, Expected: "[]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestHas(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ .V | has "a" }}`, Expected: "true", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ .V | has "a" }}`, Expected: "false", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{Input: `{{ .V | has 1 }}`, Expected: "true", Data: map[string]any{"V": []any{"b", 1, nil, struct{}{}}}},
		{Input: `{{ .V | has .Nil }}`, Expected: "true", Data: map[string]any{"Nil": nil, "V": []any{"b", 1, nil, struct{}{}}}},
		{Input: `{{ .V | has "nope" }}`, Expected: "false", Data: map[string]any{"V": []any{"b", 1, nil, struct{}{}}}},
		{Input: `{{ .V | has 1 }}`, Expected: "true", Data: map[string]any{"V": []int{1}}},
		{Input: `{{ .V | has "a" }}`, Expected: "false", Data: map[string]any{"V": nil}},
		{Input: `{{ .V | has "a" }}`, Expected: "false", Data: map[string]any{"V": 1}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestWithout(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ without .V "a" }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ without .V "a" }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{Input: `{{ without .V "a" }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e", "a"}}},
		{Input: `{{ without .V "a" }}`, Expected: "[]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ without .V "a" }}`, Expected: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ without .V "a" }}`, Expected: "[]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestRest(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ rest .V }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ rest .V }}`, Expected: "[c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{Input: `{{ rest .V }}`, Expected: "[c d e a]", Data: map[string]any{"V": []string{"b", "c", "d", "e", "a"}}},
		{Input: `{{ rest .V }}`, Expected: "[]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ rest .V }}`, Expected: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ rest .V }}`, Expected: "[]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestInitial(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ initial .V }}`, Expected: "[a b c d]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ initial .V }}`, Expected: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d"}}},
		{Input: `{{ initial .V }}`, Expected: "[]", Data: map[string]any{"V": []string{"a"}}},
		{Input: `{{ initial .V }}`, Expected: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ initial .V }}`, Expected: "[]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestFirst(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ first .V }}`, Expected: "a", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ first .V }}`, Expected: "<no value>", Data: map[string]any{"V": []string{}}},
		{Input: `{{ first .V }}`, Expected: "<no value>", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestLast(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ last .V }}`, Expected: "e", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ last .V }}`, Expected: "<no value>", Data: map[string]any{"V": []string{}}},
		{Input: `{{ last .V }}`, Expected: "<no value>", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestReverse(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ reverse .V }}`, Expected: "[e d c b a]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ reverse .V }}`, Expected: "[a b c d e]", Data: map[string]any{"V": []string{"e", "d", "c", "b", "a"}}},
		{Input: `{{ reverse .V }}`, Expected: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ reverse .V }}`, Expected: "[]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestSortAlpha(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ sortAlpha .V }}`, Expected: "[a b c d e]", Data: map[string]any{"V": []string{"e", "d", "c", "b", "a"}}},
		{Input: `{{ sortAlpha .V }}`, Expected: "[1 2 3 4 5 a]", Data: map[string]any{"V": []any{5, 4, 3, 2, 1, "a"}}},
		{Input: `{{ sortAlpha .V }}`, Expected: "[]", Data: map[string]any{"V": []string{}}},
		{Input: `{{ sortAlpha .V }}`, Expected: "[<nil>]", Data: map[string]any{"V": nil}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestSplitList(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ .V | splitList "," }}`, Expected: "[a b c d e]", Data: map[string]any{"V": "a,b,c,d,e"}},
		{Input: `{{ .V | splitList "," }}`, Expected: "[a b c d e ]", Data: map[string]any{"V": "a,b,c,d,e,"}},
		{Input: `{{ .V | splitList "," }}`, Expected: "[ a b c d e]", Data: map[string]any{"V": ",a,b,c,d,e"}},
		{Input: `{{ .V | splitList "," }}`, Expected: "[ a b c d e ]", Data: map[string]any{"V": ",a,b,c,d,e,"}},
		{Input: `{{ .V | splitList "," }}`, Expected: "[]", Data: map[string]any{"V": ""}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestStrSlice(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{ strSlice .V }}`, Expected: "[a b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{Input: `{{ strSlice .V }}`, Expected: "[5 4 3 2 1]", Data: map[string]any{"V": []int{5, 4, 3, 2, 1}}},
		{Input: `{{ strSlice .V }}`, Expected: "[5 a true 1]", Data: map[string]any{"V": []any{5, "a", true, nil, 1}}},
		{Input: `{{ strSlice .V }}`, Expected: "[]", Data: map[string]any{"V": ""}},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestUntil(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{range $i, $e := until 5}}({{$i}}{{$e}}){{end}}`, Expected: "(00)(11)(22)(33)(44)"},
		{Input: `{{range $i, $e := until -5}}({{$i}}{{$e}}){{end}}`, Expected: "(00)(1-1)(2-2)(3-3)(4-4)"},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestUntilStep(t *testing.T) {
	var tc = []pesticide.SafeTestCase{
		{Input: `{{range $i, $e := untilStep 0 5 1}}({{$i}}{{$e}}){{end}}`, Expected: "(00)(11)(22)(33)(44)"},
		{Input: `{{range $i, $e := untilStep 3 6 1}}({{$i}}{{$e}}){{end}}`, Expected: "(03)(14)(25)"},
		{Input: `{{range $i, $e := untilStep 0 -10 -2}}({{$i}}{{$e}}){{end}}`, Expected: "(00)(1-2)(2-4)(3-6)(4-8)"},
		{Input: `{{range $i, $e := untilStep 3 0 1}}({{$i}}{{$e}}){{end}}`, Expected: ""},
		{Input: `{{range $i, $e := untilStep 3 99 0}}({{$i}}{{$e}}){{end}}`, Expected: ""},
		{Input: `{{range $i, $e := untilStep 3 99 -1}}({{$i}}{{$e}}){{end}}`, Expected: ""},
		{Input: `{{range $i, $e := untilStep 3 0 0}}({{$i}}{{$e}}){{end}}`, Expected: ""},
	}

	pesticide.RunSafeTestCases(t, slices.NewRegistry(), tc)
}

func TestMustAppend(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustAppend .V "a" }}`, Expected: "[a]", Data: map[string]any{"V": []string{}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustAppend .V "a" }}`, Expected: "[a]", Data: map[string]any{"V": []string(nil)}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustAppend .V "a" }}`, Expected: "[x a]", Data: map[string]any{"V": []string{"x"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustAppend .V "a" }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot append to nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustAppend .V "a" }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot append on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustPrepend(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustPrepend .V "a" }}`, Expected: "[a]", Data: map[string]any{"V": []string{}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustPrepend .V "a" }}`, Expected: "[a]", Data: map[string]any{"V": []string(nil)}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustPrepend .V "a" }}`, Expected: "[a x]", Data: map[string]any{"V": []string{"x"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustPrepend .V "a" }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot prepend to nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustPrepend .V "a" }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot prepend on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustChunk(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustChunk 2 .V }}`, Expected: "[[a b] [c d] [e]]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustChunk 2 .V }}`, Expected: "[[a b] [c d]]", Data: map[string]any{"V": []string{"a", "b", "c", "d"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustChunk 2 .V }}`, Expected: "[[a b]]", Data: map[string]any{"V": []string{"a", "b"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustChunk 2 .V }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot chunk nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustChunk 2 .V }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot chunk type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustUniq(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustUniq .V }}`, Expected: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c", "a", "b", "c"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustUniq .V }}`, Expected: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustUniq .V }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a", "a", "a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustUniq .V }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustUniq .V }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot uniq nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustUniq .V }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot find uniq on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustCompact(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustCompact .V }}`, Expected: "[a b c]", Data: map[string]any{"V": []string{"a", "", "b", "", "c"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustCompact .V }}`, Expected: "[a a]", Data: map[string]any{"V": []string{"a", "", "a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustCompact .V }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustCompact .V }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot compact nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustCompact .V }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot compact on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustSlice(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustSlice .V }}`, Expected: "[a b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustSlice .V 1 }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustSlice .V 1 3 }}`, Expected: "[b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustSlice .V 0 1 }}`, Expected: "[a]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustSlice .V 0 1 }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot slice nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustSlice .V 0 1 }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "list should be type of slice or array but int"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustSlice .V -1 1 }}`, Expected: "", Data: map[string]any{"V": []string{"a"}}}, ExpectedErr: "start index out of bounds"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustSlice .V 0 52 }}`, Expected: "", Data: map[string]any{"V": []string{"a"}}}, ExpectedErr: "end index out of bounds"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustHas(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ .V | mustHas "a" }}`, Expected: "true", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ .V | mustHas "a" }}`, Expected: "false", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ .V | mustHas 1 }}`, Expected: "true", Data: map[string]any{"V": []any{"b", 1, nil, struct{}{}}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ .V | mustHas .Nil }}`, Expected: "true", Data: map[string]any{"Nil": nil, "V": []any{"b", 1, nil, struct{}{}}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ .V | mustHas "nope" }}`, Expected: "false", Data: map[string]any{"V": []any{"b", 1, nil, struct{}{}}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ .V | mustHas 1 }}`, Expected: "true", Data: map[string]any{"V": []int{1}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ .V | mustHas "a" }}`, Expected: "false", Data: map[string]any{"V": nil}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ .V | mustHas "a" }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot find has on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustWithout(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustWithout .V "a" }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustWithout .V "a" }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustWithout .V "a" }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e", "a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustWithout .V "a" }}`, Expected: "[]", Data: map[string]any{"V": []string{"a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustWithout .V "a" }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot without nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustWithout .V "a" }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot find without on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustRest(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustRest .V }}`, Expected: "[b c d e]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustRest .V }}`, Expected: "[c d e]", Data: map[string]any{"V": []string{"b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustRest .V }}`, Expected: "[c d e a]", Data: map[string]any{"V": []string{"b", "c", "d", "e", "a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustRest .V }}`, Expected: "[]", Data: map[string]any{"V": []string{"a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustRest .V }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot rest nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustRest .V }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot find rest on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustInitial(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustInitial .V }}`, Expected: "[a b c d]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustInitial .V }}`, Expected: "[a b c]", Data: map[string]any{"V": []string{"a", "b", "c", "d"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustInitial .V }}`, Expected: "[]", Data: map[string]any{"V": []string{"a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustInitial .V }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot initial nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustInitial .V }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot find initial on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustFirst(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustFirst .V }}`, Expected: "a", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustFirst .V }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot first nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustFirst .V }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot find first on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustLast(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustLast .V }}`, Expected: "e", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustLast .V }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot last nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustLast .V }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot find last on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}

func TestMustReverse(t *testing.T) {
	var tc = []pesticide.TestCase{
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustReverse .V }}`, Expected: "[e d c b a]", Data: map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustReverse .V }}`, Expected: "[a b c d e]", Data: map[string]any{"V": []string{"e", "d", "c", "b", "a"}}}, ExpectedErr: ""},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustReverse .V }}`, Expected: "", Data: map[string]any{"V": nil}}, ExpectedErr: "cannot reverse nil"},
		{TestCase: pesticide.SafeTestCase{Input: `{{ mustReverse .V }}`, Expected: "", Data: map[string]any{"V": 1}}, ExpectedErr: "cannot find reverse on type int"},
	}

	pesticide.RunTestCases(t, slices.NewRegistry(), tc)
}
