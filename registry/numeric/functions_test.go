package numeric_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/numeric"
)

func TestFloor(t *testing.T) {
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
		{Input: `{{ add }}`, ExpectedOutput: "0"},
		{Input: `{{ add 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ add 1 2 3 4 5 6 7 8 9 10 }}`, ExpectedOutput: "55"},
		{Input: `{{ 10.1 | add 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 }}`, ExpectedOutput: "59.6"},
		{Input: `{{ add 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAddf(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ addf }}`, ExpectedOutput: "0"},
		{Input: `{{ addf 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ addf 1 2 3 4 5 6 7 8 9 10 }}`, ExpectedOutput: "55"},
		{Input: `{{ 10.1 | addf 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 }}`, ExpectedOutput: "59.6"},
		{Input: `{{ addf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAdd1(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ add1 -1 }}`, ExpectedOutput: "0"},
		{Input: `{{ add1f -1.0}}`, ExpectedOutput: "0"},
		{Input: `{{ add1 1 }}`, ExpectedOutput: "2"},
		{Input: `{{ add1 1.1 }}`, ExpectedOutput: "2.1"},
		{Input: `{{ add1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAdd1f(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ add1f -1 }}`, ExpectedOutput: "0"},
		{Input: `{{ add1f -1.0}}`, ExpectedOutput: "0"},
		{Input: `{{ add1f 1 }}`, ExpectedOutput: "2"},
		{Input: `{{ add1f 1.1 }}`, ExpectedOutput: "2.1"},
		{Input: `{{ add1f "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSub(t *testing.T) {
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
		{Input: `{{ subf 1.1 1.1 }}`, ExpectedOutput: "0"},
		{Input: `{{ subf 1.1 2.2 }}`, ExpectedOutput: "-1.1"},
		{Input: `{{ round (3 | subf 4.5 1) 1 }}`, ExpectedOutput: "0.5"},
		{Input: `{{ subf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMulInt(t *testing.T) {
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
		{Input: `{{ round (mulf 1.1 1.1) 2 }}`, ExpectedOutput: "1.21"},
		{Input: `{{ round (mulf 1.1 2.2) 2 }}`, ExpectedOutput: "2.42"},
		{Input: `{{ round (3.3 | mulf 14.4) 2 }}`, ExpectedOutput: "47.52"},
		{Input: `{{ mulf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestDivInt(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ div 1 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ div 1 2 }}`, ExpectedOutput: "0"},
		{Input: `{{ div 1.1 1.1 }}`, ExpectedOutput: "1"},
		{Input: `{{ div 1.1 2.2 }}`, ExpectedOutput: "0"},
		{Input: `{{ 4 | div 5 }}`, ExpectedOutput: "1"},
		{Input: `{{ div 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
		{Name: "TestDivideByZero", Input: `{{ div 1 0 }}`, ExpectedErr: "cannot divide by zero"},
		{Name: "TestDivideByZeroInChain", Input: `{{ div 100 5 0 }}`, ExpectedErr: "cannot divide by zero"},
		{Name: "TestZeroDividend", Input: `{{ div 0 5 }}`, ExpectedOutput: "0"},
		{Name: "TestWithoutArgument", Input: `{{ div }}`, ExpectedOutput: "0"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestDivFloat(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ round (divf 1.1 1.1) 2 }}`, ExpectedOutput: "1"},
		{Input: `{{ round (divf 1.1 2.2) 2 }}`, ExpectedOutput: "0.5"},
		{Input: `{{ 2 | divf 5 4 }}`, ExpectedOutput: "0.625"},
		{Input: `{{ divf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMod(t *testing.T) {
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
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
	tc := []pesticide.TestCase{
		{Input: `{{ maxf 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ maxf 1.0 "1.1" }}`, ExpectedOutput: "1.1"},
		{Input: `{{ maxf -1.5 0 1.4 }}`, ExpectedOutput: "1.4"},
		{Input: `{{ maxf .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 }}`, ExpectedOutput: "0.9"},
		{Input: `{{ maxf 1 "b" }}`, ExpectedErr: "failed to convert: b to float64"},
		{Input: `{{ maxf "a" "b" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSum(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ sum }}`, ExpectedOutput: "0"},
		{Input: `{{ sum 1 }}`, ExpectedOutput: "1"},
		{Input: `{{ sum 1 2 3 4 5 6 7 8 9 10 }}`, ExpectedOutput: "55"},
		{Input: `{{ sum -5 5 }}`, ExpectedOutput: "0"},
		{Input: `{{ sum -1 -2 -3 }}`, ExpectedOutput: "-6"},
		{Input: `{{ sum "1" "2" }}`, ExpectedOutput: "3"},
		{Input: `{{ sum 1.9 1.9 }}`, ExpectedOutput: "3"},
		{Input: `{{ sum .V }}`, Data: map[string]any{"V": []int{1, 2, 3}}, ExpectedOutput: "6"},
		{Input: `{{ sum .V }}`, Data: map[string]any{"V": []any{1, "2", 3.0}}, ExpectedOutput: "6"},
		{Input: `{{ sum .V }}`, Data: map[string]any{"V": []string{"10", "20"}}, ExpectedOutput: "30"},
		{Input: `{{ sum .V }}`, Data: map[string]any{"V": []int{}}, ExpectedOutput: "0"},
		{Input: `{{ sum .V }}`, Data: map[string]any{"V": [3]int{1, 2, 3}}, ExpectedOutput: "6"},
		{Input: `{{ sum 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
		{Input: `{{ sum .V }}`, Data: map[string]any{"V": []any{1, "a"}}, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSumf(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ sumf }}`, ExpectedOutput: "0"},
		{Input: `{{ sumf 1.5 }}`, ExpectedOutput: "1.5"},
		{Input: `{{ sumf 1.5 2.25 }}`, ExpectedOutput: "3.75"},
		{Input: `{{ sumf 0.1 0.2 }}`, ExpectedOutput: "0.3"},
		{Input: `{{ sumf 1.1 2.2 3.3 }}`, ExpectedOutput: "6.6"},
		{Input: `{{ sumf -1.5 1.5 }}`, ExpectedOutput: "0"},
		{Input: `{{ sumf "1.1" "2.2" }}`, ExpectedOutput: "3.3"},
		{Input: `{{ sumf .V }}`, Data: map[string]any{"V": []float64{1.1, 2.2}}, ExpectedOutput: "3.3"},
		{Input: `{{ sumf .V }}`, Data: map[string]any{"V": []any{"0.1", "0.2"}}, ExpectedOutput: "0.3"},
		{Input: `{{ sumf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMean(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ mean 10 20 60 }}`, ExpectedOutput: "30"},
		{Input: `{{ mean 5 }}`, ExpectedOutput: "5"},
		{Input: `{{ mean 1 2 }}`, ExpectedOutput: "1"},
		{Input: `{{ mean -10 10 }}`, ExpectedOutput: "0"},
		{Input: `{{ mean -3 -4 }}`, ExpectedOutput: "-3"},
		{Input: `{{ mean "10" "20" }}`, ExpectedOutput: "15"},
		{Input: `{{ mean .V }}`, Data: map[string]any{"V": []int{2, 4, 6}}, ExpectedOutput: "4"},
		{Input: `{{ mean }}`, ExpectedErr: "cannot compute the mean of no values"},
		{Input: `{{ mean .V }}`, Data: map[string]any{"V": []int{}}, ExpectedErr: "cannot compute the mean of no values"},
		{Input: `{{ mean 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMeanf(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ meanf 1 2 }}`, ExpectedOutput: "1.5"},
		{Input: `{{ meanf 10 20 60 }}`, ExpectedOutput: "30"},
		{Input: `{{ meanf 2.5 }}`, ExpectedOutput: "2.5"},
		{Input: `{{ meanf 0.1 0.2 }}`, ExpectedOutput: "0.15"},
		{Input: `{{ meanf 1 2 3 }}`, ExpectedOutput: "2"},
		{Input: `{{ meanf -10 10 }}`, ExpectedOutput: "0"},
		{Input: `{{ meanf .V }}`, Data: map[string]any{"V": []float64{1.5, 2.5}}, ExpectedOutput: "2"},
		{Input: `{{ meanf }}`, ExpectedErr: "cannot compute the mean of no values"},
		{Input: `{{ meanf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMedian(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ median 60 10 20 }}`, ExpectedOutput: "20"},
		{Input: `{{ median 5 }}`, ExpectedOutput: "5"},
		{Input: `{{ median 1 2 }}`, ExpectedOutput: "1"},
		{Input: `{{ median 10 20 30 40 }}`, ExpectedOutput: "25"},
		{Input: `{{ median 3 1 2 }}`, ExpectedOutput: "2"},
		{Input: `{{ median -5 -1 -3 }}`, ExpectedOutput: "-3"},
		{Input: `{{ median .V }}`, Data: map[string]any{"V": []int{7, 1, 3}}, ExpectedOutput: "3"},
		{Input: `{{ median }}`, ExpectedErr: "cannot compute the median of no values"},
		{Input: `{{ median 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMedianf(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ medianf 60 10 20 }}`, ExpectedOutput: "20"},
		{Input: `{{ medianf 10 20 30 40 }}`, ExpectedOutput: "25"},
		{Input: `{{ medianf 1 2 }}`, ExpectedOutput: "1.5"},
		{Input: `{{ medianf 1.5 2.5 3.5 }}`, ExpectedOutput: "2.5"},
		{Input: `{{ medianf 0.1 0.2 }}`, ExpectedOutput: "0.15"},
		{Input: `{{ medianf .V }}`, Data: map[string]any{"V": []float64{4, 1, 2, 3}}, ExpectedOutput: "2.5"},
		{Input: `{{ medianf }}`, ExpectedErr: "cannot compute the median of no values"},
		{Input: `{{ medianf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ mode 25 50 25 }}`, ExpectedOutput: "25"},
		{Input: `{{ mode 5 }}`, ExpectedOutput: "5"},
		{Input: `{{ mode 1 2 3 }}`, ExpectedOutput: "1"},
		{Input: `{{ mode 5 3 5 3 }}`, ExpectedOutput: "5"},
		{Input: `{{ mode 3 5 3 5 }}`, ExpectedOutput: "3"},
		{Input: `{{ mode -1 -1 2 }}`, ExpectedOutput: "-1"},
		{Input: `{{ mode .V }}`, Data: map[string]any{"V": []int{9, 9, 1}}, ExpectedOutput: "9"},
		{Input: `{{ mode }}`, ExpectedErr: "cannot compute the mode of no values"},
		{Input: `{{ mode 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestModef(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ modef 2.5 1.75 2.5 }}`, ExpectedOutput: "2.5"},
		{Input: `{{ modef 1.5 }}`, ExpectedOutput: "1.5"},
		{Input: `{{ modef 1.1 2.2 1.1 }}`, ExpectedOutput: "1.1"},
		{Input: `{{ modef .V }}`, Data: map[string]any{"V": []float64{0.5, 0.5, 9}}, ExpectedOutput: "0.5"},
		{Input: `{{ modef }}`, ExpectedErr: "cannot compute the mode of no values"},
		{Input: `{{ modef 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSpread(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ spread 10 60 20 }}`, ExpectedOutput: "50"},
		{Input: `{{ spread 5 }}`, ExpectedOutput: "0"},
		{Input: `{{ spread -10 10 }}`, ExpectedOutput: "20"},
		{Input: `{{ spread -5 -1 }}`, ExpectedOutput: "4"},
		{Input: `{{ spread 1 1 1 }}`, ExpectedOutput: "0"},
		{Input: `{{ spread .V }}`, Data: map[string]any{"V": []int{3, 9, 1}}, ExpectedOutput: "8"},
		{Input: `{{ spread }}`, ExpectedErr: "cannot compute the spread of no values"},
		{Input: `{{ spread 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSpreadf(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ spreadf 1.5 4.25 }}`, ExpectedOutput: "2.75"},
		{Input: `{{ spreadf 7 }}`, ExpectedOutput: "0"},
		{Input: `{{ spreadf -1.5 1.5 }}`, ExpectedOutput: "3"},
		{Input: `{{ spreadf 0.1 0.3 }}`, ExpectedOutput: "0.2"},
		{Input: `{{ spreadf .V }}`, Data: map[string]any{"V": []float64{2.5, 0.5}}, ExpectedOutput: "2"},
		{Input: `{{ spreadf }}`, ExpectedErr: "cannot compute the spread of no values"},
		{Input: `{{ spreadf 1 "a" }}`, ExpectedErr: "failed to convert: a to float64"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}
