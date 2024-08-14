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
		{Input: `{{ floor 1.5 }}`, Expected: "1"},
		{Input: `{{ floor 1 }}`, Expected: "1"},
		{Input: `{{ floor -1.5 }}`, Expected: "-2"},
		{Input: `{{ floor -1 }}`, Expected: "-1"},
		{Input: `{{ floor 0 }}`, Expected: "0"},
		{Input: `{{ floor 123 }}`, Expected: "123"},
		{Input: `{{ floor "123" }}`, Expected: "123"},
		{Input: `{{ floor 123.9999 }}`, Expected: "123"},
		{Input: `{{ floor 123.0001 }}`, Expected: "123"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestCeil(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ ceil 1.5 }}`, Expected: "2"},
		{Input: `{{ ceil 1 }}`, Expected: "1"},
		{Input: `{{ ceil -1.5 }}`, Expected: "-1"},
		{Input: `{{ ceil -1 }}`, Expected: "-1"},
		{Input: `{{ ceil 0 }}`, Expected: "0"},
		{Input: `{{ ceil 123 }}`, Expected: "123"},
		{Input: `{{ ceil "123" }}`, Expected: "123"},
		{Input: `{{ ceil 123.9999 }}`, Expected: "124"},
		{Input: `{{ ceil 123.0001 }}`, Expected: "124"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestRound(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ round 3.746 2 }}`, Expected: "3.75"},
		{Input: `{{ round 3.746 2 0.5 }}`, Expected: "3.75"},
		{Input: `{{ round 123.5555 3 }}`, Expected: "123.556"},
		{Input: `{{ round "123.5555" 3 }}`, Expected: "123.556"},
		{Input: `{{ round 123.500001 0 }}`, Expected: "124"},
		{Input: `{{ round 123.49999999 0 }}`, Expected: "123"},
		{Input: `{{ round 123.2329999 2 .3 }}`, Expected: "123.23"},
		{Input: `{{ round 123.233 2 .3 }}`, Expected: "123.24"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAdd(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ add }}`, Expected: "0"},
		{Input: `{{ add 1 }}`, Expected: "1"},
		{Input: `{{ add 1 2 3 4 5 6 7 8 9 10 }}`, Expected: "55"},
		{Input: `{{ 10.1 | add 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 }}`, Expected: "59.6"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAddf(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ addf }}`, Expected: "0"},
		{Input: `{{ addf 1 }}`, Expected: "1"},
		{Input: `{{ addf 1 2 3 4 5 6 7 8 9 10 }}`, Expected: "55"},
		{Input: `{{ 10.1 | addf 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 }}`, Expected: "59.6"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAdd1(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ add1 -1 }}`, Expected: "0"},
		{Input: `{{ add1f -1.0}}`, Expected: "0"},
		{Input: `{{ add1 1 }}`, Expected: "2"},
		{Input: `{{ add1 1.1 }}`, Expected: "2.1"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestAdd1f(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ add1f -1 }}`, Expected: "0"},
		{Input: `{{ add1f -1.0}}`, Expected: "0"},
		{Input: `{{ add1f 1 }}`, Expected: "2"},
		{Input: `{{ add1f 1.1 }}`, Expected: "2.1"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSub(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ sub 1 1 }}`, Expected: "0"},
		{Input: `{{ sub 1 2 }}`, Expected: "-1"},
		{Input: `{{ sub 1.1 1.1 }}`, Expected: "0"},
		{Input: `{{ sub 1.1 2.2 }}`, Expected: "-1.1"},
		{Input: `{{ 3 | sub 14 }}`, Expected: "11"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestSubf(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ subf 1.1 1.1 }}`, Expected: "0"},
		{Input: `{{ subf 1.1 2.2 }}`, Expected: "-1.1"},
		{Input: `{{ round (3 | subf 4.5 1) 1 }}`, Expected: "0.5"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMulInt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ mul 1 1 }}`, Expected: "1"},
		{Input: `{{ mul 1 2 }}`, Expected: "2"},
		{Input: `{{ mul 1.1 1.1 }}`, Expected: "1"},
		{Input: `{{ mul 1.1 2.2 }}`, Expected: "2"},
		{Input: `{{ 3 | mul 14 }}`, Expected: "42"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMulFloat(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ round (mulf 1.1 1.1) 2 }}`, Expected: "1.21"},
		{Input: `{{ round (mulf 1.1 2.2) 2 }}`, Expected: "2.42"},
		{Input: `{{ round (3.3 | mulf 14.4) 2 }}`, Expected: "47.52"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestDivInt(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ div 1 1 }}`, Expected: "1"},
		{Input: `{{ div 1 2 }}`, Expected: "0"},
		{Input: `{{ div 1.1 1.1 }}`, Expected: "1"},
		{Input: `{{ div 1.1 2.2 }}`, Expected: "0"},
		{Input: `{{ 4 | div 5 }}`, Expected: "1"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestDivFloat(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ round (divf 1.1 1.1) 2 }}`, Expected: "1"},
		{Input: `{{ round (divf 1.1 2.2) 2 }}`, Expected: "0.5"},
		{Input: `{{ 2 | divf 5 4 }}`, Expected: "0.625"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMod(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ mod 10 4 }}`, Expected: "2"},
		{Input: `{{ mod 10 3 }}`, Expected: "1"},
		{Input: `{{ mod 10 2 }}`, Expected: "0"},
		{Input: `{{ mod 10 1 }}`, Expected: "0"},
		// In case of division by zero, the result is NaN defined by the
		// IEEE 754 " not-a-number" value.
		{Input: `{{ mod 10 0 }}`, Expected: strconv.Itoa(int(math.NaN()))},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMin(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ min 1 }}`, Expected: "1"},
		{Input: `{{ min 1 "1" }}`, Expected: "1"},
		{Input: `{{ min -1 0 1 }}`, Expected: "-1"},
		{Input: `{{ min 1 2 3 4 5 6 7 8 9 10 1 2 3 4 5 6 7 8 9 10 0 }}`, Expected: "0"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMinf(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ minf 1 }}`, Expected: "1"},
		{Input: `{{ minf 1 "1.1" }}`, Expected: "1"},
		{Input: `{{ minf -1.4 .0 2.1 }}`, Expected: "-1.4"},
		{Input: `{{ minf .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 .1 .2 .3 .4 .5 .6 .7 .8 .9 .10}}`, Expected: "0.1"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMax(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ max 1 }}`, Expected: "1"},
		{Input: `{{ max 1 "1" }}`, Expected: "1"},
		{Input: `{{ max -1 0 1 }}`, Expected: "1"},
		{Input: `{{ max 1 2 3 4 5 6 7 8 9 10 1 2 3 4 5 6 7 8 9 10 0 }}`, Expected: "10"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}

func TestMaxf(t *testing.T) {
	var tc = []pesticide.TestCase{
		{Input: `{{ maxf 1 }}`, Expected: "1"},
		{Input: `{{ maxf 1.0 "1.1" }}`, Expected: "1.1"},
		{Input: `{{ maxf -1.5 0 1.4 }}`, Expected: "1.4"},
		{Input: `{{ maxf .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 }}`, Expected: "0.9"},
	}

	pesticide.RunTestCases(t, numeric.NewRegistry(), tc)
}
