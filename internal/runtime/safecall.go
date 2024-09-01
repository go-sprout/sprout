package runtime

import (
	"errors"
	"fmt"
	"reflect"
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
		return nil, errors.New("fn is not a function")
	}

	// Defer a function to handle panics
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
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

	// Process the output
	if len(out) == 0 {
		return nil, nil
	}

	// If there's only one return value
	result = out[0].Interface()
	if len(out) == 1 {
		return result, nil
	}

	// If there are two return values (assuming the second is an error)
	if len(out) == 2 && out[1].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		err, _ = out[1].Interface().(error)
		return result, err
	}

	// Handle other cases as needed
	return result, nil
}
