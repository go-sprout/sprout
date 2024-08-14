package semver_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/semver"
)

func TestSemver(t *testing.T) {

	var tc = []pesticide.TestCase{
		{Input: `{{ semver "1.0.0" }}`, Expected: "1.0.0"},
		{Input: `{{ semver "1.0.0-alpha" }}`, Expected: "1.0.0-alpha"},
		{Input: `{{ semver "1.0.0-alpha.1" }}`, Expected: "1.0.0-alpha.1"},
		{Input: `{{ semver "1.0.0-alpha.1+build" }}`, Expected: "1.0.0-alpha.1+build"},
	}

	pesticide.RunTestCases(t, semver.NewRegistry(), tc)
}

func TestSemverCompare(t *testing.T) {

	var tc = []pesticide.TestCase{
		{Input: `{{ semverCompare "1.0.0" "1.0.0" }}`, Expected: "true"},
		{Input: `{{ semverCompare "1.0.0" "1.0.1" }}`, Expected: "false"},
		{Input: `{{ semverCompare "1.0.1" "1.0.0" }}`, Expected: "false"},
		{Input: `{{ semverCompare "~1.0.0" "1.0.0" }}`, Expected: "true"},
		{Input: `{{ semverCompare ">=1.0.0" "1.0.0-alpha" }}`, Expected: "false"},
		{Input: `{{ semverCompare ">1.0.0-alpha" "1.0.0-alpha.1" }}`, Expected: "true"},
		{Input: `{{ semverCompare "1.0.0-alpha.1" "1.0.0-alpha" }}`, Expected: "false"},
		{Input: `{{ semverCompare "1.0.0-alpha.1" "1.0.0-alpha.1" }}`, Expected: "true"},
	}

	pesticide.RunTestCases(t, semver.NewRegistry(), tc)

	var mtc = []pesticide.MustTestCase{
		{
			TestCase:    pesticide.TestCase{Input: `{{ semverCompare "abc" "1.0.0" }}`},
			ExpectedErr: "improper constraint",
		},
		{
			TestCase:    pesticide.TestCase{Input: `{{ semverCompare "1.0.0" "abc" }}`},
			ExpectedErr: "Invalid Semantic Version",
		},
	}

	pesticide.RunMustTestCases(t, semver.NewRegistry(), mtc)
}
