package crypto_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/crypto"
)

func TestDerivePassword(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestSimple", Input: `{{ derivePassword 1 "long" "password" "user" "example.com" }}`, ExpectedOutput: "ZedaFaxcZaso9*"},
	}

	pesticide.RunTestCases(t, crypto.NewRegistry(), tc)
}
