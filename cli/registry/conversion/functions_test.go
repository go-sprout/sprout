package conversion_test

import (
	"fmt"
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/conversion"
)

func TestToBool(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestBool", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, Expected: "bool-true", Data: map[string]any{"V": true}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, Expected: "bool-true", Data: map[string]any{"V": 1}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, Expected: "bool-true", Data: map[string]any{"V": int32(1)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestFloat64", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, Expected: "bool-true", Data: map[string]any{"V": float64(1.42)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, Expected: "bool-true", Data: map[string]any{"V": "true"}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestStringFalse", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, Expected: "bool-false", Data: map[string]any{"V": "false"}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestStringInvalid", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, Expected: "", Data: map[string]any{"V": "invalid"}}, ExpectedErr: "invalid syntax"},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestToInt(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, Expected: "int-1", Data: map[string]any{"V": 1}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, Expected: "int-1", Data: map[string]any{"V": int32(1)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestFloat64", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, Expected: "int-1", Data: map[string]any{"V": float64(1.42)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestBool", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, Expected: "int-1", Data: map[string]any{"V": true}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, Expected: "int-1", Data: map[string]any{"V": "1"}}, ExpectedErr: ""},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestToInt64(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "int64-1", Data: map[string]any{"V": 1}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "int64-1", Data: map[string]any{"V": int32(1)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestFloat64", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "int64-1", Data: map[string]any{"V": float64(1.42)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestBool", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "int64-1", Data: map[string]any{"V": true}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "int64-1", Data: map[string]any{"V": "1"}}, ExpectedErr: ""},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestToUint(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, Expected: "uint-1", Data: map[string]any{"V": 1}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, Expected: "uint-1", Data: map[string]any{"V": int32(1)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestFloat64", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, Expected: "uint-1", Data: map[string]any{"V": float64(1.42)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestBool", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, Expected: "uint-1", Data: map[string]any{"V": true}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, Expected: "uint-1", Data: map[string]any{"V": "1"}}, ExpectedErr: ""},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestToUint64(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "uint64-1", Data: map[string]any{"V": 1}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "uint64-1", Data: map[string]any{"V": int32(1)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestFloat64", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "uint64-1", Data: map[string]any{"V": float64(1.42)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestBool", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "uint64-1", Data: map[string]any{"V": true}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "uint64-1", Data: map[string]any{"V": "1"}}, ExpectedErr: ""},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestToFloat64(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "float64-1", Data: map[string]any{"V": 1}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "float64-1", Data: map[string]any{"V": int32(1)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestFloat64", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "float64-1.42", Data: map[string]any{"V": float64(1.42)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestBool", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "float64-1", Data: map[string]any{"V": true}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, Expected: "float64-1", Data: map[string]any{"V": "1"}}, ExpectedErr: ""},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestToOctal(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, Expected: "int64-511", Data: map[string]any{"V": 777}}},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, Expected: "int64-504", Data: map[string]any{"V": int32(770)}}},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, Expected: "int64-1", Data: map[string]any{"V": "1"}}},
		{TestCase: pesticide.TestCase{Name: "TestInvalid", Input: `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, Expected: "", Data: map[string]any{"V": 1.1}}, ExpectedErr: "invalid syntax"},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

type testStringer struct{}

func (s testStringer) String() string {
	return "stringer"
}

func TestToString(t *testing.T) {

	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, Expected: "string-1", Data: map[string]any{"V": 1}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, Expected: "string-1", Data: map[string]any{"V": int32(1)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestFloat64", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, Expected: "string-1.42", Data: map[string]any{"V": float64(1.42)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestBool", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, Expected: "string-true", Data: map[string]any{"V": true}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, Expected: "string-1", Data: map[string]any{"V": "1"}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestError", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, Expected: "string-error", Data: map[string]any{"V": fmt.Errorf("error")}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestStringer", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, Expected: "string-stringer", Data: map[string]any{"V": testStringer{}}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestSliceOfBytes", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, Expected: "string-abc", Data: map[string]any{"V": []byte("abc")}}, ExpectedErr: ""},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestToDate(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestDate",
				Input:    `{{$v := toDate "2006-01-02" .V }}{{typeOf $v}}-{{$v}}`,
				Expected: "time.Time-2024-05-09 00:00:00 +0000 UTC",
				Data:     map[string]any{"V": "2024-05-09"},
			},
			ExpectedErr: "",
		},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestToDuration(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Name: "TestInt", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, Expected: "time.Duration-1ns", Data: map[string]any{"V": 1}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInt32", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, Expected: "time.Duration-1µs", Data: map[string]any{"V": int32(1000)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestFloat64", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, Expected: "time.Duration-1.00042ms", Data: map[string]any{"V": float64(1000 * 1000.42)}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestString", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, Expected: "time.Duration-1m0s", Data: map[string]any{"V": "1m"}}, ExpectedErr: ""},
		{TestCase: pesticide.TestCase{Name: "TestInvalid", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, Expected: "", Data: map[string]any{"V": "aaaa"}}, ExpectedErr: "invalid duration \"aaaans\""},
		{TestCase: pesticide.TestCase{Name: "TestCallingOnIt", Input: `{{ (toDuration "1h30m").Seconds }}`, Expected: "5400"}, ExpectedErr: ""},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}

func TestMustToDate(t *testing.T) {
	var tc = []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestDate",
				Input:    `{{$v := mustToDate "2006-01-02 15:04:05 MST" .V }}{{typeOf $v}}-{{$v}}`,
				Expected: "time.Time-2024-05-09 00:00:00 +0000 UTC",
				Data:     map[string]any{"V": "2024-05-09 00:00:00 UTC"},
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestInvalidValue",
				Input:    `{{$v := mustToDate "2006-01-02" .V }}{{typeOf $v}}-{{$v}}`,
				Expected: "",
				Data:     map[string]any{"V": ""},
			},
			ExpectedErr: "cannot parse \"\" as \"2006\"",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestInvalidLayout",
				Input:    `{{$v := mustToDate "invalid" .V }}{{typeOf $v}}-{{$v}}`,
				Expected: "",
				Data:     map[string]any{"V": "2024-05-09"},
			},
			ExpectedErr: "cannot parse \"2024-05-09\" as \"invalid\"",
		},
	}

	pesticide.RunMustTestCases(t, conversion.NewRegistry(), tc)
}
