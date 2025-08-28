package semver_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/semver"
)

func TestSemver(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ semver "1.0.0" }}`, ExpectedOutput: "1.0.0"},
		{Input: `{{ semver "1.0.0-alpha" }}`, ExpectedOutput: "1.0.0-alpha"},
		{Input: `{{ semver "1.0.0-alpha.1" }}`, ExpectedOutput: "1.0.0-alpha.1"},
		{Input: `{{ semver "1.0.0-alpha.1+build" }}`, ExpectedOutput: "1.0.0-alpha.1+build"},
	}

	pesticide.RunTestCases(t, semver.NewRegistry(), tc)
}

func TestSemverCompare(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ semverCompare "1.0.0" "1.0.0" }}`, ExpectedOutput: "true"},
		{Input: `{{ semverCompare "1.0.0" "1.0.1" }}`, ExpectedOutput: "false"},
		{Input: `{{ semverCompare "1.0.1" "1.0.0" }}`, ExpectedOutput: "false"},
		{Input: `{{ semverCompare "~1.0.0" "1.0.0" }}`, ExpectedOutput: "true"},
		{Input: `{{ semverCompare ">=1.0.0" "1.0.0-alpha" }}`, ExpectedOutput: "false"},
		{Input: `{{ semverCompare ">1.0.0-alpha" "1.0.0-alpha.1" }}`, ExpectedOutput: "true"},
		{Input: `{{ semverCompare "1.0.0-alpha.1" "1.0.0-alpha" }}`, ExpectedOutput: "false"},
		{Input: `{{ semverCompare "1.0.0-alpha.1" "1.0.0-alpha.1" }}`, ExpectedOutput: "true"},
	}

	pesticide.RunTestCases(t, semver.NewRegistry(), tc)

	mtc := []pesticide.TestCase{
		{Input: `{{ semverCompare "abc" "1.0.0" }}`, ExpectedErr: "improper constraint"},
		{Input: `{{ semverCompare "1.0.0" "abc" }}`, ExpectedErr: "invalid semantic version"},
	}

	pesticide.RunTestCases(t, semver.NewRegistry(), mtc)
}
