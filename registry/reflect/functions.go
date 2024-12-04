package reflect

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mitchellh/copystructure"
)

// TypeIs compares the type of 'src' to a target type string 'target'.
// It returns true if the type of 'src' matches the 'target'.
//
// Parameters:
//
//	target string - the string representation of the type to check against.
//	src any - the variable whose type is being checked.
//
// Returns:
//
//	bool - true if 'src' is of type 'target', false otherwise.
//
// For an example of this function in a go template, refer to [Sprout Documentation: typeIs].
//
// [Sprout Documentation: typeIs]: https://docs.atom.codes/sprout/registries/reflect#typeis
func (rr *ReflectRegistry) TypeIs(target string, src any) bool {
	return target == rr.TypeOf(src)
}

// TypeIsLike compares the type of 'src' to a target type string 'target',
// including a wildcard '*' prefix option. It returns true if 'src' matches
// 'target' or '*target'. Useful for checking if a variable is of a specific
// type or a pointer to that type.
//
// Parameters:
//
//	target string - the string representation of the type or its wildcard version.
//	src any - the variable whose type is being checked.
//
// Returns:
//
//	bool - true if the type of 'src' matches 'target' or '*'+target, false otherwise.
//
// For an example of this function in a go template, refer to [Sprout Documentation: typeIsLike].
//
// [Sprout Documentation: typeIsLike]: https://docs.atom.codes/sprout/registries/reflect#typeislike
func (rr *ReflectRegistry) TypeIsLike(target string, src any) bool {
	t := rr.TypeOf(src)
	return target == t || "*"+target == t
}

// TypeOf returns the type of 'src' as a string.
//
// Parameters:
//
//	src any - the variable whose type is being determined.
//
// Returns:
//
//	string - the string representation of 'src's type.
//
// For an example of this function in a go template, refer to [Sprout Documentation: typeOf].
//
// [Sprout Documentation: typeOf]: https://docs.atom.codes/sprout/registries/reflect#typeof
func (rr *ReflectRegistry) TypeOf(src any) string {
	return fmt.Sprintf("%T", src)
}

// KindIs compares the kind of 'src' to a target kind string 'target'.
// It returns true if the kind of 'src' matches the 'target'.
//
// Parameters:
//
//	target string - the string representation of the kind to check against.
//	src any - the variable whose kind is being checked.
//
// Returns:
//
//	bool - true if 'src's kind is 'target', false otherwise.
//	error - when 'src' is nil.
//
// For an example of this function in a go template, refer to [Sprout Documentation: kindIs].
//
// [Sprout Documentation: kindIs]: https://docs.atom.codes/sprout/registries/reflect#kindis
func (rr *ReflectRegistry) KindIs(target string, src any) (bool, error) {
	result, err := rr.KindOf(src)
	if err != nil {
		return false, err
	}

	return result == target, nil
}

// KindOf returns the kind of 'src' as a string.
//
// Parameters:
//
//	src any - the variable whose kind is being determined.
//
// Returns:
//
//	string - the string representation of 'src's kind.
//	error - when 'src' is nil.
//
// For an example of this function in a go template, refer to [Sprout Documentation: kindOf].
//
// [Sprout Documentation: kindOf]: https://docs.atom.codes/sprout/registries/reflect#kindof
func (rr *ReflectRegistry) KindOf(src any) (string, error) {
	if src == nil {
		return "", errors.New("src must not be nil")
	}

	return reflect.ValueOf(src).Kind().String(), nil
}

// HasField checks whether a struct has a field with a given name.
//
// Parameters:
//
//	name string - the name of the field that is being checked.
//	src any - the struct that is being checked.
//
// Returns:
//
//	bool - true if the struct 'src' contains a field with the name 'name', false otherwise.
//	error - when the last argument is not a struct.
//
// For an example of this function in a go template, refer to [Sprout Documentation: hasField].
//
// [Sprout Documentation: hasField]: https://docs.atom.codes/sprout/registries/reflect#hasfield
func (rr *ReflectRegistry) HasField(name string, src any) (bool, error) {
	rv := reflect.Indirect(reflect.ValueOf(src))
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
// For an example of this function in a go template, refer to [Sprout Documentation: deepEqual].
//
// [Sprout Documentation: deepEqual]: https://docs.atom.codes/sprout/registries/reflect#deepequal
func (rr *ReflectRegistry) DeepEqual(x, y any) bool {
	return reflect.DeepEqual(y, x)
}

// DeepCopy performs a deep copy of 'element' and panics if copying fails.
// It relies on MustDeepCopy to perform the copy and handle errors internally.
//
// Parameters:
//
//	element any - the element to be deeply copied.
//
// Returns:
//
//	any - a deep copy of 'element'.
//	error - when 'element' is nil.
//
// For an example of this function in a go template, refer to [Sprout Documentation: deepCopy].
//
// [Sprout Documentation: deepCopy]: https://docs.atom.codes/sprout/registries/reflect#deepcopy
func (rr *ReflectRegistry) DeepCopy(element any) (any, error) {
	if element == nil {
		return nil, errors.New("element cannot be nil")
	}
	return copystructure.Copy(element)
}
