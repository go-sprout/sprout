package backward_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/backward"
)

func TestFail(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{fail "This is an error"}}`, ExpectedErr: "This is an error"},
	}

	pesticide.RunTestCases(t, backward.NewRegistry(), tc)
}

func TestUrlParse(t *testing.T) {
	tc := []pesticide.TestCase{
		{Input: `{{ urlParse "https://example.com" | urlJoin }}`, ExpectedOutput: "https://example.com"},
		{Input: `{{ urlParse "https://example.com/path" | urlJoin }}`, ExpectedOutput: "https://example.com/path"},
		{Input: `{{ urlParse "https://user:pass@example.com/path?query=1" | urlJoin }}`, ExpectedOutput: "https://user:pass@example.com/path?query=1"},
		{Input: `{{ urlParse "://" }}`, ExpectedErr: "unable to parse url"},
	}

	pesticide.RunTestCases(t, backward.NewRegistry(), tc)
}

func TestGetHostByName(t *testing.T) {
	ipv6 := `^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`
	ipv4 := `^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`

	ipAddressRegexp := ipv4 + `|` + ipv6

	tc := []pesticide.RegexpTestCase{
		{Template: `{{ getHostByName "example.com" }}`, Regexp: ipAddressRegexp, Length: -1},
		{Template: `{{ getHostByName "github.com" }}`, Regexp: ipAddressRegexp, Length: -1},
		{Template: `{{ getHostByName "127.0.0.1" }}`, Regexp: ipAddressRegexp, Length: -1},
	}

	pesticide.RunRegexpTestCases(t, backward.NewRegistry(), tc)
}
