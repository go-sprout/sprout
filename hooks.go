package sprout

import (
	"fmt"
	"reflect"
)

var errorType = reflect.TypeFor[error]()
var reflectValueType = reflect.TypeFor[reflect.Value]()

type HookedHandler interface {
	Handler

	// ExecFunction executes the function with the given name and arguments.
	// The function name is case-sensitive.
	ExecFunction(name string, args ...any) (any, error)

	// OnCastHook is called before the function is executed, allowing you to
	// apply any necessary actions before the function is called.
	OnCastHook(name string, args ...any) error

	// OnSuccessHook is called after the function is executed successfully, allowing
	// you to apply any necessary actions after the function is called.
	OnSuccessHook(name string, args ...any) error

	// OnFailureHook is called after the function is executed with an error, allowing
	// you to apply any necessary actions after the function is called.
	OnFailureHook(name string, args ...any) error
}

func (d *DefaultHandler) OnCastHook(name string, args ...any) error {
	return nil
}

func (d *DefaultHandler) OnSuccessHook(name string, args ...any) error {
	return nil
}

func (d *DefaultHandler) OnFailureHook(name string, args ...any) error {
	return nil
}

func (d *DefaultHandler) ExecFunction(name string, args ...reflect.Value) (any, error) {
	d.rwmu.RLock()
	fn := d.cachedFuncsMap[name]
	d.rwmu.RUnlock()

	fn = indirectInterface(fn)
	if !fn.IsValid() {
		return nil, fmt.Errorf("call of nil")
	}
	typ := fn.Type()
	if typ.Kind() != reflect.Func {
		return nil, fmt.Errorf("non-function %s of type %s", name, typ)
	}

	if err := goodFunc(name, typ); err != nil {
		return nil, err
	}
	numIn := typ.NumIn()
	var dddType reflect.Type
	if typ.IsVariadic() {
		if len(args) < numIn-1 {
			return nil, fmt.Errorf("wrong number of args for %s: got %d want at least %d", name, len(args), numIn-1)
		}
		dddType = typ.In(numIn - 1).Elem()
	} else {
		if len(args) != numIn {
			return nil, fmt.Errorf("wrong number of args for %s: got %d want %d", name, len(args), numIn)
		}
	}
	argv := make([]reflect.Value, len(args))
	for i, arg := range args {
		arg = indirectInterface(arg)
		// Compute the expected type. Clumsy because of variadics.
		argType := dddType
		if !typ.IsVariadic() || i < numIn-1 {
			argType = typ.In(i)
		}

		var err error
		if argv[i], err = prepareArg(arg, argType); err != nil {
			return nil, fmt.Errorf("arg %d: %w", i, err)
		}
	}
	out, err := safeCall(fn, argv)
	return out.Interface(), err
}

// Optimized version of reflectValueOf that avoids unnecessary allocations
func safeReflectValueOf(arg any) reflect.Value {
	if arg == nil {
		return reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
	}

	return reflect.ValueOf(arg)
}

// goodFunc reports whether the function or method has the right result signature.
func goodFunc(name string, typ reflect.Type) error {
	// We allow functions with 1 result or 2 results where the second is an error.
	switch numOut := typ.NumOut(); {
	case numOut == 1:
		return nil
	case numOut == 2 && typ.Out(1) == errorType:
		return nil
	case numOut == 2:
		return fmt.Errorf("invalid function signature for %s: second return value should be error; is %s", name, typ.Out(1))
	default:
		return fmt.Errorf("function %s has %d return values; should be 1 or 2", name, typ.NumOut())
	}
}

// safeCall runs fun.Call(args), and returns the resulting value and error, if
// any. If the call panics, the panic value is returned as an error.
func safeCall(fun reflect.Value, args []reflect.Value) (val reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	ret := fun.Call(args)
	if len(ret) == 2 && !ret[1].IsNil() {
		return ret[0], ret[1].Interface().(error)
	}
	return ret[0], nil
}

// indirectInterface returns the concrete value in an interface value,
// or else the zero reflect.Value.
// That is, if v represents the interface value x, the result is the same as reflect.ValueOf(x):
// the fact that x was an interface value is forgotten.
func indirectInterface(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Interface {
		return v
	}
	if v.IsNil() {
		return reflect.Value{}
	}
	return v.Elem()
}

// prepareArg checks if value can be used as an argument of type argType, and
// converts an invalid value to appropriate zero if possible.
func prepareArg(value reflect.Value, argType reflect.Type) (reflect.Value, error) {
	if !value.IsValid() {
		if !canBeNil(argType) {
			return reflect.Value{}, fmt.Errorf("value is nil; should be of type %s", argType)
		}
		value = reflect.Zero(argType)
	}
	if value.Type().AssignableTo(argType) {
		return value, nil
	}
	if intLike(value.Kind()) && intLike(argType.Kind()) && value.Type().ConvertibleTo(argType) {
		value = value.Convert(argType)
		return value, nil
	}
	return reflect.Value{}, fmt.Errorf("value has type %s; should be %s", value.Type(), argType)
}

// canBeNil reports whether an untyped nil can be assigned to the type. See reflect.Zero.
func canBeNil(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return true
	case reflect.Struct:
		return typ == reflectValueType
	}
	return false
}

func intLike(typ reflect.Kind) bool {
	switch typ {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	}
	return false
}
