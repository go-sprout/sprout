package sprout

import (
	"fmt"
	"testing"
)

func TestToBool(t *testing.T) {
	var tests = testCases{
		{"TestBool", `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, "bool-true", map[string]any{"V": true}},
		{"TestInt", `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, "bool-true", map[string]any{"V": 1}},
		{"TestInt32", `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, "bool-true", map[string]any{"V": int32(1)}},
		{"TestFloat64", `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, "bool-true", map[string]any{"V": float64(1.42)}},
		{"TestString", `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, "bool-true", map[string]any{"V": "true"}},
		{"TestStringFalse", `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, "bool-false", map[string]any{"V": "false"}},
		{"TestStringInvalid", `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, "bool-false", map[string]any{"V": "invalid"}},
	}

	runTestCases(t, tests)
}

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

func TestToUint(t *testing.T) {
	var tests = testCases{
		{"TestInt", `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, "uint-1", map[string]any{"V": 1}},
		{"TestInt32", `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, "uint-1", map[string]any{"V": int32(1)}},
		{"TestFloat64", `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, "uint-1", map[string]any{"V": float64(1.42)}},
		{"TestBool", `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, "uint-1", map[string]any{"V": true}},
		{"TestString", `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, "uint-1", map[string]any{"V": "1"}},
	}

	runTestCases(t, tests)
}

func TestToUint64(t *testing.T) {
	var tests = testCases{
		{"TestInt", `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, "uint64-1", map[string]any{"V": 1}},
		{"TestInt32", `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, "uint64-1", map[string]any{"V": int32(1)}},
		{"TestFloat64", `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, "uint64-1", map[string]any{"V": float64(1.42)}},
		{"TestBool", `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, "uint64-1", map[string]any{"V": true}},
		{"TestString", `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, "uint64-1", map[string]any{"V": "1"}},
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

func TestToDuration(t *testing.T) {
	var tests = testCases{
		{"TestInt", `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, "time.Duration-1ns", map[string]any{"V": 1}},
		{"TestInt32", `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, "time.Duration-1µs", map[string]any{"V": int32(1000)}},
		{"TestFloat64", `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, "time.Duration-1.00042ms", map[string]any{"V": float64(1000 * 1000.42)}},
		{"TestString", `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, "time.Duration-1m0s", map[string]any{"V": "1m"}},
		{"TestInvalid", `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, "time.Duration-0s", map[string]any{"V": "aaaa"}},
		{"TestCallingOnIt", `{{ (toDuration "1h30m").Seconds }}`, "5400", nil},
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
