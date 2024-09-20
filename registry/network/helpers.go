package network

import (
	"net"
)

// determineIPVersion determines the IP version (IPv4 or IPv6) for a given net.IP.
//
// It checks the length of the IP address to identify the IP version:
// - Returns 4 if it's an IPv4 address.
// - Returns 6 if it's an IPv6 address.
//
// Parameters:
//
//	ip net.IP - the IP address to check.
//
// Returns:
//
//	int - 4 for IPv4, 6 for IPv6.
func (nr *NetworkRegistry) determineIPVersion(ip net.IP) int {
	// Check if it's an IPv4 address
	if len(ip) == net.IPv4len || (len(ip) == net.IPv6len && ip.To4() != nil) {
		return 4
	}

	return 6
}

// calculateLastIP calculates the last IP address in the range for the given CIDR block.
//
// It takes the network address and applies the subnet mask to find the last usable IP in the CIDR range.
//
// Parameters:
//
//	cidr *net.IPNet - the CIDR block.
//
// Returns:
//
//	net.IP - the last IP address in the CIDR block.
func (nr *NetworkRegistry) calculateLastIP(cidr *net.IPNet) net.IP {
	// Get the last IP of the CIDR
	lastIP := make(net.IP, len(cidr.IP))
	copy(lastIP, cidr.IP)

	for i := 0; i < len(cidr.Mask); i++ {
		lastIP[i] |= ^cidr.Mask[i]
	}

	return lastIP
}
