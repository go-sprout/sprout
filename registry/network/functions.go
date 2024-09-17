package network

import (
	"errors"
	"fmt"
	"net"
)

func (nr *NetworkRegistry) ParseIP(str string) (net.IP, error) {
	ip := net.ParseIP(str)
	if ip == nil {
		return nil, errors.New("invalid IP address")
	}

	return ip, nil
}

func (nr *NetworkRegistry) ParseMAC(str string) (net.HardwareAddr, error) {
	mac, err := net.ParseMAC(str)
	if err != nil {
		return nil, fmt.Errorf("cannot parse MAC address: %w", err)
	}

	return mac, nil
}

func (nr *NetworkRegistry) ParseCIDR(str string) (*net.IPNet, error) {
	_, cidr, err := net.ParseCIDR(str)
	if err != nil {
		return nil, fmt.Errorf("cannot parse CIDR: %w", err)
	}

	return cidr, nil
}

func (nr *NetworkRegistry) CIDRContains(cidr string, ip string) (bool, error) {
	parsedCIRDR, err := nr.ParseCIDR(cidr)
	if err != nil {
		return false, err
	}

	parsedIP, err := nr.ParseIP(ip)
	if err != nil {
		return false, err
	}

	return parsedCIRDR.Contains(parsedIP), nil
}

func (nr *NetworkRegistry) SizeOfCIDR(cidr string) (string, error) {
	parsedCIRDR, err := nr.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}

	ones, _ := parsedCIRDR.Mask.Size()
	return fmt.Sprintf("%d", 1<<uint(32-ones)), nil
}

func (nr *NetworkRegistry) IPVersion(ipStr string) (int, error) {
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
