package main_test

import (
	"errors"
	"testing"
)

func BenchmarkDraft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPVersion("175.24.14.24")
		IPVersion("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
	}
}

func BenchmarkBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ipv("175.24.14.24")
		ipv("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
	}
}

func IPVersion(ipStr string) (int, error) {
	isIPv4, isIPv6 := false, false
	for i := 0; i < len(ipStr); i++ {
		c := ipStr[i]
		if c == '.' {
			isIPv4 = true
		} else if c == ':' {
			isIPv6 = true
		}
	}
	if isIPv4 && isIPv6 {
		return 0, errors.New("invalid IP address format")
	} else if isIPv4 {
		return 4, nil
	} else if isIPv6 {
		return 6, nil
	}

	return 0, errors.New("unknown IP address format")
}

func ipv(ipStr string) (int, error) {
	var flags uint8 = 0
	for i := 0; i < len(ipStr); i++ {
		c := ipStr[i]
		flags |= ((c ^ '.') * 1) | ((c ^ ':') * 2)
	}

	if flags == 1 || flags == 2 {
		return 4 + 2*(int(flags)>>1), nil
	}

	return 0, errors.New("invalid IP address format")
}
