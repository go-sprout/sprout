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
// Example:
//
//	{{ "int", 42 | typeIs }} // Output: true
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
// Example:
//
//	{{ "*int", 42 | typeIsLike }} // Output: true
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
// Example:
//
//	{{ 42 | typeOf }} // Output: "int"
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
//
// Example:
//
//	{{ "int", 42 | kindIs }} // Output: true
func (rr *ReflectRegistry) KindIs(target string, src any) bool {
	return target == rr.KindOf(src)
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
//
// Example:
//
//	{{ 42 | kindOf }} // Output: "int"
func (rr *ReflectRegistry) KindOf(src any) string {
	return reflect.ValueOf(src).Kind().String()
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
//
// Example:
//
//	{{ hasField "someExistingField" .someStruct }} // Output: true
//	{{ hasField "someNonExistingField" .someStruct }} // Output: false
func (rr *ReflectRegistry) HasField(name string, src any) bool {
	rv := reflect.Indirect(reflect.ValueOf(src))
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
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
// Example:
//
//	{{ {"a":1}, {"a":1} | deepEqual }} // Output: true
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
//
// Example:
//
//	{{ {"name":"John"} | deepCopy }} // Output: {"name":"John"}
func (rr *ReflectRegistry) DeepCopy(element any) any {
	c, err := rr.MustDeepCopy(element)
	if err != nil {
		return nil
	}

	return c
}

func (rr *ReflectRegistry) MustDeepCopy(element any) (any, error) {
	if element == nil {
		return nil, errors.New("element cannot be nil")
	}
	return copystructure.Copy(element)
}
