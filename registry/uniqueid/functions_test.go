package uniqueid_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/uniqueid"
)

func TestUuidv4(t *testing.T) {
	var tc = []pesticide.RegexpTestCase{
		{Name: "TestUuidv4", Template: `{{uuidv4}}`, Regexp: `^[\da-f]{8}-[\da-f]{4}-4[\da-f]{3}-[\da-f]{4}-[\da-f]{12}$`, Length: 36},
	}

	pesticide.RunRegexpTestCases(t, uniqueid.NewRegistry(), tc)
}
