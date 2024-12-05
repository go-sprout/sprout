package reflect

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mitchellh/copystructure"
)

// TypeIs compares the type of 'value' to a target type string 'target'.
// It returns true if the type of 'value' matches the 'target'.
//
// Parameters:
//
//	target string - the string representation of the type to check against.
//	value any - the variable whose type is being checked.
//
// Returns:
//
//	bool - true if 'value' is of type 'target', false otherwise.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: typeIs].
//
// [Sprout Documentation: typeIs]: https://docs.atom.codes/sprout/registries/reflect#typeis
func (rr *ReflectRegistry) TypeIs(target string, value any) bool {
	return target == rr.TypeOf(value)
}

// TypeIsLike compares the type of 'value' to a target type string 'target',
// including a wildcard '*' prefix option. It returns true if 'value' matches
// 'target' or '*target'. Useful for checking if a variable is of a specific
// type or a pointer to that type.
//
// Parameters:
//
//	target string - the string representation of the type or its wildcard version.
//	value any - the variable whose type is being checked.
//
// Returns:
//
//	bool - true if the type of 'value' matches 'target' or '*'+target, false otherwise.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: typeIsLike].
//
// [Sprout Documentation: typeIsLike]: https://docs.atom.codes/sprout/registries/reflect#typeislike
func (rr *ReflectRegistry) TypeIsLike(target string, value any) bool {
	t := rr.TypeOf(value)
	return target == t || "*"+target == t
}

// TypeOf returns the type of 'value' as a string.
//
// Parameters:
//
//	value any - the variable whose type is being determined.
//
// Returns:
//
//	string - the string representation of 'value's type.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: typeOf].
//
// [Sprout Documentation: typeOf]: https://docs.atom.codes/sprout/registries/reflect#typeof
func (rr *ReflectRegistry) TypeOf(value any) string {
	return fmt.Sprintf("%T", value)
}

// KindIs compares the kind of 'value' to a target kind string 'target'.
// It returns true if the kind of 'value' matches the 'target'.
//
// Parameters:
//
//	target string - the string representation of the kind to check against.
//	value any - the variable whose kind is being checked.
//
// Returns:
//
//	bool - true if 'value's kind is 'target', false otherwise.
//	error - when 'value' is nil.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: kindIs].
//
// [Sprout Documentation: kindIs]: https://docs.atom.codes/sprout/registries/reflect#kindis
func (rr *ReflectRegistry) KindIs(target string, value any) (bool, error) {
	result, err := rr.KindOf(value)
	if err != nil {
		return false, err
	}

	return result == target, nil
}

// KindOf returns the kind of 'value' as a string.
//
// Parameters:
//
//	value any - the variable whose kind is being determined.
//
// Returns:
//
//	string - the string representation of 'value's kind.
//	error - when 'value' is nil.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: kindOf].
//
// [Sprout Documentation: kindOf]: https://docs.atom.codes/sprout/registries/reflect#kindof
func (rr *ReflectRegistry) KindOf(value any) (string, error) {
	if value == nil {
		return "", errors.New("value must not be nil")
	}

	return reflect.ValueOf(value).Kind().String(), nil
}

// HasField checks whether a struct has a field with a given name.
//
// Parameters:
//
//	name string - the name of the field that is being checked.
//	value any - the struct that is being checked.
//
// Returns:
//
//	bool - true if the struct 'value' contains a field with the name 'name', false otherwise.
//	error - when the last argument is not a struct.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: hasField].
//
// [Sprout Documentation: hasField]: https://docs.atom.codes/sprout/registries/reflect#hasfield
func (rr *ReflectRegistry) HasField(name string, value any) (bool, error) {
	rv := reflect.Indirect(reflect.ValueOf(value))
	if rv.Kind() != reflect.Struct {
		return false, errors.New("last argument must be a struct")
	}
	return rv.FieldByName(name).IsValid(), nil
}

// DeepEqual determines if two variables, 'x' and 'y', are deeply equal.
// It uses reflect.DeepEqual to evaluate equality.
//
// Parameters:
//
//	x, y any - the variables to be compared.
//
// Returns:
//
//	bool - true if 'x' and 'y' are deeply equal, false otherwise.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: deepEqual].
//
// [Sprout Documentation: deepEqual]: https://docs.atom.codes/sprout/registries/reflect#deepequal
func (rr *ReflectRegistry) DeepEqual(x, y any) bool {
	return reflect.DeepEqual(y, x)
}

// DeepCopy performs a deep copy of 'value' and panics if copying fails.
// It relies on MustDeepCopy to perform the copy and handle errors internally.
//
// Parameters:
//
//	value any - the element to be deeply copied.
//
// Returns:
//
//	any - a deep copy of 'value'.
//	error - when 'value' is nil.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: deepCopy].
//
// [Sprout Documentation: deepCopy]: https://docs.atom.codes/sprout/registries/reflect#deepcopy
func (rr *ReflectRegistry) DeepCopy(value any) (any, error) {
	if value == nil {
		return nil, errors.New("value cannot be nil")
	}
	return copystructure.Copy(value)
}
