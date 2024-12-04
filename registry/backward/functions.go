package backward

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/url"
)

// ! DEPRECATED: This should be removed in the next major version.
//
// Fail creates an error with a specified message and returns a nil pointer
// alongside the created error. This function is typically used to indicate
// failure conditions in functions that return a pointer and an error.
//
// Parameters:
//
//	message string - the error message to be associated with the returned error.
//
// Returns:
//
//	*uint - always returns nil, indicating no value is associated with the failure.
//	error - the error object containing the provided message.
//
// For an example of this function in a go template, refer to [Sprout Documentation: sha256Sum].
//
// [Sprout Documentation: sha256Sum]: https://docs.atom.codes/sprout/registries/backward#fail
func (bcr *BackwardCompatibilityRegistry) Fail(message string) (*uint, error) {
	return nil, errors.New(message)
}

// ! DEPRECATED: This should be removed in the next major version.
// UrlParse parses a given URL string and returns a map with its components.
//
// Parameters:
//
//	v string - the URL string to parse.
//
// Returns:
//
//	map[string]any - a map containing the URL components: "scheme", "host",
//									"hostname", "path", "query", "opaque", "fragment", and "userinfo".
//	error - an error object if the URL string is invalid.
//
// For an example of this function in a go template, refer to [Sprout Documentation: urlParse].
//
// [Sprout Documentation: urlParse]: https://docs.atom.codes/sprout/registries/backward#urlparse
func (bcr *BackwardCompatibilityRegistry) UrlParse(v string) (map[string]any, error) {
	dict := map[string]any{}
	parsedURL, err := url.Parse(v)
	if err != nil {
		return dict, fmt.Errorf("unable to parse url: %w", err)
	}
	dict["scheme"] = parsedURL.Scheme
	dict["host"] = parsedURL.Host
	dict["hostname"] = parsedURL.Hostname()
	dict["path"] = parsedURL.Path
	dict["query"] = parsedURL.RawQuery
	dict["opaque"] = parsedURL.Opaque
	dict["fragment"] = parsedURL.Fragment
	if parsedURL.User != nil {
		dict["userinfo"] = parsedURL.User.String()
	} else {
		dict["userinfo"] = ""
	}

	return dict, nil
}

// ! DEPRECATED: This should be removed in the next major version.
// UrlJoin constructs a URL string from a given map of URL components.
//
// Parameters:
//
//	d map[string]any - a map containing the URL components: "scheme", "host",
//											"path", "query", "opaque", "fragment", and "userinfo".
//
// Returns:
//
//	string - the constructed URL string.
//	error - an error object if the URL components are invalid.
//
// For an example of this function in a go template, refer to [Sprout Documentation: urlJoin].
//
// [Sprout Documentation: urlJoin]: https://docs.atom.codes/sprout/registries/backward#urljoin
func (bcr *BackwardCompatibilityRegistry) UrlJoin(d map[string]any) (string, error) {
	resURL := url.URL{
		Scheme:   bcr.get(d, "scheme").(string),
		Host:     bcr.get(d, "host").(string),
		Path:     bcr.get(d, "path").(string),
		RawQuery: bcr.get(d, "query").(string),
		Opaque:   bcr.get(d, "opaque").(string),
		Fragment: bcr.get(d, "fragment").(string),
	}
	userinfo := bcr.get(d, "userinfo").(string)
	var user *url.Userinfo
	if userinfo != "" {
		tempURL, err := url.Parse(fmt.Sprintf("proto://%s@host", userinfo))
		if err != nil {
			return "", fmt.Errorf("unable to parse userinfo in dict: %w", err)
		}
		user = tempURL.User
	}

	resURL.User = user
	return resURL.String(), nil
}

// ! DEPRECATED: This should be removed in the next major version.
// GetHostByName returns a random IP address associated with a given hostname.
//
// Parameters:
//
//	name string - the hostname to resolve.
//
// Returns:
//
//	string - a randomly selected IP address associated with the hostname.
//	error - an error object if the hostname cannot be resolved.
//
// Note: This function currently lacks error handling
//
// For an example of this function in a go template, refer to [Sprout Documentation: getHostByName].
//
// [Sprout Documentation: getHostByName]: https://docs.atom.codes/sprout/registries/checksum#gethostbyname
func (bcr *BackwardCompatibilityRegistry) GetHostByName(name string) (string, error) {
	addrs, err := net.LookupHost(name)
	if err != nil {
		return "", fmt.Errorf("unable to resolve hostname: %w", err)
	}
	return addrs[rand.Intn(len(addrs))], nil
}
