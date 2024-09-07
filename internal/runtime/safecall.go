package runtime

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrPanicRecovery         = errors.New("cannot safecall function: recovered from a panic")
	ErrInvalidFunction       = errors.New("cannot safecall function: first argument is not a function")
	ErrIncorrectArguments    = errors.New("cannot safecall function: number of arguments does not match function's input arity")
	ErrMoreThanTwoReturns    = errors.New("cannot safecall function: function returns more than two values")
	ErrInvalidLastReturnType = errors.New("cannot safecall function: invalid last return type (expected error)")
)

// SafeCall safely calls a function using reflection. It handles potential
// panics by recovering and returning an error. The function `fn` is expected
// to be a function, and `args` are the arguments to pass to that function.
// It returns the result of the function call (if any) and an error if one
// occurred during the call or if a panic was recovered.
func SafeCall(fn any, args ...any) (result any, err error) {
	// Ensure fn is a function
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()
	if fnValue.Kind() != reflect.Func {
		return nil, ErrInvalidFunction
	}

	// Check if the number of arguments passed matches the function's input arity
	if len(args) != fnType.NumIn() && !fnType.IsVariadic() {
		return nil, fmt.Errorf("%w: expected %d, got %d", ErrIncorrectArguments, fnType.NumIn(), len(args))
	}

	// Defer a function to handle panics
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%w: %v", ErrPanicRecovery, r)
		}
	}()

	// Convert args to reflect.Value slice
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		if arg == nil {
			if fnType.IsVariadic() && i >= fnType.NumIn()-1 {
				// Get the type of the variadic argument's element (e.g., interface{})
				variadicType := fnType.In(fnType.NumIn() - 1).Elem()
				// Create a zero Value of that type
				in[i] = reflect.Zero(variadicType)
			} else {
				// Regular arguments
				in[i] = reflect.Zero(fnType.In(i))
			}
			continue
		}
		in[i] = reflect.ValueOf(arg)
	}

	// Call the function using reflection
	out := fnValue.Call(in)

	switch len(out) {
	case 0:
		return nil, nil
	case 1:
		return out[0].Interface(), nil
	case 2:
		if out[1].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			err, _ = out[1].Interface().(error)
			return out[0].Interface(), err
		}
		return out[0].Interface(), ErrInvalidLastReturnType
	default:
		return out[0].Interface(), ErrMoreThanTwoReturns
	}
}
