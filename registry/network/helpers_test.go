package network

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDetermineIPVersion tests the determineIPVersion function.
func TestDetermineIPVersion(t *testing.T) {
	nr := &NetworkRegistry{}

	// Test case: IPv4 address
	ipv4 := net.ParseIP("192.168.0.1")
	assert.Equal(t, 4, nr.determineIPVersion(ipv4))

	// Test case: IPv6 address
	ipv6 := net.ParseIP("2001:db8::1")
	assert.Equal(t, 6, nr.determineIPVersion(ipv6))

	// Test case: IPv4-mapped IPv6 address
	ipv4MappedIPv6 := net.ParseIP("::ffff:192.168.0.1")
	assert.Equal(t, 4, nr.determineIPVersion(ipv4MappedIPv6))
}

// TestCalculateLastIP tests the calculateLastIP function.
func TestCalculateLastIP(t *testing.T) {
	nr := &NetworkRegistry{}

	// Test case: IPv4 CIDR
	_, cidrIPv4, err := net.ParseCIDR("192.168.0.0/24")
	require.NoError(t, err)
	expectedIPv4 := net.ParseIP("192.168.0.255").To4() // Use To4() to ensure it's in 4-byte format
	assert.Equal(t, expectedIPv4, nr.calculateLastIP(cidrIPv4).To4())

	// Test case: IPv6 CIDR
	_, cidrIPv6, err := net.ParseCIDR("2001:db8::/64")
	require.NoError(t, err)
	expectedIPv6 := net.ParseIP("2001:db8::ffff:ffff:ffff:ffff")
	assert.Equal(t, expectedIPv6, nr.calculateLastIP(cidrIPv6))

	// Test case: Smaller IPv4 CIDR block
	_, cidrIPv4Small, err := net.ParseCIDR("192.168.1.0/30")
	require.NoError(t, err)
	expectedIPv4Small := net.ParseIP("192.168.1.3").To4() // Use To4() to ensure it's in 4-byte format
	assert.Equal(t, expectedIPv4Small, nr.calculateLastIP(cidrIPv4Small).To4())
}
