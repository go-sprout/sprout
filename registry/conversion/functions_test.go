package conversion_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/conversion"
	rtime "github.com/go-sprout/sprout/registry/time"
)

func TestToBool(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestBool", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "bool-true", Data: map[string]any{"V": true}},
		{Name: "TestInt", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "bool-true", Data: map[string]any{"V": 1}},
		{Name: "TestInt32", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "bool-true", Data: map[string]any{"V": int32(1)}},
		{Name: "TestFloat64", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "bool-true", Data: map[string]any{"V": float64(1.42)}},
		{Name: "TestString", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "bool-true", Data: map[string]any{"V": "true"}},
		{Name: "TestStringFalse", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "bool-false", Data: map[string]any{"V": "false"}},
		{Name: "TestStringInvalid", Input: `{{$v := toBool .V }}{{kindOf $v}}-{{$v}}`, ExpectedErr: "invalid syntax", Data: map[string]any{"V": "invalid"}},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

func TestToInt(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestInt", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "int-1", Data: map[string]any{"V": 1}},
		{Name: "TestInt32", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "int-1", Data: map[string]any{"V": int32(1)}},
		{Name: "TestFloat64", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "int-1", Data: map[string]any{"V": float64(1.42)}},
		{Name: "TestBool", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "int-1", Data: map[string]any{"V": true}},
		{Name: "TestString", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "int-1", Data: map[string]any{"V": "1"}},
		{Name: "TestStringInvalid", Input: `{{$v := toInt .V }}{{kindOf $v}}-{{$v}}`, ExpectedErr: "error calling toInt: unable to cast", Data: map[string]any{"V": "invalid"}},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

func TestToInt64(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestInt", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "int64-1", Data: map[string]any{"V": 1}},
		{Name: "TestInt32", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "int64-1", Data: map[string]any{"V": int32(1)}},
		{Name: "TestFloat64", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "int64-1", Data: map[string]any{"V": float64(1.42)}},
		{Name: "TestBool", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "int64-1", Data: map[string]any{"V": true}},
		{Name: "TestString", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "int64-1", Data: map[string]any{"V": "1"}},
		{Name: "TestStringInvalid", Input: `{{$v := toInt64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedErr: "error calling toInt64: unable to cast", Data: map[string]any{"V": "invalid"}},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

func TestToUint(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestInt", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "uint-1", Data: map[string]any{"V": 1}},
		{Name: "TestInt32", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "uint-1", Data: map[string]any{"V": int32(1)}},
		{Name: "TestFloat64", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "uint-1", Data: map[string]any{"V": float64(1.42)}},
		{Name: "TestBool", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "uint-1", Data: map[string]any{"V": true}},
		{Name: "TestString", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, ExpectedOutput: "uint-1", Data: map[string]any{"V": "1"}},
		{Name: "TestStringInvalid", Input: `{{$v := toUint .V }}{{kindOf $v}}-{{$v}}`, ExpectedErr: "error calling toUint: unable to cast", Data: map[string]any{"V": "invalid"}},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

func TestToUint64(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestInt", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "uint64-1", Data: map[string]any{"V": 1}},
		{Name: "TestInt32", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "uint64-1", Data: map[string]any{"V": int32(1)}},
		{Name: "TestFloat64", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "uint64-1", Data: map[string]any{"V": float64(1.42)}},
		{Name: "TestBool", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "uint64-1", Data: map[string]any{"V": true}},
		{Name: "TestString", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "uint64-1", Data: map[string]any{"V": "1"}},
		{Name: "TestStringInvalid", Input: `{{$v := toUint64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedErr: "error calling toUint64: unable to cast", Data: map[string]any{"V": "invalid"}},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

func TestToFloat64(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestInt", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "float64-1", Data: map[string]any{"V": 1}},
		{Name: "TestInt32", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "float64-1", Data: map[string]any{"V": int32(1)}},
		{Name: "TestFloat64", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "float64-1.42", Data: map[string]any{"V": float64(1.42)}},
		{Name: "TestBool", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "float64-1", Data: map[string]any{"V": true}},
		{Name: "TestString", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "float64-1", Data: map[string]any{"V": "1"}},
		{Name: "TestStringInvalid", Input: `{{$v := toFloat64 .V }}{{typeOf $v}}-{{$v}}`, ExpectedErr: "error calling toFloat64: unable to cast", Data: map[string]any{"V": "invalid"}},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

func TestToOctal(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestInt", Input: `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "int64-511", Data: map[string]any{"V": 777}},
		{Name: "TestInt32", Input: `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "int64-504", Data: map[string]any{"V": int32(770)}},
		{Name: "TestString", Input: `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "int64-1", Data: map[string]any{"V": "1"}},
		{Name: "TestInvalid", Input: `{{$v := toOctal .V }}{{typeOf $v}}-{{$v}}`, ExpectedErr: "failed to parse octal", Data: map[string]any{"V": 1.1}},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

type testStringer struct{}

func (s testStringer) String() string {
	return "stringer"
}

func TestToString(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestInt", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "string-1", Data: map[string]any{"V": 1}},
		{Name: "TestInt32", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "string-1", Data: map[string]any{"V": int32(1)}},
		{Name: "TestFloat64", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "string-1.42", Data: map[string]any{"V": float64(1.42)}},
		{Name: "TestBool", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "string-true", Data: map[string]any{"V": true}},
		{Name: "TestString", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "string-1", Data: map[string]any{"V": "1"}},
		{Name: "TestError", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "string-error", Data: map[string]any{"V": fmt.Errorf("error")}},
		{Name: "TestStringer", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "string-stringer", Data: map[string]any{"V": testStringer{}}},
		{Name: "TestSliceOfBytes", Input: `{{$v := toString .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "string-abc", Data: map[string]any{"V": []byte("abc")}},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

func TestToDate(t *testing.T) {
	t.Run("dates with numeric timezone offset", func(t *testing.T) {
		// Please refer to https://pkg.go.dev/time#Parse
		// When parsing a time with a zone offset like -0700,
		// if the offset corresponds to a time zone used by the current location (Local),
		// then Parse uses that location and zone in the returned time.
		// Otherwise it records the time as being in a fabricated location with time fixed at the given zone offset.

		// So we have to temporarily force time.Local a known timezone
		// to validate the behavior of toDate function
		local, err := time.LoadLocation("America/New_York")
		require.NoError(t, err)

		// temporarily force time.Local to New York
		rtime.ForceTimeLocal(t, local)

		tc := []pesticide.TestCase{
			{
				Name:           "date with UTC timezone",
				Input:          `{{$v := toDate "2006-01-02 15:04:05 -0700" .V }}{{typeOf $v}}-{{$v}}`,
				Data:           map[string]any{"V": "2024-05-09 00:00:00 +0000"},
				ExpectedOutput: "time.Time-2024-05-09 00:00:00 +0000 +0000",
			},
			{
				Name:           "date with non-UTC timezone equal to local timezone",
				Input:          `{{$v := toDate "2006-01-02 15:04:05 -0700" .V }}{{typeOf $v}}-{{$v}}`,
				Data:           map[string]any{"V": "2024-05-09 00:00:00 -0400"},
				ExpectedOutput: "time.Time-2024-05-09 00:00:00 -0400 EDT",
			},
			{
				Name:           "date with non-UTC timezone different than local",
				Input:          `{{$v := toDate "2006-01-02 15:04:05 -0700" .V }}{{typeOf $v}}-{{$v}}`,
				Data:           map[string]any{"V": "2024-05-09 00:00:00 -0700"},
				ExpectedOutput: "time.Time-2024-05-09 00:00:00 -0700 -0700",
			},
		}

		pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
	})

	t.Run("dates with abbreviated timezone", func(t *testing.T) {
		// Please refer to https://pkg.go.dev/time#Parse
		// When parsing a time with a zone abbreviation like MST,
		// if the zone abbreviation has a defined offset in the current location,
		// then that offset is used.
		// The zone abbreviation "UTC" is recognized as UTC regardless of location.
		// To avoid such problems, prefer time layouts that use a numeric zone offset, or use ParseInLocation.

		// So we have to temporarily force time.Local a known timezone
		// to validate the behavior of toDate function
		local, err := time.LoadLocation("America/New_York")
		require.NoError(t, err)

		// temporarily force time.Local to New York
		rtime.ForceTimeLocal(t, local)

		tc := []pesticide.TestCase{
			{
				Name:           "date with UTC timezone",
				Input:          `{{$v := toDate "2006-01-02 15:04:05 MST" .V }}{{typeOf $v}}-{{$v}}`,
				Data:           map[string]any{"V": "2024-05-09 00:00:00 UTC"},
				ExpectedOutput: "time.Time-2024-05-09 00:00:00 +0000 UTC",
			},
			{
				Name:           "date with non-UTC timezone equal to local timezone",
				Input:          `{{$v := toDate "2006-01-02 15:04:05 MST" .V }}{{typeOf $v}}-{{$v}}`,
				Data:           map[string]any{"V": "2024-05-09 00:00:00 EDT"},
				ExpectedOutput: "time.Time-2024-05-09 00:00:00 -0400 EDT",
			},
			{
				Name:           "date with non-UTC timezone different than local",
				Input:          `{{$v := toDate "2006-01-02 15:04:05 MST" .V }}{{typeOf $v}}-{{$v}}`,
				Data:           map[string]any{"V": "2024-05-09 00:00:00 MST"},
				ExpectedOutput: "time.Time-2024-05-09 00:00:00 +0000 MST",
			},
		}

		pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
	})

	t.Run("dates without timezone (local time should be assumed)", func(t *testing.T) {
		t.Run("UTC", func(t *testing.T) {
			// temporarily force time.Local to UTC
			rtime.ForceTimeLocal(t, time.UTC)

			tc := []pesticide.TestCase{
				{
					Name:           "short date",
					Input:          `{{$v := toDate "2006-01-02" .V }}{{typeOf $v}}-{{$v}}`,
					Data:           map[string]any{"V": "2024-05-09"},
					ExpectedOutput: "time.Time-2024-05-09 00:00:00 +0000 UTC",
				},
				{
					Name:           "datetime ",
					Input:          `{{$v := toDate "2006-01-02 15:04:05" .V }}{{typeOf $v}}-{{$v}}`,
					Data:           map[string]any{"V": "2024-05-09 01:02:03"},
					ExpectedOutput: "time.Time-2024-05-09 01:02:03 +0000 UTC",
				},
			}

			pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
		})

		t.Run("New York timezone", func(t *testing.T) {
			local, err := time.LoadLocation("America/New_York")
			require.NoError(t, err)

			// temporarily force time.Local to New York
			rtime.ForceTimeLocal(t, local)

			tc := []pesticide.TestCase{
				{
					Name:           "short date",
					Input:          `{{$v := toDate "2006-01-02" .V }}{{typeOf $v}}-{{$v}}`,
					Data:           map[string]any{"V": "2024-05-09"},
					ExpectedOutput: "time.Time-2024-05-09 00:00:00 -0400 EDT",
				},
				{
					Name:           "datetime ",
					Input:          `{{$v := toDate "2006-01-02 15:04:05" .V }}{{typeOf $v}}-{{$v}}`,
					Data:           map[string]any{"V": "2024-05-09 01:02:03"},
					ExpectedOutput: "time.Time-2024-05-09 01:02:03 -0400 EDT",
				},
			}

			pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
		})
	})

	t.Run("invalid layout", func(t *testing.T) {
		tc := []pesticide.TestCase{
			{
				Name:        "TestInvalidValue",
				Input:       `{{$v := toDate "2006-01-02" .V }}{{typeOf $v}}-{{$v}}`,
				Data:        map[string]any{"V": ""},
				ExpectedErr: `cannot parse "" as "2006"`,
			},
			{
				Name:        "TestInvalidLayout",
				Input:       `{{$v := toDate "invalid" .V }}{{typeOf $v}}-{{$v}}`,
				Data:        map[string]any{"V": "2024-05-09"},
				ExpectedErr: `cannot parse "2024-05-09" as "invalid"`,
			},
		}

		pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
	})
}

func TestToLocalDate(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestLocalDate",
			Input:          `{{$v := toLocalDate "2006-01-02" "Europe/Paris" .V }}{{typeOf $v}}-{{$v}}`,
			Data:           map[string]any{"V": "2024-05-09"},
			ExpectedOutput: "time.Time-2024-05-09 00:00:00 +0200 CEST",
		},
		{
			Name:           "TestLocalDate",
			Input:          `{{$v := toLocalDate "2006-01-02 15:04:05 MST" "MST" .V }}{{typeOf $v}}-{{$v}}`,
			Data:           map[string]any{"V": "2024-05-09 00:00:00 MST"},
			ExpectedOutput: "time.Time-2024-05-09 00:00:00 -0700 MST",
		},
		{
			Name:        "TestInvalidValue",
			Input:       `{{$v := toLocalDate "2006-01-02" "UTC" .V }}{{typeOf $v}}-{{$v}}`,
			Data:        map[string]any{"V": ""},
			ExpectedErr: "cannot parse \"\" as \"2006\"",
		},
		{
			Name:        "TestInvalidLayout",
			Input:       `{{$v := toLocalDate "invalid" "UTC" .V }}{{typeOf $v}}-{{$v}}`,
			Data:        map[string]any{"V": "2024-05-09"},
			ExpectedErr: "cannot parse \"2024-05-09\" as \"invalid\"",
		},
		{
			Name:        "TestInvalidTimezone",
			Input:       `{{$v := toLocalDate "2006-01-02" "invalid" .V }}{{typeOf $v}}-{{$v}}`,
			Data:        map[string]any{"V": "2024-05-09"},
			ExpectedErr: "unknown time zone invalid",
		},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}

func TestToDuration(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestInt", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "time.Duration-1ns", Data: map[string]any{"V": 1}},
		{Name: "TestInt32", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "time.Duration-1µs", Data: map[string]any{"V": int32(1000)}},
		{Name: "TestFloat64", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "time.Duration-1.00042ms", Data: map[string]any{"V": float64(1000 * 1000.42)}},
		{Name: "TestString", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, ExpectedOutput: "time.Duration-1m0s", Data: map[string]any{"V": "1m"}},
		{Name: "TestInvalid", Input: `{{$v := toDuration .V }}{{typeOf $v}}-{{$v}}`, ExpectedErr: "invalid duration", Data: map[string]any{"V": "aaaa"}},
		{Name: "TestCallingOnIt", Input: `{{ (toDuration "1h30m").Seconds }}`, ExpectedOutput: "5400"},
	}

	pesticide.RunTestCases(t, conversion.NewRegistry(), tc)
}
