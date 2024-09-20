package network_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/network"
)

func TestParseIP(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "ValidIPv4", Input: `{{ .V | parseIP }}`, Data: map[string]any{"V": "10.42.0.1"}, ExpectedOutput: "10.42.0.1"},
		{Name: "ValidIPv6", Input: `{{ .V | parseIP }}`, Data: map[string]any{"V": "2001:db8::"}, ExpectedOutput: "2001:db8::"},
		{Name: "InvalidIP", Input: `{{ .V | parseIP }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
		{Name: "EmptyIP", Input: `{{ .V | parseIP }}`, Data: map[string]any{"V": ""}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestParseMAC(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "ValidMAC", Input: `{{ .V | parseMAC }}`, Data: map[string]any{"V": "01:23:45:67:89:ab"}, ExpectedOutput: "01:23:45:67:89:ab"},
		{Name: "InvalidMAC", Input: `{{ .V | parseMAC }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "cannot parse MAC address"},
		{Name: "EmptyMAC", Input: `{{ .V | parseMAC }}`, Data: map[string]any{"V": ""}, ExpectedErr: "cannot parse MAC address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestParseCIDR(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "ValidCIDR", Input: `{{ .V | parseCIDR }}`, Data: map[string]any{"V": "10.42.0.0/24"}, ExpectedOutput: "10.42.0.0/24"},
		{Name: "ValidCIDRWithIP", Input: `{{ .V | parseCIDR }}`, Data: map[string]any{"V": "10.42.0.2/24"}, ExpectedOutput: "10.42.0.0/24"},
		{Name: "InvalidCIDR", Input: `{{ .V | parseCIDR }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid CIDR address"},
		{Name: "EmptyCIDR", Input: `{{ .V | parseCIDR }}`, Data: map[string]any{"V": ""}, ExpectedErr: "invalid CIDR address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestIPVersion(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "IPv4", Input: `{{ .V | ipVersion }}`, Data: map[string]any{"V": "10.42.0.0"}, ExpectedOutput: "4"},
		{Name: "IPv6", Input: `{{ .V | ipVersion }}`, Data: map[string]any{"V": "2001:db8::"}, ExpectedOutput: "6"},
		{Name: "MixedIP", Input: `{{ .V | ipVersion }}`, Data: map[string]any{"V": "2001:db8::10.42.0.1"}, ExpectedOutput: "6"},
		{Name: "InvalidIP", Input: `{{ .V | ipVersion }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
		{Name: "EmptyIP", Input: `{{ .V | ipVersion }}`, Data: map[string]any{"V": ""}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestIPIsLoopback(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "LoopbackIPv4", Input: `{{ .V | ipIsLoopback }}`, Data: map[string]any{"V": "127.0.0.1"}, ExpectedOutput: "true"},
		{Name: "LoopbackIPv6", Input: `{{ .V | ipIsLoopback }}`, Data: map[string]any{"V": "::1"}, ExpectedOutput: "true"},
		{Name: "NonLoopbackIPv4", Input: `{{ .V | ipIsLoopback }}`, Data: map[string]any{"V": "10.42.0.1"}, ExpectedOutput: "false"},
		{Name: "NonLoopbackIPv6", Input: `{{ .V | ipIsLoopback }}`, Data: map[string]any{"V": "2001:db8::"}, ExpectedOutput: "false"},
		{Name: "InvalidIP", Input: `{{ .V | ipIsLoopback }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
		{Name: "EmptyIP", Input: `{{ .V | ipIsLoopback }}`, Data: map[string]any{"V": ""}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestIPIsGlobalUnicast(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "GlobalUnicastIPv4", Input: `{{ .V | ipIsGlobalUnicast }}`, Data: map[string]any{"V": "8.8.8.8"}, ExpectedOutput: "true"},
		{Name: "GlobalUnicastIPv6", Input: `{{ .V | ipIsGlobalUnicast }}`, Data: map[string]any{"V": "2001:4860:4860::8888"}, ExpectedOutput: "true"},
		{Name: "NonGlobalUnicastIPv4", Input: `{{ .V | ipIsGlobalUnicast }}`, Data: map[string]any{"V": "127.0.0.1"}, ExpectedOutput: "false"},
		{Name: "NonGlobalUnicastIPv6", Input: `{{ .V | ipIsGlobalUnicast }}`, Data: map[string]any{"V": "::1"}, ExpectedOutput: "false"},
		{Name: "InvalidIP", Input: `{{ .V | ipIsGlobalUnicast }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
		{Name: "EmptyIP", Input: `{{ .V | ipIsGlobalUnicast }}`, Data: map[string]any{"V": ""}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestIPIsMulticast(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "MulticastIPv4", Input: `{{ .V | ipIsMulticast }}`, Data: map[string]any{"V": "224.0.0.1"}, ExpectedOutput: "true"},
		{Name: "MulticastIPv6", Input: `{{ .V | ipIsMulticast }}`, Data: map[string]any{"V": "ff02::1"}, ExpectedOutput: "true"},
		{Name: "NonMulticastIPv4", Input: `{{ .V | ipIsMulticast }}`, Data: map[string]any{"V": "127.0.0.1"}, ExpectedOutput: "false"},
		{Name: "NonMulticastIPv6", Input: `{{ .V | ipIsMulticast }}`, Data: map[string]any{"V": "::1"}, ExpectedOutput: "false"},
		{Name: "InvalidIP", Input: `{{ .V | ipIsMulticast }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
		{Name: "EmptyIP", Input: `{{ .V | ipIsMulticast }}`, Data: map[string]any{"V": ""}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestIPIsPrivate(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "PrivateIPv4", Input: `{{ .V | ipIsPrivate }}`, Data: map[string]any{"V": "10.42.0.1"}, ExpectedOutput: "true"},
		{Name: "PrivateIPv6", Input: `{{ .V | ipIsPrivate }}`, Data: map[string]any{"V": "fd00::1"}, ExpectedOutput: "true"},
		{Name: "NonPrivateIPv4", Input: `{{ .V | ipIsPrivate }}`, Data: map[string]any{"V": "42.42.42.42"}, ExpectedOutput: "false"},
		{Name: "NonPrivateIPv6", Input: `{{ .V | ipIsPrivate }}`, Data: map[string]any{"V": "2001:db8::"}, ExpectedOutput: "false"},
		{Name: "InvalidIP", Input: `{{ .V | ipIsPrivate }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
		{Name: "EmptyIP", Input: `{{ .V | ipIsPrivate }}`, Data: map[string]any{"V": ""}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestIPIncrement(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "IncrementOverBasic", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "10.42.0.1"}, ExpectedOutput: "10.42.0.2"},
		{Name: "IncrementToMax", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "10.42.254.254"}, ExpectedOutput: "10.42.254.255"},
		{Name: "IncrementOverMax", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "10.42.254.255"}, ExpectedOutput: "10.42.255.0"},
		{Name: "IncrementOverMaxTwice", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "10.42.255.255"}, ExpectedOutput: "10.43.0.0"},
		{Name: "IncrementOverMaxThrice", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "10.255.255.255"}, ExpectedOutput: "11.0.0.0"},
		{Name: "IncrementOverflow", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "255.255.255.255"}, ExpectedErr: "ip increment overflow"},
		{Name: "IncrementIPv6", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "2001:db8::1"}, ExpectedOutput: "2001:db8::2"},
		{Name: "IncrementIPv6ToMax", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "2001:db8::ffff"}, ExpectedOutput: "2001:db8::1:0"},
		{Name: "IncrementIPv6OverMax", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "2001:db8::1:ffff"}, ExpectedOutput: "2001:db8::2:0"},
		{Name: "IncrementIPv6OverMaxTwice", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "2001:db8::ffff:ffff"}, ExpectedOutput: "2001:db8::1:0:0"},
		{Name: "IncrementIPv6Overflow", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"}, ExpectedErr: "ip increment overflow"},
		{Name: "IncrementInvalidIP", Input: `{{ .V | parseIP | ipIncrement }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestIPDecrement(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "DecrementOverBasic", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "10.42.0.2"}, ExpectedOutput: "10.42.0.1"},
		{Name: "DecrementToMin", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "10.42.0.1"}, ExpectedOutput: "10.42.0.0"},
		{Name: "DecrementOverMin", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "10.42.1.0"}, ExpectedOutput: "10.42.0.255"},
		{Name: "DecrementOverMinTwice", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "10.42.0.0"}, ExpectedOutput: "10.41.255.255"},
		{Name: "DecrementOverMinThrice", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "10.0.0.0"}, ExpectedOutput: "9.255.255.255"},
		{Name: "DecrementUnderflow", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "0.0.0.0"}, ExpectedErr: "ip decrement underflow"},
		{Name: "DecrementIPv6", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "2001:db8::2"}, ExpectedOutput: "2001:db8::1"},
		{Name: "DecrementIPv6ToMin", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "2001:db8::1"}, ExpectedOutput: "2001:db8::"},
		{Name: "DecrementIPv6OverMin", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "2001:db8::1:0"}, ExpectedOutput: "2001:db8::ffff"},
		{Name: "DecrementIPv6OverMinTwice", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "2001:db8::1:0:0"}, ExpectedOutput: "2001:db8::ffff:ffff"},
		{Name: "DecrementIPv6Underflow", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "0:0:0:0:0:0:0:0"}, ExpectedErr: "ip decrement underflow"},
		{Name: "DecrementInvalidIP", Input: `{{ .V | parseIP | ipDecrement }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestCIDRContains(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "IPv4Contains", Input: `{{ .V | cidrContains "10.42.0.0/16" }}`, Data: map[string]any{"V": "10.42.0.1"}, ExpectedOutput: "true"},
		{Name: "IPv4NotContains", Input: `{{ .V | cidrContains "10.42.0.0/16" }}`, Data: map[string]any{"V": "10.51.0.1"}, ExpectedOutput: "false"},
		{Name: "IPv6Contains", Input: `{{ .V | cidrContains "2001:db8::/32" }}`, Data: map[string]any{"V": "2001:db8::1"}, ExpectedOutput: "true"},
		{Name: "IPv6NotContains", Input: `{{ .V | cidrContains "2001:db8::/32" }}`, Data: map[string]any{"V": "2001:db9::1"}, ExpectedOutput: "false"},
		{Name: "InvalidCIDR", Input: `{{ .V | cidrContains "invalid" }}`, Data: map[string]any{"V": "10.42.0.1"}, ExpectedErr: "invalid CIDR address"},
		{Name: "InvalidIP", Input: `{{ .V | cidrContains "10.42.0.0/16" }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid IP address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestCIDRSize(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "CIDRSizeIPv4", Input: `{{ .V | cidrSize }}`, Data: map[string]any{"V": "10.42.0.0/16"}, ExpectedOutput: "65536"},
		{Name: "CIDRSizeIPv6", Input: `{{ .V | cidrSize }}`, Data: map[string]any{"V": "2001:db8::/32"}, ExpectedOutput: "79228162514264337593543950336"},
		{Name: "InvalidCIDR", Input: `{{ .V | cidrSize }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid CIDR address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestCIDRRangeList(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "CIDRRangeListIPv4", Input: `{{ .V | cidrRangeList | len }}`, Data: map[string]any{"V": "10.42.1.1/16"}, ExpectedOutput: "65536"},
		{Name: "CIDRRangeListIPv4WithOne", Input: `{{ .V | cidrRangeList | len }}`, Data: map[string]any{"V": "10.42.1.1/32"}, ExpectedOutput: "1"},
		{Name: "CIDRRangeListIPv6", Input: `{{ .V | cidrRangeList | len }}`, Data: map[string]any{"V": "2001:db8::1/120"}, ExpectedOutput: "256"},
		{Name: "InvalidCIDR", Input: `{{ .V | cidrRangeList }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid CIDR address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestCIDRFirstIP(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "CIDRFirstIPv4", Input: `{{ .V | cidrFirst }}`, Data: map[string]any{"V": "10.42.0.0/24"}, ExpectedOutput: "10.42.0.0"},
		{Name: "CIDRFirstIPv6", Input: `{{ .V | cidrFirst }}`, Data: map[string]any{"V": "2001:db8::/120"}, ExpectedOutput: "2001:db8::"},
		{Name: "InvalidCIDR", Input: `{{ .V | cidrFirst }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid CIDR address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestCIDRLastIP(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "CIDRLastIPv4", Input: `{{ .V | cidrLast }}`, Data: map[string]any{"V": "10.42.0.0/24"}, ExpectedOutput: "10.42.0.255"},
		{Name: "CIDRLastIPv4WithOne", Input: `{{ .V | cidrLast }}`, Data: map[string]any{"V": "10.42.0.0/32"}, ExpectedOutput: "10.42.0.0"},
		{Name: "CIDRLastIPv4Max", Input: `{{ .V | cidrLast }}`, Data: map[string]any{"V": "10.42.0.0/0"}, ExpectedOutput: "255.255.255.255"},
		{Name: "CIDRLastIPv6", Input: `{{ .V | cidrLast }}`, Data: map[string]any{"V": "2001:db8::/120"}, ExpectedOutput: "2001:db8::ff"},
		{Name: "CIDRLastIPv6WithOne", Input: `{{ .V | cidrLast }}`, Data: map[string]any{"V": "2001:db8::/128"}, ExpectedOutput: "2001:db8::"},
		{Name: "CIDRLastIPv6Max", Input: `{{ .V | cidrLast }}`, Data: map[string]any{"V": "2001:db8::/0"}, ExpectedOutput: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"},
		{Name: "InvalidCIDR", Input: `{{ .V | cidrLast }}`, Data: map[string]any{"V": "invalid"}, ExpectedErr: "invalid CIDR address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}

func TestCIDROverlap(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "CIDROverlapIPv4Subset", Input: `{{ cidrOverlap "10.42.0.0/24" "10.42.0.0/16" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv4NoOverlapDifferentRange", Input: `{{ cidrOverlap "192.168.1.0/24" "10.0.0.0/8" }}`, ExpectedOutput: "false"},
		{Name: "CIDROverlapIPv4PartialOverlap", Input: `{{ cidrOverlap "10.0.1.0/24" "10.0.0.0/16" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv4AdjacentRanges", Input: `{{ cidrOverlap "192.168.1.0/24" "192.168.2.0/24" }}`, ExpectedOutput: "false"},
		{Name: "CIDROverlapIPv4SameRange", Input: `{{ cidrOverlap "192.168.1.0/24" "192.168.1.0/24" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv4SubsetSmall", Input: `{{ cidrOverlap "10.10.0.0/30" "10.10.0.0/16" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv4DifferentNonOverlapping", Input: `{{ cidrOverlap "172.16.0.0/24" "192.168.0.0/24" }}`, ExpectedOutput: "false"},
		{Name: "CIDROverlapIPv4OverlapInLargeRange", Input: `{{ cidrOverlap "172.16.10.0/24" "172.16.0.0/12" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv4SubsetInSmallerRange", Input: `{{ cidrOverlap "192.168.1.128/25" "192.168.1.0/24" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv4NoOverlapEdgeCases", Input: `{{ cidrOverlap "10.255.255.0/24" "11.0.0.0/8" }}`, ExpectedOutput: "false"},
		{Name: "CIDROverlapIPv6Subset", Input: `{{ cidrOverlap "2001:db8::/64" "2001:db8::/32" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv6NoOverlapDifferentRange", Input: `{{ cidrOverlap "2001:db8::/64" "3001:db8::/32" }}`, ExpectedOutput: "false"},
		{Name: "CIDROverlapIPv6PartialOverlap", Input: `{{ cidrOverlap "2001:db8:1234::/64" "2001:db8::/32" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv6AdjacentRanges", Input: `{{ cidrOverlap "2001:db8:abcd::/64" "2001:db8:abce::/64" }}`, ExpectedOutput: "false"},
		{Name: "CIDROverlapIPv6SameRange", Input: `{{ cidrOverlap "2001:db8::/48" "2001:db8::/48" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv6SmallSubset", Input: `{{ cidrOverlap "2001:db8:abcd::/126" "2001:db8:abcd::/64" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv6DifferentNonOverlapping", Input: `{{ cidrOverlap "fe80::/10" "2001:db8::/32" }}`, ExpectedOutput: "false"},
		{Name: "CIDROverlapIPv6OverlapInLargeRange", Input: `{{ cidrOverlap "2001:db8:ffff::/48" "2001:db8::/32" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv6SubsetInSmallerRange", Input: `{{ cidrOverlap "2001:db8:abcd::/64" "2001:db8::/32" }}`, ExpectedOutput: "true"},
		{Name: "CIDROverlapIPv6NoOverlapEdgeCases", Input: `{{ cidrOverlap "2001:db8:ffff::/64" "2001:db9::/32" }}`, ExpectedOutput: "false"},
		{Name: "InvalidCIDRA", Input: `{{ cidrOverlap "invalid" "10.42.0.0/32" }}`, ExpectedErr: "invalid CIDR address"},
		{Name: "InvalidCIDRB", Input: `{{ cidrOverlap "10.42.0.0/32" "invalid" }}`, ExpectedErr: "invalid CIDR address"},
	}

	pesticide.RunTestCases(t, network.NewRegistry(), tc)
}
