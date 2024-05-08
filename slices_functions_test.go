package sprout

import "testing"

func TestList(t *testing.T) {
	var tests = testCases{
		{"", `{{ list }}`, "[]", nil},
		{"", `{{ .V | list "ab" true 4 5 }}`, "[ab true 4 5 <nil>]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestAppend(t *testing.T) {
	var tests = testCases{
		{"", `{{ append .V "a" }}`, "[a]", map[string]any{"V": []string{}}},
		{"", `{{ append .V "a" }}`, "[a]", map[string]any{"V": []string(nil)}},
		{"", `{{ append .V "a" }}`, "[]", map[string]any{"V": nil}},
		{"", `{{ append .V "a" }}`, "[x a]", map[string]any{"V": []string{"x"}}},
	}

	runTestCases(t, tests)
}

func TestPrepend(t *testing.T) {
	var tests = testCases{
		{"", `{{ prepend .V "a" }}`, "[a]", map[string]any{"V": []string{}}},
		{"", `{{ prepend .V "a" }}`, "[a]", map[string]any{"V": []string(nil)}},
		{"", `{{ prepend .V "a" }}`, "[]", map[string]any{"V": nil}},
		{"", `{{ prepend .V "a" }}`, "[a x]", map[string]any{"V": []string{"x"}}},
	}

	runTestCases(t, tests)
}

func TestConcat(t *testing.T) {
	var tests = testCases{
		{"", `{{ concat .V (list 1 2 3) }}`, "[a 1 2 3]", map[string]any{"V": []string{"a"}}},
		{"", `{{ list 4 5 | concat (list 1 2 3) }}`, "[1 2 3 4 5]", nil},
		{"", `{{ concat .V (list 1 2 3) }}`, "[1 2 3]", map[string]any{"V": nil}},
		{"", `{{ concat .V (list 1 2 3) }}`, "[1 2 3]", map[string]any{"V": "a"}},
		{"", `{{ concat .V (list 1 2 3) }}`, "[x 1 2 3]", map[string]any{"V": []string{"x"}}},
		{"", `{{ concat .V (list 1 2 3) }}`, "[[x] 1 2 3]", map[string]any{"V": [][]string{{"x"}}}},
	}

	runTestCases(t, tests)
}

func TestChunk(t *testing.T) {
	var tests = testCases{
		{"", `{{ chunk 2 .V }}`, "[[a b] [c d] [e]]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ chunk 2 .V }}`, "[[a b] [c d]]", map[string]any{"V": []string{"a", "b", "c", "d"}}},
		{"", `{{ chunk 2 .V }}`, "[[a b]]", map[string]any{"V": []string{"a", "b"}}},
		{"", `{{ chunk 2 .V }}`, "[]", map[string]any{"V": []string{}}},
		{"", `{{ chunk 2 .V }}`, "[]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestUniq(t *testing.T) {
	var tests = testCases{
		{"", `{{ uniq .V }}`, "[a b c]", map[string]any{"V": []string{"a", "b", "c", "a", "b", "c"}}},
		{"", `{{ uniq .V }}`, "[a b c]", map[string]any{"V": []string{"a", "b", "c"}}},
		{"", `{{ uniq .V }}`, "[a]", map[string]any{"V": []string{"a", "a", "a"}}},
		{"", `{{ uniq .V }}`, "[a]", map[string]any{"V": []string{"a"}}},
		{"", `{{ uniq .V }}`, "[]", map[string]any{"V": []string{}}},
		{"", `{{ uniq .V }}`, "[]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestCompact(t *testing.T) {
	var tests = testCases{
		{"", `{{ compact .V }}`, "[a b c]", map[string]any{"V": []string{"a", "", "b", "", "c"}}},
		{"", `{{ compact .V }}`, "[a a]", map[string]any{"V": []string{"a", "", "a"}}},
		{"", `{{ compact .V }}`, "[a]", map[string]any{"V": []string{"a"}}},
		{"", `{{ compact .V }}`, "[]", map[string]any{"V": []string{}}},
		{"", `{{ compact .V }}`, "[]", map[string]any{"V": nil}},
		{"", `{{ list 1 0 "" "hello" | compact }}`, "[1 hello]", nil},
		{"", `{{ list "" "" | compact }}`, "[]", nil},
		{"", `{{ list | compact }}`, "[]", nil},
	}

	runTestCases(t, tests)
}

func TestSlice(t *testing.T) {
	var tests = testCases{
		{"", `{{ slice .V }}`, "[a b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ slice .V 1 }}`, "[b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ slice .V 1 3 }}`, "[b c]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ slice .V 0 1 }}`, "[a]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ slice .V 0 1  }}`, "[a]", map[string]any{"V": []string{"a"}}},
		{"", `{{ slice .V 0 1 }}`, "<no value>", map[string]any{"V": []string{}}},
		{"", `{{ slice .V 0 1 }}`, "[]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestWithout(t *testing.T) {
	var tests = testCases{
		{"", `{{ without .V "a" }}`, "[b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ without .V "a" }}`, "[b c d e]", map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{"", `{{ without .V "a" }}`, "[b c d e]", map[string]any{"V": []string{"b", "c", "d", "e", "a"}}},
		{"", `{{ without .V "a" }}`, "[]", map[string]any{"V": []string{"a"}}},
		{"", `{{ without .V "a" }}`, "[]", map[string]any{"V": []string{}}},
		{"", `{{ without .V "a" }}`, "[]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestRest(t *testing.T) {
	var tests = testCases{
		{"", `{{ rest .V }}`, "[b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ rest .V }}`, "[c d e]", map[string]any{"V": []string{"b", "c", "d", "e"}}},
		{"", `{{ rest .V }}`, "[c d e a]", map[string]any{"V": []string{"b", "c", "d", "e", "a"}}},
		{"", `{{ rest .V }}`, "[]", map[string]any{"V": []string{"a"}}},
		{"", `{{ rest .V }}`, "[]", map[string]any{"V": []string{}}},
		{"", `{{ rest .V }}`, "[]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestInitial(t *testing.T) {
	var tests = testCases{
		{"", `{{ initial .V }}`, "[a b c d]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ initial .V }}`, "[a b c]", map[string]any{"V": []string{"a", "b", "c", "d"}}},
		{"", `{{ initial .V }}`, "[]", map[string]any{"V": []string{"a"}}},
		{"", `{{ initial .V }}`, "[]", map[string]any{"V": []string{}}},
		{"", `{{ initial .V }}`, "[]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestFirst(t *testing.T) {
	var tests = testCases{
		{"", `{{ first .V }}`, "a", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ first .V }}`, "<no value>", map[string]any{"V": []string{}}},
		{"", `{{ first .V }}`, "<no value>", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestLast(t *testing.T) {
	var tests = testCases{
		{"", `{{ last .V }}`, "e", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ last .V }}`, "<no value>", map[string]any{"V": []string{}}},
		{"", `{{ last .V }}`, "<no value>", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestReverse(t *testing.T) {
	var tests = testCases{
		{"", `{{ reverse .V }}`, "[e d c b a]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ reverse .V }}`, "[a b c d e]", map[string]any{"V": []string{"e", "d", "c", "b", "a"}}},
		{"", `{{ reverse .V }}`, "[]", map[string]any{"V": []string{}}},
		{"", `{{ reverse .V }}`, "[]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestSortAlpha(t *testing.T) {
	var tests = testCases{
		{"", `{{ sortAlpha .V }}`, "[a b c d e]", map[string]any{"V": []string{"e", "d", "c", "b", "a"}}},
		{"", `{{ sortAlpha .V }}`, "[1 2 3 4 5 a]", map[string]any{"V": []any{5, 4, 3, 2, 1, "a"}}},
		{"", `{{ sortAlpha .V }}`, "[]", map[string]any{"V": []string{}}},
		{"", `{{ sortAlpha .V }}`, "[<nil>]", map[string]any{"V": nil}},
	}

	runTestCases(t, tests)
}

func TestSplitList(t *testing.T) {
	var tests = testCases{
		{"", `{{ .V | splitList "," }}`, "[a b c d e]", map[string]any{"V": "a,b,c,d,e"}},
		{"", `{{ .V | splitList "," }}`, "[a b c d e ]", map[string]any{"V": "a,b,c,d,e,"}},
		{"", `{{ .V | splitList "," }}`, "[ a b c d e]", map[string]any{"V": ",a,b,c,d,e"}},
		{"", `{{ .V | splitList "," }}`, "[ a b c d e ]", map[string]any{"V": ",a,b,c,d,e,"}},
		{"", `{{ .V | splitList "," }}`, "[]", map[string]any{"V": ""}},
	}

	runTestCases(t, tests)
}

func TestStrSlice(t *testing.T) {
	var tests = testCases{
		{"", `{{ strSlice .V }}`, "[a b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}},
		{"", `{{ strSlice .V }}`, "[5 4 3 2 1]", map[string]any{"V": []int{5, 4, 3, 2, 1}}},
		{"", `{{ strSlice .V }}`, "[5 a true 1]", map[string]any{"V": []any{5, "a", true, nil, 1}}},
		{"", `{{ strSlice .V }}`, "[]", map[string]any{"V": ""}},
	}

	runTestCases(t, tests)
}

func TestMustAppend(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustAppend .V "a" }}`, "[a]", map[string]any{"V": []string{}}}, ""},
		{testCase{"", `{{ mustAppend .V "a" }}`, "[a]", map[string]any{"V": []string(nil)}}, ""},
		{testCase{"", `{{ mustAppend .V "a" }}`, "[x a]", map[string]any{"V": []string{"x"}}}, ""},
		{testCase{"", `{{ mustAppend .V "a" }}`, "", map[string]any{"V": nil}}, "cannot append to nil"},
		{testCase{"", `{{ mustAppend .V "a" }}`, "", map[string]any{"V": 1}}, "cannot append on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustPrepend(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustPrepend .V "a" }}`, "[a]", map[string]any{"V": []string{}}}, ""},
		{testCase{"", `{{ mustPrepend .V "a" }}`, "[a]", map[string]any{"V": []string(nil)}}, ""},
		{testCase{"", `{{ mustPrepend .V "a" }}`, "[a x]", map[string]any{"V": []string{"x"}}}, ""},
		{testCase{"", `{{ mustPrepend .V "a" }}`, "", map[string]any{"V": nil}}, "cannot prepend to nil"},
		{testCase{"", `{{ mustPrepend .V "a" }}`, "", map[string]any{"V": 1}}, "cannot prepend on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustChunk(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustChunk 2 .V }}`, "[[a b] [c d] [e]]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustChunk 2 .V }}`, "[[a b] [c d]]", map[string]any{"V": []string{"a", "b", "c", "d"}}}, ""},
		{testCase{"", `{{ mustChunk 2 .V }}`, "[[a b]]", map[string]any{"V": []string{"a", "b"}}}, ""},
		{testCase{"", `{{ mustChunk 2 .V }}`, "", map[string]any{"V": nil}}, "cannot chunk nil"},
		{testCase{"", `{{ mustChunk 2 .V }}`, "", map[string]any{"V": 1}}, "cannot chunk type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustUniq(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustUniq .V }}`, "[a b c]", map[string]any{"V": []string{"a", "b", "c", "a", "b", "c"}}}, ""},
		{testCase{"", `{{ mustUniq .V }}`, "[a b c]", map[string]any{"V": []string{"a", "b", "c"}}}, ""},
		{testCase{"", `{{ mustUniq .V }}`, "[a]", map[string]any{"V": []string{"a", "a", "a"}}}, ""},
		{testCase{"", `{{ mustUniq .V }}`, "[a]", map[string]any{"V": []string{"a"}}}, ""},
		{testCase{"", `{{ mustUniq .V }}`, "", map[string]any{"V": nil}}, "cannot uniq nil"},
		{testCase{"", `{{ mustUniq .V }}`, "", map[string]any{"V": 1}}, "cannot find uniq on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustCompact(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustCompact .V }}`, "[a b c]", map[string]any{"V": []string{"a", "", "b", "", "c"}}}, ""},
		{testCase{"", `{{ mustCompact .V }}`, "[a a]", map[string]any{"V": []string{"a", "", "a"}}}, ""},
		{testCase{"", `{{ mustCompact .V }}`, "[a]", map[string]any{"V": []string{"a"}}}, ""},
		{testCase{"", `{{ mustCompact .V }}`, "", map[string]any{"V": nil}}, "cannot compact nil"},
		{testCase{"", `{{ mustCompact .V }}`, "", map[string]any{"V": 1}}, "cannot compact on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustSlice(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustSlice .V }}`, "[a b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustSlice .V 1 }}`, "[b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustSlice .V 1 3 }}`, "[b c]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustSlice .V 0 1 }}`, "[a]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustSlice .V 0 1 }}`, "", map[string]any{"V": nil}}, "cannot slice nil"},
		{testCase{"", `{{ mustSlice .V 0 1 }}`, "", map[string]any{"V": 1}}, "list should be type of slice or array but int"},
	}

	runMustTestCases(t, tests)
}

func TestMustWithout(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustWithout .V "a" }}`, "[b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustWithout .V "a" }}`, "[b c d e]", map[string]any{"V": []string{"b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustWithout .V "a" }}`, "[b c d e]", map[string]any{"V": []string{"b", "c", "d", "e", "a"}}}, ""},
		{testCase{"", `{{ mustWithout .V "a" }}`, "[]", map[string]any{"V": []string{"a"}}}, ""},
		{testCase{"", `{{ mustWithout .V "a" }}`, "", map[string]any{"V": nil}}, "cannot without nil"},
		{testCase{"", `{{ mustWithout .V "a" }}`, "", map[string]any{"V": 1}}, "cannot find without on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustRest(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustRest .V }}`, "[b c d e]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustRest .V }}`, "[c d e]", map[string]any{"V": []string{"b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustRest .V }}`, "[c d e a]", map[string]any{"V": []string{"b", "c", "d", "e", "a"}}}, ""},
		{testCase{"", `{{ mustRest .V }}`, "[]", map[string]any{"V": []string{"a"}}}, ""},
		{testCase{"", `{{ mustRest .V }}`, "", map[string]any{"V": nil}}, "cannot rest nil"},
		{testCase{"", `{{ mustRest .V }}`, "", map[string]any{"V": 1}}, "cannot find rest on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustInitial(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustInitial .V }}`, "[a b c d]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustInitial .V }}`, "[a b c]", map[string]any{"V": []string{"a", "b", "c", "d"}}}, ""},
		{testCase{"", `{{ mustInitial .V }}`, "[]", map[string]any{"V": []string{"a"}}}, ""},
		{testCase{"", `{{ mustInitial .V }}`, "", map[string]any{"V": nil}}, "cannot initial nil"},
		{testCase{"", `{{ mustInitial .V }}`, "", map[string]any{"V": 1}}, "cannot find initial on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustFirst(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustFirst .V }}`, "a", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustFirst .V }}`, "", map[string]any{"V": nil}}, "cannot first nil"},
		{testCase{"", `{{ mustFirst .V }}`, "", map[string]any{"V": 1}}, "cannot find first on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustLast(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustLast .V }}`, "e", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustLast .V }}`, "", map[string]any{"V": nil}}, "cannot last nil"},
		{testCase{"", `{{ mustLast .V }}`, "", map[string]any{"V": 1}}, "cannot find last on type int"},
	}

	runMustTestCases(t, tests)
}

func TestMustReverse(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"", `{{ mustReverse .V }}`, "[e d c b a]", map[string]any{"V": []string{"a", "b", "c", "d", "e"}}}, ""},
		{testCase{"", `{{ mustReverse .V }}`, "[a b c d e]", map[string]any{"V": []string{"e", "d", "c", "b", "a"}}}, ""},
		{testCase{"", `{{ mustReverse .V }}`, "", map[string]any{"V": nil}}, "cannot reverse nil"},
		{testCase{"", `{{ mustReverse .V }}`, "", map[string]any{"V": 1}}, "cannot find reverse on type int"},
	}

	runMustTestCases(t, tests)
}
