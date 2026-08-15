package uniqueid

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// predefinedNamespaces maps the well-known namespace names defined by RFC 9562
// to their UUID, to let templates use `dns` instead of the raw UUID.
var predefinedNamespaces = map[string]uuid.UUID{
	"dns":  uuid.NameSpaceDNS,
	"url":  uuid.NameSpaceURL,
	"oid":  uuid.NameSpaceOID,
	"x500": uuid.NameSpaceX500,
}

// computeNamespace returns the UUID of the given namespace. The namespace can
// either be one of the predefined names (dns, url, oid, x500) or any valid UUID
// to use a custom namespace.
func computeNamespace(namespace string) (uuid.UUID, error) {
	if ns, ok := predefinedNamespaces[strings.ToLower(namespace)]; ok {
		return ns, nil
	}

	ns, err := uuid.Parse(namespace)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid namespace %q: must be one of dns, url, oid, x500 or a valid UUID", namespace)
	}
	return ns, nil
}
