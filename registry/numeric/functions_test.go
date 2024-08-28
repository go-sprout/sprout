package numeric_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/numeric"
)

func TestFloor(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ floor 1.5 }}`, ExpectedOutput: "1"},
		{Input: `{{ floor 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ floor -1.5 }}`, ExpectedOutput: "-2"},
		{Input: `{{ floor -1 }}`, ExpectedOutput: "-1"},
		{Input: `{{ floor 0 }}`, ExpectedOutput: "0"},
		{Input: `{{ floor 123 }}`, ExpectedOutput: "123"},
		{Input: `{{ floor "123" }}`, ExpectedOutput: "123"},
		{Input: `{{ floor "123.9999" }}`, ExpectedOutput: "123"},
		{Input: `{{ floor 123.0001 }}`, ExpectedOutput: "123"},
		{Input: `{{ floor "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestCeil(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ ceil 1.5 }}`, ExpectedOutput: "2"},
		{Input: `{{ ceil 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ ceil -1.5 }}`, ExpectedOutput: "-1"},
		{Input: `{{ ceil -1 }}`, ExpectedOutput: "-1"},
		{Input: `{{ ceil 0 }}`, ExpectedOutput: "0"},
		{Input: `{{ ceil 123 }}`, ExpectedOutput: "123"},
		{Input: `{{ ceil "123" }}`, ExpectedOutput: "123"},
		{Input: `{{ ceil "123.9999" }}`, ExpectedOutput: "124"},
		{Input: `{{ ceil 123.0001 }}`, ExpectedOutput: "124"},
		{Input: `{{ ceil "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestRound(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ round 3.746 2 }}`, ExpectedOutput: "3.75"},
		{Input: `{{ round 3.746 2 0.5 }}`, ExpectedOutput: "3.75"},
		{Input: `{{ round 123.5555 3 }}`, ExpectedOutput: "123.556"},
		{Input: `{{ round "123.5555" 3 }}`, ExpectedOutput: "123.556"},
		{Input: `{{ round 123.500001 0 }}`, ExpectedOutput: "124"},
		{Input: `{{ round 123.49999999 0 }}`, ExpectedOutput: "123"},
		{Input: `{{ round 123.2329999 2 .3 }}`, ExpectedOutput: "123.23"},
		{Input: `{{ round 123.233 2 .3 }}`, ExpectedOutput: "123.24"},
		{Input: `{{ round "a" 2 }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAdd(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ add }}`, ExpectedOutput: "0"},
		{Input: `{{ add 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ add 1 2 3 4 5 6 7 8 9 10 }}`, ExpectedOutput: "55"},
		{Input: `{{ 10.1 | add 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 }}`, ExpectedOutput: "59.6"},
		{Input: `{{ add 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAddf(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ addf }}`, ExpectedOutput: "0"},
		{Input: `{{ addf 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ addf 1 2 3 4 5 6 7 8 9 10 }}`, ExpectedOutput: "55"},
		{Input: `{{ 10.1 | addf 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 }}`, ExpectedOutput: "59.6"},
		{Input: `{{ addf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAdd1(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ add1 -1 }}`, ExpectedOutput: "0"},
		{Input: `{{ add1f -1.0}}`, ExpectedOutput: "0"},
		{Input: `{{ add1 1 }}`, ExpectedOutput: "2"},
		{Input: `{{ add1 1.1 }}`, ExpectedOutput: "2.1"},
		{Input: `{{ add1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAdd1f(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ add1f -1 }}`, ExpectedOutput: "0"},
		{Input: `{{ add1f -1.0}}`, ExpectedOutput: "0"},
		{Input: `{{ add1f 1 }}`, ExpectedOutput: "2"},
		{Input: `{{ add1f 1.1 }}`, ExpectedOutput: "2.1"},
		{Input: `{{ add1f "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSub(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ sub 1 1 }}`, ExpectedOutput: "0"},
		{Input: `{{ sub 1 2 }}`, ExpectedOutput: "-1"},
		{Input: `{{ sub 1.1 1.1 }}`, ExpectedOutput: "0"},
		{Input: `{{ sub 1.1 2.2 }}`, ExpectedOutput: "-1.1"},
		{Input: `{{ 3 | sub 14 }}`, ExpectedOutput: "11"},
		{Input: `{{ sub 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSubf(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ subf 1.1 1.1 }}`, ExpectedOutput: "0"},
		{Input: `{{ subf 1.1 2.2 }}`, ExpectedOutput: "-1.1"},
		{Input: `{{ round (3 | subf 4.5 1) 1 }}`, ExpectedOutput: "0.5"},
		{Input: `{{ subf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMulInt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ mul 1 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ mul 1 2 }}`, ExpectedOutput: "2"},
		{Input: `{{ mul 1.1 1.1 }}`, ExpectedOutput: "1"},
		{Input: `{{ mul 1.1 2.2 }}`, ExpectedOutput: "2"},
		{Input: `{{ 3 | mul 14 }}`, ExpectedOutput: "42"},
		{Input: `{{ mul 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMulFloat(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ round (mulf 1.1 1.1) 2 }}`, ExpectedOutput: "1.21"},
		{Input: `{{ round (mulf 1.1 2.2) 2 }}`, ExpectedOutput: "2.42"},
		{Input: `{{ round (3.3 | mulf 14.4) 2 }}`, ExpectedOutput: "47.52"},
		{Input: `{{ mulf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestDivInt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ div 1 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ div 1 2 }}`, ExpectedOutput: "0"},
		{Input: `{{ div 1.1 1.1 }}`, ExpectedOutput: "1"},
		{Input: `{{ div 1.1 2.2 }}`, ExpectedOutput: "0"},
		{Input: `{{ 4 | div 5 }}`, ExpectedOutput: "1"},
		{Input: `{{ div 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestDivFloat(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ round (divf 1.1 1.1) 2 }}`, ExpectedOutput: "1"},
		{Input: `{{ round (divf 1.1 2.2) 2 }}`, ExpectedOutput: "0.5"},
		{Input: `{{ 2 | divf 5 4 }}`, ExpectedOutput: "0.625"},
		{Input: `{{ divf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMod(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ mod 10 4 }}`, ExpectedOutput: "2"},
		{Input: `{{ mod 10 3 }}`, ExpectedOutput: "1"},
		{Input: `{{ mod 10 2 }}`, ExpectedOutput: "0"},
		{Input: `{{ mod 10 1 }}`, ExpectedOutput: "0"},
		{Input: `{{ mod 10 0.5 }}`, ExpectedOutput: "0"},
		{Input: `{{ mod "a" 10 }}`, ExpectedErr: "failed to convert: a to float64"},
		{Input: `{{ mod 10 "b" }}`, ExpectedErr: "failed to convert: b to float64"},
		// In case of division by zero, the result is NaN defined by the
		// IEEE 754 " not-a-number" value.
		{Input: `{{ mod 10 0 }}`, ExpectedOutput: strconv.Itoa(int(math.NaN()))},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMin(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ min 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ min 1 "1" }}`, ExpectedOutput: "1"},
		{Input: `{{ min -1 0 1 }}`, ExpectedOutput: "-1"},
		{Input: `{{ min 1 2 3 4 5 6 7 8 9 10 1 2 3 4 5 6 7 8 9 10 0 }}`, ExpectedOutput: "0"},
		{Input: `{{ min "a" "b" }}`, ExpectedErr: "failed to convert: a to int64"},
		{Input: `{{ min 1 "b" }}`, ExpectedErr: "failed to convert: b to int64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMinf(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ minf 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ minf 1 "1.1" }}`, ExpectedOutput: "1"},
		{Input: `{{ minf -1.4 .0 2.1 }}`, ExpectedOutput: "-1.4"},
		{Input: `{{ minf .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 .1 .2 .3 .4 .5 .6 .7 .8 .9 .10}}`, ExpectedOutput: "0.1"},
		{Input: `{{ minf 1 "b" }}`, ExpectedErr: "failed to convert: b to float64"},
		{Input: `{{ minf "a" "b" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMax(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ max 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ max 1 "1" }}`, ExpectedOutput: "1"},
		{Input: `{{ max -1 0 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ max 1 2 3 4 5 6 7 8 9 10 1 2 3 4 5 6 7 8 9 10 0 }}`, ExpectedOutput: "10"},
		{Input: `{{ max 1 "b" }}`, ExpectedErr: "failed to convert: b to int64"},
		{Input: `{{ max "a" "b" }}`, ExpectedErr: "failed to convert: a to int64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMaxf(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ maxf 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ maxf 1.0 "1.1" }}`, ExpectedOutput: "1.1"},
		{Input: `{{ maxf -1.5 0 1.4 }}`, ExpectedOutput: "1.4"},
		{Input: `{{ maxf .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 }}`, ExpectedOutput: "0.9"},
		{Input: `{{ maxf 1 "b" }}`, ExpectedErr: "failed to convert: b to float64"},
		{Input: `{{ maxf "a" "b" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}
