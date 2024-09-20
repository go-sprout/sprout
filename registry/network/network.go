package network

import (
	"github.com/go-sprout/sprout"
)

type NetworkRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *NetworkRegistry {
	return &NetworkRegistry{}
}

// Uid returns the unique identifier of the registry.
func (nr *NetworkRegistry) Uid() string {
	return "network" // ! Must be unique and in camel case
}

// LinkHandler links the handler to the registry at runtime.
func (nr *NetworkRegistry) LinkHandler(fh sprout.Handler) error {
	nr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (nr *NetworkRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "parseIP", nr.ParseIP)
	sprout.AddFunction(funcsMap, "parseMAC", nr.ParseMAC)
	sprout.AddFunction(funcsMap, "parseCIDR", nr.ParseCIDR)
	sprout.AddFunction(funcsMap, "ipVersion", nr.IPVersion)
	sprout.AddFunction(funcsMap, "ipIsLoopback", nr.IPIsLoopback)
	sprout.AddFunction(funcsMap, "ipIsGlobalUnicast", nr.IPIsGlobalUnicast)
	sprout.AddFunction(funcsMap, "ipIsMulticast", nr.IPIsMulticast)
	sprout.AddFunction(funcsMap, "ipIsPrivate", nr.IPIsPrivate)
	sprout.AddFunction(funcsMap, "ipIncrement", nr.IPIncrement)
	sprout.AddFunction(funcsMap, "ipDecrement", nr.IPDecrement)
	sprout.AddFunction(funcsMap, "cidrContains", nr.CIDRContains)
	sprout.AddFunction(funcsMap, "cidrSize", nr.CIDRSize)
	sprout.AddFunction(funcsMap, "cidrRangeList", nr.CIDRRangeList)
	sprout.AddFunction(funcsMap, "cidrFirst", nr.CIDRFirst)
	sprout.AddFunction(funcsMap, "cidrLast", nr.CIDRLast)
	sprout.AddFunction(funcsMap, "cidrOverlap", nr.CIDROverlap)
	return nil
}

func (nr *NetworkRegistry) RegisterAliases(aliasMap sprout.FunctionAliasMap) error {
	// Register your alias here if you have any or remove this method
	return nil
}

func (nr *NetworkRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	// Register your notices here if you have any or remove this method
	return nil
}
