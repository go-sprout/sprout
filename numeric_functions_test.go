package sprout

import (
	"math"
	"strconv"
	"testing"
)

func TestFloor(t *testing.T) {
	var tests = testCases{
		{"", `{{ floor 1.5 }}`, "1", nil},
		{"", `{{ floor 1 }}`, "1", nil},
		{"", `{{ floor -1.5 }}`, "-2", nil},
		{"", `{{ floor -1 }}`, "-1", nil},
		{"", `{{ floor 0 }}`, "0", nil},
		{"", `{{ floor 123 }}`, "123", nil},
		{"", `{{ floor "123" }}`, "123", nil},
		{"", `{{ floor 123.9999 }}`, "123", nil},
		{"", `{{ floor 123.0001 }}`, "123", nil},
	}

	runTestCases(t, tests)
}

func TestCeil(t *testing.T) {
	var tests = testCases{
		{"", `{{ ceil 1.5 }}`, "2", nil},
		{"", `{{ ceil 1 }}`, "1", nil},
		{"", `{{ ceil -1.5 }}`, "-1", nil},
		{"", `{{ ceil -1 }}`, "-1", nil},
		{"", `{{ ceil 0 }}`, "0", nil},
		{"", `{{ ceil 123 }}`, "123", nil},
		{"", `{{ ceil "123" }}`, "123", nil},
		{"", `{{ ceil 123.9999 }}`, "124", nil},
		{"", `{{ ceil 123.0001 }}`, "124", nil},
	}

	runTestCases(t, tests)
}

func TestRound(t *testing.T) {
	var tests = testCases{
		{"", `{{ round 3.746 2 }}`, "3.75", nil},
		{"", `{{ round 3.746 2 0.5 }}`, "3.75", nil},
		{"", `{{ round 123.5555 3 }}`, "123.556", nil},
		{"", `{{ round "123.5555" 3 }}`, "123.556", nil},
		{"", `{{ round 123.500001 0 }}`, "124", nil},
		{"", `{{ round 123.49999999 0 }}`, "123", nil},
		{"", `{{ round 123.2329999 2 .3 }}`, "123.23", nil},
		{"", `{{ round 123.233 2 .3 }}`, "123.24", nil},
	}

	runTestCases(t, tests)
}

func TestAdd(t *testing.T) {
	var tests = testCases{
		{"", `{{ add }}`, "0", nil},
		{"", `{{ add 1 }}`, "1", nil},
		{"", `{{ add 1 2 3 4 5 6 7 8 9 10 }}`, "55", nil},
		{"", `{{ 10.1 | add 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 }}`, "59.6", nil},
	}

	runTestCases(t, tests)
}

func TestAddf(t *testing.T) {
	var tests = testCases{
		{"", `{{ addf }}`, "0", nil},
		{"", `{{ addf 1 }}`, "1", nil},
		{"", `{{ addf 1 2 3 4 5 6 7 8 9 10 }}`, "55", nil},
		{"", `{{ 10.1 | addf 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 }}`, "59.6", nil},
	}

	runTestCases(t, tests)
}

func TestAdd1(t *testing.T) {
	var tests = testCases{
		{"", `{{ add1 -1 }}`, "0", nil},
		{"", `{{ add1f -1.0}}`, "0", nil},
		{"", `{{ add1 1 }}`, "2", nil},
		{"", `{{ add1 1.1 }}`, "2.1", nil},
	}

	runTestCases(t, tests)
}

func TestAdd1f(t *testing.T) {
	var tests = testCases{
		{"", `{{ add1f -1 }}`, "0", nil},
		{"", `{{ add1f -1.0}}`, "0", nil},
		{"", `{{ add1f 1 }}`, "2", nil},
		{"", `{{ add1f 1.1 }}`, "2.1", nil},
	}

	runTestCases(t, tests)
}

func TestSub(t *testing.T) {
	var tests = testCases{
		{"", `{{ sub 1 1 }}`, "0", nil},
		{"", `{{ sub 1 2 }}`, "-1", nil},
		{"", `{{ sub 1.1 1.1 }}`, "0", nil},
		{"", `{{ sub 1.1 2.2 }}`, "-1.1", nil},
		{"", `{{ 3 | sub 14 }}`, "11", nil},
	}

	runTestCases(t, tests)
}

func TestSubf(t *testing.T) {
	var tests = testCases{
		{"", `{{ subf 1.1 1.1 }}`, "0", nil},
		{"", `{{ subf 1.1 2.2 }}`, "-1.1", nil},
		{"", `{{ round (3 | subf 4.5 1) 1 }}`, "0.5", nil},
	}

	runTestCases(t, tests)
}

func TestMulInt(t *testing.T) {
	var tests = testCases{
		{"", `{{ mul 1 1 }}`, "1", nil},
		{"", `{{ mul 1 2 }}`, "2", nil},
		{"", `{{ mul 1.1 1.1 }}`, "1", nil},
		{"", `{{ mul 1.1 2.2 }}`, "2", nil},
		{"", `{{ 3 | mul 14 }}`, "42", nil},
	}

	runTestCases(t, tests)
}

func TestMulFloat(t *testing.T) {
	var tests = testCases{
		{"", `{{ round (mulf 1.1 1.1) 2 }}`, "1.21", nil},
		{"", `{{ round (mulf 1.1 2.2) 2 }}`, "2.42", nil},
		{"", `{{ round (3.3 | mulf 14.4) 2 }}`, "47.52", nil},
	}

	runTestCases(t, tests)
}

func TestDivInt(t *testing.T) {
	var tests = testCases{
		{"", `{{ div 1 1 }}`, "1", nil},
		{"", `{{ div 1 2 }}`, "0", nil},
		{"", `{{ div 1.1 1.1 }}`, "1", nil},
		{"", `{{ div 1.1 2.2 }}`, "0", nil},
		{"", `{{ 4 | div 5 }}`, "1", nil},
	}

	runTestCases(t, tests)
}

func TestDivFloat(t *testing.T) {
	var tests = testCases{
		{"", `{{ round (divf 1.1 1.1) 2 }}`, "1", nil},
		{"", `{{ round (divf 1.1 2.2) 2 }}`, "0.5", nil},
		{"", `{{ 2 | divf 5 4 }}`, "0.625", nil},
	}

	runTestCases(t, tests)
}

func TestMod(t *testing.T) {
	var tests = testCases{
		{"", `{{ mod 10 4 }}`, "2", nil},
		{"", `{{ mod 10 3 }}`, "1", nil},
		{"", `{{ mod 10 2 }}`, "0", nil},
		{"", `{{ mod 10 1 }}`, "0", nil},
		// In case of division by zero, the result is NaN defined by the
		// IEEE 754 " not-a-number" value.
		{"", `{{ mod 10 0 }}`, strconv.Itoa(int(math.NaN())), nil},
	}

	runTestCases(t, tests)
}

func TestMin(t *testing.T) {
	var tests = testCases{
		{"", `{{ min 1 }}`, "1", nil},
		{"", `{{ min 1 "1" }}`, "1", nil},
		{"", `{{ min -1 0 1 }}`, "-1", nil},
		{"", `{{ min 1 2 3 4 5 6 7 8 9 10 1 2 3 4 5 6 7 8 9 10 0 }}`, "0", nil},
	}

	runTestCases(t, tests)
}

func TestMinf(t *testing.T) {
	var tests = testCases{
		{"", `{{ minf 1 }}`, "1", nil},
		{"", `{{ minf 1 "1.1" }}`, "1", nil},
		{"", `{{ minf -1.4 .0 2.1 }}`, "-1.4", nil},
		{"", `{{ minf .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 .1 .2 .3 .4 .5 .6 .7 .8 .9 .10}}`, "0.1", nil},
	}

	runTestCases(t, tests)
}

func TestMax(t *testing.T) {
	var tests = testCases{
		{"", `{{ max 1 }}`, "1", nil},
		{"", `{{ max 1 "1" }}`, "1", nil},
		{"", `{{ max -1 0 1 }}`, "1", nil},
		{"", `{{ max 1 2 3 4 5 6 7 8 9 10 1 2 3 4 5 6 7 8 9 10 0 }}`, "10", nil},
	}

	runTestCases(t, tests)
}

func TestMaxf(t *testing.T) {
	var tests = testCases{
		{"", `{{ maxf 1 }}`, "1", nil},
		{"", `{{ maxf 1.0 "1.1" }}`, "1.1", nil},
		{"", `{{ maxf -1.5 0 1.4 }}`, "1.4", nil},
		{"", `{{ maxf .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 .1 .2 .3 .4 .5 .6 .7 .8 .9 .10 }}`, "0.9", nil},
	}

	runTestCases(t, tests)
}
