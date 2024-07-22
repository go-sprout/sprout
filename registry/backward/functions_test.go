package backward_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/backward"
)

func TestFail(t *testing.T) {

	var tc = []pesticide.MustTestCase{
		{TestCase: pesticide.TestCase{Input: `{{fail "This is an error"}}`}, ExpectedErr: "This is an error"},
	}

	pesticide.RunMustTestCases(t, backward.NewRegistry(), tc)
}
