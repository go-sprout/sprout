package sprout

import (
	"fmt"
	"testing"
)

func TestToInt(t *testing.T) {
	var tests = testCases{
		{"TestInt", `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, "int-1", map[string]any{"V": 1}},
		{"TestInt32", `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, "int-1", map[string]any{"V": int32(1)}},
		{"TestFloat64", `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, "int-1", map[string]any{"V": float64(1.42)}},
		{"TestBool", `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, "int-1", map[string]any{"V": true}},
		{"TestString", `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, "int-1", map[string]any{"V": "1"}},
	}

	runTestCases(t, tests)
}

func TestToInt64(t *testing.T) {
	var tests = testCases{
		{"TestInt", `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, "int64-1", map[string]any{"V": 1}},
		{"TestInt32", `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, "int64-1", map[string]any{"V": int32(1)}},
		{"TestFloat64", `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, "int64-1", map[string]any{"V": float64(1.42)}},
		{"TestBool", `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, "int64-1", map[string]any{"V": true}},
		{"TestString", `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, "int64-1", map[string]any{"V": "1"}},
	}

	runTestCases(t, tests)
}

func TestToFloat64(t *testing.T) {
	var tests = testCases{
		{"TestInt", `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, "float64-1", map[string]any{"V": 1}},
		{"TestInt32", `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, "float64-1", map[string]any{"V": int32(1)}},
		{"TestFloat64", `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, "float64-1.42", map[string]any{"V": float64(1.42)}},
		{"TestBool", `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, "float64-1", map[string]any{"V": true}},
		{"TestString", `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, "float64-1", map[string]any{"V": "1"}},
	}

	runTestCases(t, tests)
}

func TestToOctal(t *testing.T) {
	var tests = testCases{
		{"TestInt", `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, "int64-511", map[string]any{"V": 777}},
		{"TestInt32", `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, "int64-504", map[string]any{"V": int32(770)}},
		{"TestString", `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, "int64-1", map[string]any{"V": "1"}},
		{"TestInvalid", `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, "int64-0", map[string]any{"V": 1.1}},
	}

	runTestCases(t, tests)
}

type testStringer struct{}

func (s testStringer) String() string {
	return "stringer"
}

func TestToString(t *testing.T) {

	var tests = testCases{
		{"TestInt", `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, "string-1", map[string]any{"V": 1}},
		{"TestInt32", `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, "string-1", map[string]any{"V": int32(1)}},
		{"TestFloat64", `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, "string-1.42", map[string]any{"V": float64(1.42)}},
		{"TestBool", `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, "string-true", map[string]any{"V": true}},
		{"TestString", `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, "string-1", map[string]any{"V": "1"}},
		{"TestError", `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, "string-error", map[string]any{"V": fmt.Errorf("error")}},
		{"TestStringer", `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, "string-stringer", map[string]any{"V": testStringer{}}},
		{"TestSliceOfBytes", `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, "string-abc", map[string]any{"V": []byte("abc")}},
	}

	runTestCases(t, tests)
}

func TestToDate(t *testing.T) {
	var tests = testCases{
		{"TestDate", `{{$v := toDate "2006-01-02" .V }}{{typeOf $v}}-{{$v}}`, "time.Time-2024-05-09 00:00:00 +0000 UTC", map[string]any{"V": "2024-05-09"}},
	}

	runTestCases(t, tests)
}

func TestMustToDate(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestDate", `{{$v := mustToDate "2006-01-02" .V }}{{typeOf $v}}-{{$v}}`, "time.Time-2024-05-09 00:00:00 +0000 UTC", map[string]any{"V": "2024-05-09"}}, ""},
		{testCase{"TestInvalidValue", `{{$v := mustToDate "2006-01-02" .V }}{{typeOf $v}}-{{$v}}`, "", map[string]any{"V": ""}}, "cannot parse \"\" as \"2006\""},
		{testCase{"TestInvalidLayout", `{{$v := mustToDate "invalid" .V }}{{typeOf $v}}-{{$v}}`, "", map[string]any{"V": "2024-05-09"}}, "cannot parse \"2024-05-09\" as \"invalid\""},
	}

	runMustTestCases(t, tests)
}
