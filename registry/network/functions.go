package network

import (
	"errors"
	"fmt"
	"math/big"
	"net"
)

// ParseIP parses a string representation of an IP address and returns its [net.IP] form.
//
// It attempts to parse the string as either an IPv4 or IPv6 address.
// If the provided string is not a valid IP address, an error is returned.
//
// Parameters:
//
//	str string - the string representation of the IP address.
//
// Returns:
//
//	net.IP - the parsed IP address in its net.IP format.
//	error - an error if the string cannot be parsed as a valid IP address.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: parseIP].
//
// [Sprout Documentation: parseIP]: https://docs.atom.codes/sprout/registries/network#parseip
func (nr *NetworkRegistry) ParseIP(str string) (net.IP, error) {
	ip := net.ParseIP(str)
	if ip == nil {
		return nil, errors.New("invalid IP address")
	}

	return ip, nil
}

// ParseMAC parses a string representation of a MAC address and returns its [net.HardwareAddr] form.
//
// It attempts to parse the provided string as a MAC address. If the string is not a valid MAC address,
// an error is returned.
//
// Parameters:
//
//	str string - the string representation of the MAC address.
//
// Returns:
//
//	net.HardwareAddr - the parsed MAC address in its net.HardwareAddr format.
//	error - an error if the string cannot be parsed as a valid MAC address.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: parseMAC].
//
// [Sprout Documentation: parseMAC]: https://docs.atom.codes/sprout/registries/network#parsemac
func (nr *NetworkRegistry) ParseMAC(str string) (net.HardwareAddr, error) {
	mac, err := net.ParseMAC(str)
	if err != nil {
		return nil, fmt.Errorf("cannot parse MAC address: %w", err)
	}

	return mac, nil
}

// ParseCIDR parses a string representation of an IP address and prefix length (CIDR notation)
// and returns its [*net.IPNet] form.
//
// It attempts to parse the provided string as a CIDR (Classless Inter-Domain Routing) block.
// If the string is not valid CIDR notation, an error is returned.
//
// Parameters:
//
//	str string - the string representation of the CIDR block.
//
// Returns:
//
//	*net.IPNet - the parsed IP network in its *net.IPNet format.
//	error - an error if the string cannot be parsed as valid CIDR notation.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: parseCIDR].
//
// [Sprout Documentation: parseCIDR]: https://docs.atom.codes/sprout/registries/network#parsecidr
func (nr *NetworkRegistry) ParseCIDR(str string) (*net.IPNet, error) {
	_, cidr, err := net.ParseCIDR(str)
	if err != nil {
		return nil, fmt.Errorf("cannot parse CIDR: %w", err)
	}

	return cidr, nil
}

// IPVersion determines the IP version (IPv4 or IPv6) from a string
// representation of an IP address.
//
// Parameters:
//
//	ipStr string - the string representation of the IP address.
//
// Returns:
//
//	uint8 - the IP version, 4 for IPv4 or 16 for IPv6.
//	error - an error if the IP address is invalid or cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ipVersion].
//
// [Sprout Documentation: ipVersion]: https://docs.atom.codes/sprout/registries/network#ipversion
func (nr *NetworkRegistry) IPVersion(ipStr string) (uint8, error) {
	ip, err := nr.ParseIP(ipStr)
	if err != nil {
		return 0, err
	}

	return nr.determineIPVersion(ip), nil
}

// IPIsLoopback checks if the given IP address is a loopback address.
//
// It parses the provided string as an IP address and checks whether it is a
// loopback address (e.g., 127.0.0.1 for IPv4, ::1 for IPv6).
//
// Parameters:
//
//	ipStr string - the string representation of the IP address.
//
// Returns:
//
//	bool - true if the IP address is a loopback address.
//	error - an error if the IP address is invalid or cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ipisLoopback].
//
// [Sprout Documentation: ipisLoopback]: https://docs.atom.codes/sprout/registries/network#ipisloopback
func (nr *NetworkRegistry) IPIsLoopback(ipStr string) (bool, error) {
	ip, err := nr.ParseIP(ipStr)
	if err != nil {
		return false, err
	}

	return ip.IsLoopback(), nil
}

// IPIsGlobalUnicast checks if the given IP address is a global unicast address.
//
// It parses the provided string as an IP address and checks whether it is a
// global unicast address. Global unicast addresses are globally unique and routable
// (not multicast, loopback, or private).
//
// Parameters:
//
//	ipStr string - the string representation of the IP address.
//
// Returns:
//
//	bool - true if the IP address is a global unicast address.
//	error - an error if the IP address is invalid or cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ipIsGlobalUnicast].
//
// [Sprout Documentation: ipIsGlobalUnicast]: https://docs.atom.codes/sprout/registries/network#ipisglobalunicast
func (nr *NetworkRegistry) IPIsGlobalUnicast(ipStr string) (bool, error) {
	ip, err := nr.ParseIP(ipStr)
	if err != nil {
		return false, err
	}

	return ip.IsGlobalUnicast(), nil
}

// IPIsMulticast checks if the given IP address is a multicast address.
//
// It parses the provided string as an IP address and checks whether it is a
// multicast address. Multicast addresses are used to send data to multiple receivers.
//
// Parameters:
//
//	ipStr string - the string representation of the IP address.
//
// Returns:
//
//	bool - true if the IP address is a multicast address.
//	error - an error if the IP address is invalid or cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ipIsMulticast].
//
// [Sprout Documentation: ipIsMulticast]: https://docs.atom.codes/sprout/registries/network#ipismulticast
func (nr *NetworkRegistry) IPIsMulticast(ipStr string) (bool, error) {
	ip, err := nr.ParseIP(ipStr)
	if err != nil {
		return false, err
	}

	return ip.IsMulticast(), nil
}

// IPIsPrivate checks if the given IP address is a private address.
//
// It parses the provided string as an IP address and checks whether it is
// a private address. Private addresses are typically used for local
// communication within a network (e.g., 192.168.x.x).
//
// Parameters:
//
//	ipStr string - the string representation of the IP address.
//
// Returns:
//
//	bool - true if the IP address is a private address.
//	error - an error if the IP address is invalid or cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ipIsPrivate].
//
// [Sprout Documentation: ipIsPrivate]: https://docs.atom.codes/sprout/registries/network#ipisprivate
func (nr *NetworkRegistry) IPIsPrivate(ipStr string) (bool, error) {
	ip, err := nr.ParseIP(ipStr)
	if err != nil {
		return false, err
	}

	return ip.IsPrivate(), nil
}

// IPIncrement increments the given IP address by one unit.
// This function works for both IPv4 and IPv6 addresses.
//
// It converts the IP to the correct byte length depending on the version (IPv4 or IPv6)
// and increments the address by 1. In case of an overflow (e.g., incrementing
// 255.255.255.255 in IPv4), an error is returned.
//
// Parameters:
//
//	ip net.IP - the IP address to increment.
//
// Returns:
//
//	net.IP - the incremented IP address.
//	error - an error if the IP address overflows or if the IP version cannot be determined.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ipIncrement].
//
// [Sprout Documentation: ipIncrement]: https://docs.atom.codes/sprout/registries/network#ipincrement
func (nr *NetworkRegistry) IPIncrement(ip net.IP) (net.IP, error) {
	switch nr.determineIPVersion(ip) {
	case 4:
		ip = ip.To4()
	case 6:
		ip = ip.To16()
	}

	inc := 1 // increment value
	for i := len(ip) - 1; i >= 0 && inc > 0; i-- {
		if ip[i]+byte(inc) < ip[i] { // detect overflow
			ip[i] = 0
		} else {
			ip[i] += byte(inc)
			inc = 0 // stop further increments
		}
	}

	if inc > 0 {
		return net.IP{}, errors.New("ip increment overflow")
	}

	return ip, nil
}

// IPDecrement decrements the given IP address by one unit.
// This function works for both IPv4 and IPv6 addresses.
//
// It converts the IP to the correct byte length depending on the version (IPv4 or IPv6)
// and decrements the address by 1. In case of an underflow (e.g., decrementing
// 0.0.0.0 in IPv4), an error is returned.
//
// Parameters:
//
//	ip net.IP - the IP address to decrement.
//
// Returns:
//
//	net.IP - the decremented IP address.
//	error - an error if the IP address underflows or if the IP version cannot be determined.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ipDecrement].
//
// [Sprout Documentation: ipDecrement]: https://docs.atom.codes/sprout/registries/network#ipdecrement
func (nr *NetworkRegistry) IPDecrement(ip net.IP) (net.IP, error) {
	switch nr.determineIPVersion(ip) {
	case 4:
		ip = ip.To4()
	case 6:
		ip = ip.To16()
	}

	dec := 1 // decrement value
	for i := len(ip) - 1; i >= 0 && dec > 0; i-- {
		if ip[i] < byte(dec) { // detect underflow
			ip[i] = 0xFF
		} else {
			ip[i] -= byte(dec)
			dec = 0 // stop further decrements
		}
	}

	if dec > 0 {
		return net.IP{}, errors.New("ip decrement underflow")
	}

	return ip, nil
}

// CIDRContains checks if a given IP address is contained within a specified CIDR block.
//
// It parses both the CIDR block and the IP address, and checks whether the IP falls
// within the network range defined by the CIDR. If either the CIDR or the IP address is invalid,
// an error is returned.
//
// Parameters:
//
//	cidrStr string - the string representation of the CIDR block.
//	ip string - the string representation of the IP address to check.
//
// Returns:
//
//	bool - true if the IP address is within the CIDR block, false otherwise.
//	error - an error if the CIDR block or IP address cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: cidrContains].
//
// [Sprout Documentation: cidrContains]: https://docs.atom.codes/sprout/registries/network#cidrcontains
func (nr *NetworkRegistry) CIDRContains(cidrStr string, ip string) (bool, error) {
	parsedCIRDR, err := nr.ParseCIDR(cidrStr)
	if err != nil {
		return false, err
	}

	parsedIP, err := nr.ParseIP(ip)
	if err != nil {
		return false, err
	}

	return parsedCIRDR.Contains(parsedIP), nil
}

// CIDRSize calculates the total number of IP addresses in the given CIDR block.
// It works for both IPv4 and IPv6 CIDR blocks.
//
// The function parses the CIDR string, determines the IP version, and calculates the
// size of the network range based on the prefix length.
//
// Parameters:
//
//	cidr string - the string representation of the CIDR block.
//
// Returns:
//
//	*big.Int - the total number of IP addresses in the CIDR block.
//	error - an error if the CIDR block cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: cidrSize].
//
//	{{ cidrSize "2001:db8::/32" }}    // Output: 79228162514264337593543950336 (IPv6 range)
//
// [Sprout Documentation: cidrSize]: https://docs.atom.codes/sprout/registries/network#cidrsize
func (nr *NetworkRegistry) CIDRSize(cidrStr string) (*big.Int, error) {
	cidr, err := nr.ParseCIDR(cidrStr)
	if err != nil {
		return nil, err
	}

	ones, bits := cidr.Mask.Size()

	// Calculate the number of addresses: 2^(bits - ones)
	// Use big.Int for handling large values
	size := new(big.Int)
	size.Exp(big.NewInt(2), big.NewInt(int64(bits-ones)), nil)

	return size, nil
}

// CIDRRangeList generates a list of all IP addresses within the given CIDR block.
// It works for both IPv4 and IPv6 CIDR blocks,
// ! WARNING that generating all IPs in a large IPv4/IPv6 block may consume
// ! significant memory and processing time.
//
// Parameters:
//
//	cidr string - the string representation of the CIDR block.
//
// Returns:
//
//	[]net.IP - a slice containing all IP addresses within the CIDR block.
//	error - an error if the CIDR block cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: cidrRangeList].
//
// [Sprout Documentation: cidrRangeList]: https://docs.atom.codes/sprout/registries/network#cidrrangelist
func (nr *NetworkRegistry) CIDRRangeList(cidrStr string) ([]net.IP, error) {
	cidr, err := nr.ParseCIDR(cidrStr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR block: %w", err)
	}

	// Get the starting IP
	startIP := cidr.IP
	ones, bits := cidr.Mask.Size()
	totalIPs := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bits-ones)), nil)

	// Prepare a slice to store all IP addresses
	ipList := make([]net.IP, 0, totalIPs.Int64())

	// Use a big.Int to increment IP addresses, as they can
	// overflow standard integer types
	currentIP := new(big.Int).SetBytes(startIP)

	for i := big.NewInt(0); i.Cmp(totalIPs) < 0; i.Add(i, big.NewInt(1)) {
		// Convert the current big.Int to an IP address
		ipBytes := currentIP.Bytes()

		// Convert back to net.IP format (ensuring the length is correct for IPv4 or IPv6)
		ip := make(net.IP, len(startIP))
		copy(ip[len(ip)-len(ipBytes):], ipBytes)

		ipList = append(ipList, ip)

		// Increment the current IP
		currentIP.Add(currentIP, big.NewInt(1))
	}

	return ipList, nil
}

// CIDRFirst returns the first IP address in the given CIDR block.
//
// Parameters:
//
//	cidrStr string - the string representation of the CIDR block.
//
// Returns:
//
//	string - the first IP address as a string.
//	error - an error if the CIDR block cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: cidrFirst].
//
// [Sprout Documentation: cidrFirst]: https://docs.atom.codes/sprout/registries/network#cidrfirst
func (nr *NetworkRegistry) CIDRFirst(cidrStr string) (string, error) {
	cidr, err := nr.ParseCIDR(cidrStr)
	if err != nil {
		return "", err
	}

	return cidr.IP.String(), nil
}

// CIDRLast returns the last IP address in the given CIDR block.
//
// Parameters:
//
//	cidrStr string - the string representation of the CIDR block.
//
// Returns:
//
//	string - the last IP address as a string.
//	error - an error if the CIDR block cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: cidrLast].
//
// [Sprout Documentation: cidrLast]: https://docs.atom.codes/sprout/registries/network#cidrlast
func (nr *NetworkRegistry) CIDRLast(cidrStr string) (string, error) {
	cidr, err := nr.ParseCIDR(cidrStr)
	if err != nil {
		return "", err
	}

	return nr.calculateLastIP(cidr).String(), nil
}

// CIDROverlap checks if two CIDR blocks overlap.
// It parses both CIDR blocks and determines whether they overlap.
//
// Parameters:
//
//	cidrStrA string - the first CIDR block.
//	cidrStrB string - the second CIDR block.
//
// Returns:
//
//	bool - true if the two CIDR blocks overlap, false otherwise.
//	error - an error if either of the CIDR blocks cannot be parsed.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: cidrOverlap].
//
// [Sprout Documentation: cidrOverlap]: https://docs.atom.codes/sprout/registries/network#cidroverlap
func (nr *NetworkRegistry) CIDROverlap(cidrStrA, cidrStrB string) (bool, error) {
	cidrA, err := nr.ParseCIDR(cidrStrA)
	if err != nil {
		return false, err
	}

	cidrB, err := nr.ParseCIDR(cidrStrB)
	if err != nil {
		return false, err
	}

	return cidrB.Contains(cidrA.IP.To16()) ||
			cidrB.Contains(nr.calculateLastIP(cidrA)) ||
			cidrA.Contains(cidrB.IP.To16()) ||
			cidrA.Contains(nr.calculateLastIP(cidrB)),
		nil
}
