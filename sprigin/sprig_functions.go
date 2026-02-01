package sprigin

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-sprout/sprout/registry/encoding"
	"github.com/go-sprout/sprout/registry/maps"
	rreflect "github.com/go-sprout/sprout/registry/reflect"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/time"
)

// isList checks if a value is a slice or array type.
// Used for signature detection in compatibility functions.
func isList(v any) bool {
	if v == nil {
		return false
	}
	rt := reflect.TypeOf(v)
	return rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array
}

// isMap checks if a value is a map type.
// Used for signature detection in compatibility functions.
func isMap(v any) bool {
	if v == nil {
		return false
	}
	_, ok := v.(map[string]any)
	return ok
}

// This file contains Sprig-compatible function wrappers that provide backward
// compatibility with Sprig's function signatures. These functions have different
// argument orders compared to Sprout's native functions.
//
// Sprig uses: function(target, args...)
// Sprout uses: function(args..., target)
//
// These wrappers are only used through the sprigin compatibility layer.
// Deprecation warnings are only logged when old Sprig signature is detected.

// sprigGet handles both Sprig and Sprout signatures for get.
// Sprig: get(dict, key) - dict first
// Sprout: get(key, dict) - dict last (for piping)
// Detection is based on which argument is a map type.
func (sh *SprigHandler) sprigGet(mr *maps.MapsRegistry) func(args ...any) (any, error) {
	return func(args ...any) (any, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("get requires exactly 2 arguments, got %d", len(args))
		}

		firstIsMap := isMap(args[0])
		secondIsMap := isMap(args[1])

		switch {
		case firstIsMap && !secondIsMap:
			// Old Sprig signature: get(dict, key)
			sh.SignatureWarn("get", "{{ get $dict \"key\" }}", "{{ $dict | get \"key\" }}")
			dict := args[0].(map[string]any)
			key, ok := args[1].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			return mr.Get(key, dict)
		case !firstIsMap && secondIsMap:
			// New Sprout signature: get(key, dict)
			key, ok := args[0].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			dict := args[1].(map[string]any)
			return mr.Get(key, dict)
		default:
			// Ambiguous or invalid - try Sprig behavior (with warning)
			sh.SignatureWarn("get", "{{ get $dict \"key\" }}", "{{ $dict | get \"key\" }}")
			dict, ok := args[0].(map[string]any)
			if !ok {
				return nil, fmt.Errorf("first argument must be a map[string]any")
			}
			key, ok := args[1].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			return mr.Get(key, dict)
		}
	}
}

// sprigSet handles both Sprig and Sprout signatures for set.
// Sprig: set(dict, key, value) - dict first
// Sprout: set(key, value, dict) - dict last (for piping)
// Detection is based on which argument is a map type.
func (sh *SprigHandler) sprigSet(mr *maps.MapsRegistry) func(args ...any) (map[string]any, error) {
	return func(args ...any) (map[string]any, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("set requires exactly 3 arguments, got %d", len(args))
		}

		firstIsMap := isMap(args[0])
		lastIsMap := isMap(args[2])

		switch {
		case firstIsMap && !lastIsMap:
			// Old Sprig signature: set(dict, key, value)
			sh.SignatureWarn("set", "{{ set $dict \"key\" \"value\" }}", "{{ $dict | set \"key\" \"value\" }}")
			dict := args[0].(map[string]any)
			key, ok := args[1].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			value := args[2]
			return mr.Set(key, value, dict)
		case !firstIsMap && lastIsMap:
			// New Sprout signature: set(key, value, dict)
			key, ok := args[0].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			value := args[1]
			dict := args[2].(map[string]any)
			return mr.Set(key, value, dict)
		default:
			// Ambiguous - try Sprig behavior (with warning)
			sh.SignatureWarn("set", "{{ set $dict \"key\" \"value\" }}", "{{ $dict | set \"key\" \"value\" }}")
			dict, ok := args[0].(map[string]any)
			if !ok {
				return nil, fmt.Errorf("first argument must be a map[string]any")
			}
			key, ok := args[1].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			value := args[2]
			return mr.Set(key, value, dict)
		}
	}
}

// sprigUnset handles both Sprig and Sprout signatures for unset.
// Sprig: unset(dict, key) - dict first
// Sprout: unset(key, dict) - dict last (for piping)
// Detection is based on which argument is a map type.
func (sh *SprigHandler) sprigUnset(mr *maps.MapsRegistry) func(args ...any) (map[string]any, error) {
	return func(args ...any) (map[string]any, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("unset requires exactly 2 arguments, got %d", len(args))
		}

		firstIsMap := isMap(args[0])
		secondIsMap := isMap(args[1])

		switch {
		case firstIsMap && !secondIsMap:
			// Old Sprig signature: unset(dict, key)
			sh.SignatureWarn("unset", "{{ unset $dict \"key\" }}", "{{ $dict | unset \"key\" }}")
			dict := args[0].(map[string]any)
			key, ok := args[1].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			return mr.Unset(key, dict)
		case !firstIsMap && secondIsMap:
			// New Sprout signature: unset(key, dict)
			key, ok := args[0].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			dict := args[1].(map[string]any)
			return mr.Unset(key, dict)
		default:
			// Ambiguous - try Sprig behavior (with warning)
			sh.SignatureWarn("unset", "{{ unset $dict \"key\" }}", "{{ $dict | unset \"key\" }}")
			dict, ok := args[0].(map[string]any)
			if !ok {
				return nil, fmt.Errorf("first argument must be a map[string]any")
			}
			key, ok := args[1].(string)
			if !ok {
				return nil, fmt.Errorf("key must be a string")
			}
			return mr.Unset(key, dict)
		}
	}
}

// sprigHasKey handles both Sprig and Sprout signatures for hasKey.
// Sprig: hasKey(dict, key) - dict first
// Sprout: hasKey(key, dict) - dict last (for piping)
// Detection is based on which argument is a map type.
func (sh *SprigHandler) sprigHasKey(mr *maps.MapsRegistry) func(args ...any) (bool, error) {
	return func(args ...any) (bool, error) {
		if len(args) != 2 {
			return false, fmt.Errorf("hasKey requires exactly 2 arguments, got %d", len(args))
		}

		firstIsMap := isMap(args[0])
		secondIsMap := isMap(args[1])

		switch {
		case firstIsMap && !secondIsMap:
			// Old Sprig signature: hasKey(dict, key)
			sh.SignatureWarn("hasKey", "{{ hasKey $dict \"key\" }}", "{{ $dict | hasKey \"key\" }}")
			dict := args[0].(map[string]any)
			key, ok := args[1].(string)
			if !ok {
				return false, fmt.Errorf("key must be a string")
			}
			return mr.HasKey(key, dict)
		case !firstIsMap && secondIsMap:
			// New Sprout signature: hasKey(key, dict)
			key, ok := args[0].(string)
			if !ok {
				return false, fmt.Errorf("key must be a string")
			}
			dict := args[1].(map[string]any)
			return mr.HasKey(key, dict)
		default:
			// Ambiguous - try Sprig behavior (with warning)
			sh.SignatureWarn("hasKey", "{{ hasKey $dict \"key\" }}", "{{ $dict | hasKey \"key\" }}")
			dict, ok := args[0].(map[string]any)
			if !ok {
				return false, fmt.Errorf("first argument must be a map[string]any")
			}
			key, ok := args[1].(string)
			if !ok {
				return false, fmt.Errorf("key must be a string")
			}
			return mr.HasKey(key, dict)
		}
	}
}

// sprigPick handles both Sprig and Sprout signatures for pick.
// Sprig: pick(dict, keys...) - dict first
// Sprout: pick(keys..., dict) - dict last (for piping)
// Detection is based on which argument is a map type.
func (sh *SprigHandler) sprigPick(mr *maps.MapsRegistry) func(args ...any) (map[string]any, error) {
	return func(args ...any) (map[string]any, error) {
		if len(args) < 2 {
			return nil, fmt.Errorf("pick requires at least two arguments")
		}

		firstIsMap := isMap(args[0])
		lastIsMap := isMap(args[len(args)-1])

		switch {
		case firstIsMap && !lastIsMap:
			// Old Sprig signature: pick(dict, keys...)
			// Convert to Sprout: pick(keys..., dict)
			sh.SignatureWarn("pick", "{{ pick $dict \"key1\" \"key2\" }}", "{{ $dict | pick \"key1\" \"key2\" }}")
			dict := args[0]
			keys := args[1:]
			return mr.Pick(append(keys, dict)...)
		case !firstIsMap && lastIsMap:
			// New Sprout signature: pick(keys..., dict)
			// Already in correct order
			return mr.Pick(args...)
		default:
			// Ambiguous - default to Sprig behavior (with warning)
			sh.SignatureWarn("pick", "{{ pick $dict \"key1\" \"key2\" }}", "{{ $dict | pick \"key1\" \"key2\" }}")
			dict := args[0]
			keys := args[1:]
			return mr.Pick(append(keys, dict)...)
		}
	}
}

// sprigOmit handles both Sprig and Sprout signatures for omit.
// Sprig: omit(dict, keys...) - dict first
// Sprout: omit(keys..., dict) - dict last (for piping)
// Detection is based on which argument is a map type.
func (sh *SprigHandler) sprigOmit(mr *maps.MapsRegistry) func(args ...any) (map[string]any, error) {
	return func(args ...any) (map[string]any, error) {
		if len(args) < 2 {
			if isMap(args[0]) {
				return args[0].(map[string]any), nil
			}
			return nil, fmt.Errorf("omit requires at least two arguments")
		}

		firstIsMap := isMap(args[0])
		lastIsMap := isMap(args[len(args)-1])

		switch {
		case firstIsMap && !lastIsMap:
			// Old Sprig signature: omit(dict, keys...)
			// Convert to Sprout: omit(keys..., dict)
			sh.SignatureWarn("omit", "{{ omit $dict \"key1\" \"key2\" }}", "{{ $dict | omit \"key1\" \"key2\" }}")
			dict := args[0]
			keys := args[1:]
			return mr.Omit(append(keys, dict)...)
		case !firstIsMap && lastIsMap:
			// New Sprout signature: omit(keys..., dict)
			// Already in correct order
			return mr.Omit(args...)
		default:
			// Ambiguous - default to Sprig behavior (with warning)
			sh.SignatureWarn("omit", "{{ omit $dict \"key1\" \"key2\" }}", "{{ $dict | omit \"key1\" \"key2\" }}")
			dict := args[0]
			keys := args[1:]
			return mr.Omit(append(keys, dict)...)
		}
	}
}

// sprigDig implements Sprig's dig signature: dig(keys..., default, dict)
// Note: Sprig's dig doesn't split keys on dots, unlike Sprout's native dig.
// This always uses the old Sprig signature, so we always warn.
func (sh *SprigHandler) sprigDig() func(args ...any) (any, error) {
	return func(args ...any) (any, error) {
		if len(args) < 3 {
			return nil, errors.New("dig requires at least three arguments: a sequence of keys, a default value, and a dictionary")
		}

		dict, ok := args[len(args)-1].(map[string]any)
		if !ok {
			return nil, errors.New("last argument must be a map[string]any")
		}

		// Log deprecation warning for old Sprig signature
		sh.SignatureWarn("dig", "{{ dig \"key\" \"default\" $dict }}", "{{ $dict | dig \"key\" | default \"default\" }}")

		// The second-to-last argument is the default value (can be any type in Sprig)
		defaultValue := args[len(args)-2]

		keys, err := parseKeys(args[:len(args)-2])
		if err != nil {
			return nil, fmt.Errorf("cannot parse keys: %w", err)
		}

		value, err := digIntoDict(dict, keys)
		if err != nil || value == nil {
			return defaultValue, nil
		}

		return value, nil
	}
}

// sprigAppend handles both Sprig and Sprout signatures for append.
// Sprig: append(list, value) - list first
// Sprout: append(value, list) - value first (for piping)
// Detection is based on which argument is a list type.
// If ambiguous (both are lists), defaults to Sprig behavior for backward compatibility.
func (sh *SprigHandler) sprigAppend(sr *slices.SlicesRegistry) func(args ...any) ([]any, error) {
	return func(args ...any) ([]any, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("append requires exactly 2 arguments, got %d", len(args))
		}

		first, second := args[0], args[1]
		firstIsList := isList(first)
		secondIsList := isList(second)

		switch {
		case firstIsList && !secondIsList:
			// Old Sprig signature: append(list, value)
			sh.SignatureWarn("append", "{{ append $list \"value\" }}", "{{ $list | append \"value\" }}")
			return sr.Append(second, first)
		case !firstIsList && secondIsList:
			// New Sprout signature: append(value, list)
			return sr.Append(first, second)
		default:
			// Ambiguous or neither is list - default to Sprig behavior (with warning)
			sh.SignatureWarn("append", "{{ append $list \"value\" }}", "{{ $list | append \"value\" }}")
			return sr.Append(second, first)
		}
	}
}

// sprigPrepend handles both Sprig and Sprout signatures for prepend.
// Sprig: prepend(list, value) - list first
// Sprout: prepend(value, list) - value first (for piping)
// Detection is based on which argument is a list type.
// If ambiguous (both are lists), defaults to Sprig behavior for backward compatibility.
func (sh *SprigHandler) sprigPrepend(sr *slices.SlicesRegistry) func(args ...any) ([]any, error) {
	return func(args ...any) ([]any, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("prepend requires exactly 2 arguments, got %d", len(args))
		}

		first, second := args[0], args[1]
		firstIsList := isList(first)
		secondIsList := isList(second)

		switch {
		case firstIsList && !secondIsList:
			// Old Sprig signature: prepend(list, value)
			sh.SignatureWarn("prepend", "{{ prepend $list \"value\" }}", "{{ $list | prepend \"value\" }}")
			return sr.Prepend(second, first)
		case !firstIsList && secondIsList:
			// New Sprout signature: prepend(value, list)
			return sr.Prepend(first, second)
		default:
			// Ambiguous or neither is list - default to Sprig behavior (with warning)
			sh.SignatureWarn("prepend", "{{ prepend $list \"value\" }}", "{{ $list | prepend \"value\" }}")
			return sr.Prepend(second, first)
		}
	}
}

// sprigSlice handles both Sprig and Sprout signatures for slice.
// Sprig: slice(list, indices...) - list first
// Sprout: slice(indices..., list) - list last (for piping)
// Detection is based on which argument is a list type.
func (sh *SprigHandler) sprigSlice(sr *slices.SlicesRegistry) func(args ...any) (any, error) {
	return func(args ...any) (any, error) {
		if len(args) < 1 {
			return nil, fmt.Errorf("slice requires at least one argument")
		}

		if len(args) == 1 {
			return args[0], nil
		}

		firstIsList := isList(args[0])
		lastIsList := isList(args[len(args)-1])

		switch {
		case firstIsList && !lastIsList:
			// Old Sprig signature: slice(list, indices...)
			// Convert to Sprout: slice(indices..., list)
			sh.SignatureWarn("slice", "{{ slice $list 1 3 }}", "{{ $list | slice 1 3 }}")
			list := args[0]
			indices := args[1:]
			return sr.Slice(append(indices, list)...)
		case !firstIsList && lastIsList:
			// New Sprout signature: slice(indices..., list)
			// Already in correct order
			return sr.Slice(args...)
		default:
			// Ambiguous - default to Sprig behavior (with warning)
			sh.SignatureWarn("slice", "{{ slice $list 1 3 }}", "{{ $list | slice 1 3 }}")
			list := args[0]
			indices := args[1:]
			return sr.Slice(append(indices, list)...)
		}
	}
}

// sprigWithout handles both Sprig and Sprout signatures for without.
// Sprig: without(list, omit...) - list first
// Sprout: without(omit..., list) - list last (for piping)
// Detection is based on which argument is a list type.
func (sh *SprigHandler) sprigWithout(sr *slices.SlicesRegistry) func(args ...any) ([]any, error) {
	return func(args ...any) ([]any, error) {
		if len(args) < 2 {
			return nil, fmt.Errorf("without requires at least two arguments")
		}

		firstIsList := isList(args[0])
		lastIsList := isList(args[len(args)-1])

		switch {
		case firstIsList && !lastIsList:
			// Old Sprig signature: without(list, omit...)
			// Convert to Sprout: without(omit..., list)
			sh.SignatureWarn("without", "{{ without $list \"a\" \"b\" }}", "{{ $list | without \"a\" \"b\" }}")
			list := args[0]
			omit := args[1:]
			return sr.Without(append(omit, list)...)
		case !firstIsList && lastIsList:
			// New Sprout signature: without(omit..., list)
			// Already in correct order
			return sr.Without(args...)
		default:
			// Ambiguous - default to Sprig behavior (with warning)
			sh.SignatureWarn("without", "{{ without $list \"a\" \"b\" }}", "{{ $list | without \"a\" \"b\" }}")
			list := args[0]
			omit := args[1:]
			return sr.Without(append(omit, list)...)
		}
	}
}

// Helper functions for dig (copied from maps registry to avoid circular dependency)

func parseKeys(keySet []any) ([]string, error) {
	keys := make([]string, 0, len(keySet))
	for i, element := range keySet {
		key, ok := element.(string)
		if !ok {
			return nil, fmt.Errorf("all keys must be strings, got %T at position %d", element, i)
		}
		keys = append(keys, key)
	}
	return keys, nil
}

func digIntoDict(dict map[string]any, keys []string) (any, error) {
	current := any(dict)

	for i, key := range keys {
		switch typedCurrent := current.(type) {
		case map[string]any:
			value, exists := typedCurrent[key]
			if !exists {
				return nil, nil
			}

			if i == len(keys)-1 {
				return value, nil
			}

			current = value
		default:
			return nil, fmt.Errorf("value at key %q is not a nested dictionary but %T", keys[i-1], current)
		}
	}

	return current, nil
}

// sprigBase64Decode returns the error message as a string like sprig does,
// instead of returning an empty string when decoding fails.
// Sprig returns just the underlying error message without the "base64 decode error:" prefix.
func (sh *SprigHandler) sprigBase64Decode(er *encoding.EncodingRegistry) func(value string) string {
	return func(value string) string {
		result, err := er.Base64Decode(value)
		if err != nil {
			// Unwrap to get the underlying error message like sprig does
			if unwrapped := errors.Unwrap(err); unwrapped != nil {
				return unwrapped.Error()
			}
			return err.Error()
		}
		return result
	}
}

// sprigBase32Decode returns the error message as a string like sprig does,
// instead of returning an empty string when decoding fails.
// Sprig returns just the underlying error message without the "base32 decode error:" prefix.
func (sh *SprigHandler) sprigBase32Decode(er *encoding.EncodingRegistry) func(value string) string {
	return func(value string) string {
		result, err := er.Base32Decode(value)
		if err != nil {
			// Unwrap to get the underlying error message like sprig does
			if unwrapped := errors.Unwrap(err); unwrapped != nil {
				return unwrapped.Error()
			}
			return err.Error()
		}
		return result
	}
}

// sprigDate uses local timezone like sprig does, instead of the timezone
// from the time value itself.
func (sh *SprigHandler) sprigDate(tr *time.TimeRegistry) func(layout string, date any) (string, error) {
	return func(layout string, date any) (string, error) {
		return tr.DateInZone(layout, date, "Local")
	}
}

func (sh *SprigHandler) sprigSubstr() func(start, end int, value string) string {
	return func(start, end int, value string) string {
		if start < 0 {
			return value[:end]
		}
		if end > len(value) {
			return value[start:]
		}

		if start == 0 && end == 0 {
			sh.BreakingWarn("substr", "`substr 0 0` will return the full string in future versions instead of an empty string. Replace with `\"\"` if you need an empty result.")
			return ""
		}

		if end <= 0 {
			if end < 0 {
				sh.BreakingWarn("substr", "Negative `end` values will change from being ignored to counting from the string's end. Use `0` as end value to extract until the end of the string.")
			}
			return value[start:]
		}

		return value[start:end]
	}
}

func (sh *SprigHandler) sprigKindOf(rr *rreflect.ReflectRegistry) func(value any) (string, error) {
	return func(value any) (string, error) {
		if value == nil {
			sh.BreakingWarn("kindOf", "`kindOf nil` will return an error in future versions instead of \"invalid\". Check for nil before calling kindOf.")
			return "invalid", nil
		}

		return rr.KindOf(value)
	}
}

// sprigTitle uses the deprecated strings.Title function to match Sprig's behavior.
// Sprig's title function has two bugs compared to proper Unicode title casing:
// 1. It doesn't lowercase letters before capitalizing (so "HELLO" stays "HELLO")
// 2. It treats apostrophes as word separators (so "it's" becomes "It'S")
//
//nolint:staticcheck // SA1019: strings.Title is deprecated but needed for Sprig compatibility
func (sh *SprigHandler) sprigTitle() func(value string) string {
	return func(value string) string {
		sh.BreakingWarn("title", "Sprig's `title` uses deprecated strings.Title which doesn't properly handle Unicode or apostrophes. In future versions, `title` will use proper Unicode title casing. Use `toTitleCase` for the new behavior.")
		return strings.Title(value)
	}
}

// sprigAbbrevboth replicates Sprig's abbrevboth behavior which has a bug:
// when offset <= 4, it ignores the offset and just truncates from the right.
// This is for backward compatibility with Sprig's buggy behavior.
func (sh *SprigHandler) sprigAbbrevboth() func(offset int, maxWidth int, value string) string {
	return func(offset int, maxWidth int, value string) string {
		if value == "" {
			return ""
		}

		// Sprig's check: maxWidth < 4 or (offset > 0 and maxWidth < 7)
		if maxWidth < 4 || (offset > 0 && maxWidth < 7) {
			return value
		}

		strLen := len(value)
		if strLen <= maxWidth {
			return value
		}

		// This is the Sprig bug: when offset <= 4, it ignores the offset
		// and just truncates from the right (no left ellipsis)
		if offset <= 4 {
			if offset > 0 {
				sh.BreakingWarn("abbrevboth", "When offset <= 4, Sprig ignores the offset. In future versions, the offset will be respected. Use `ellipsis` if you want right-only truncation.")
			}
			// Right-only truncation: str[0:maxWidth-3] + "..."
			return value[0:maxWidth-3] + "..."
		}

		// Adjust offset if string is too short from offset position
		if offset > strLen {
			offset = strLen
		}
		if strLen-offset < (maxWidth - 3) {
			offset = strLen - (maxWidth - 3)
		}

		// Now handle the case where we need both ellipses
		if (offset + maxWidth - 3) < strLen {
			// "..." + abbreviate(str[offset:], maxWidth-3)
			remaining := value[offset:]
			if len(remaining) > maxWidth-3 {
				return "..." + remaining[0:maxWidth-6] + "..."
			}
			return "..." + remaining
		}

		// "..." + str[len-maxWidth+3:]
		return "..." + value[strLen-(maxWidth-3):]
	}
}
