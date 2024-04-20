// Package errors implements a sophisticated error handling system that extends
// Go's standard error capabilities. It supports error wrapping, stack tracing,
// and linking errors into chains to provide rich context and debugging
// information.
//
// Example:
//
//	func doSomething() error {
//		err := doSomethingElse()
//		if err != nil {
//			return errors.New("failed to do something", err)
//		}
//		return nil
//	}
//
//	func doSomethingElse() error {
//		return errors.New("failed to do something else")
//	}
//
//	func main() {
//		err := doSomething()
//		if err = errors.Cast(err); err != nil {
//			fmt.Println(err)
//			fmt.Println("Stack trace:")
//			for _, line := range err.Stack() {
//				fmt.Println(line)
//			}
//		}
//	}
//
// The errors package provides a rich set of functions for creating, wrapping,
// and handling errors. It allows you to create detailed error messages, capture
// stack traces, and link errors together to form a chain. This can be useful for
// debugging and logging errors in complex applications.
//
// The errors package is designed to be compatible with Go's standard error
// interface, so you can use it seamlessly with existing code. It provides
// additional functionality on top of the standard error interface, making it
// easy to work with errors in a more sophisticated way.
package errors

import (
	"errors"
	"path"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

// Error is an interface that models an error with extended functionality over
// Go's built-in error interface. It supports unwrapping, checking error equality,
// and retrieving a stack trace.
type Error interface {
	error

	Err() error
	Unwrap() error
	Is(err error) bool
	Stack() []string
}

// ErrorChain defines additional behavior for errors that can be linked into
// a chain, providing access to previous and next errors in the chain, as well
// as the ability to determine the root cause of the chain.
type ErrorChain interface {
	Error

	Prev() ErrorChain
	Next() ErrorChain
	Cause() ErrorChain
}

// Stackliteable is an interface for errors that include a stack trace line.
// It allows retrieving a formatted string representing where the error occurred.
type Stackliteable interface {
	StackLine() string
	WithFrameSkip(skip int) Stackliteable
}

// errorStruct is the concrete implementation of Error and ErrorChain interfaces.
// It links together errors in a chain, capturing each error's context and
// stack trace.
type errorStruct struct {
	prev *errorStruct
	next *errorStruct
	err  error

	stackliteSkip int
	stacklite     *stacklite
}

// stacklite contains details about the stack at the point where the error
// was captured.
type stacklite struct {
	pack string
	fn   string
	file string
	line int
}

// New creates a new error instance with the specified text message. It captures
// the current stack trace and optionally links the new error with previous errors
// to form a chain.
// The optional previousErrs parameter allows linking the new error to existing
// errors, forming a chain.
//
// WARNING: if you want to keep the err comparable with errors.Is, you should
// use the Cast function to wrap the error not `New(err.Error())`.
func New(text string, previousErrs ...error) Error {
	var prev *errorStruct

	if len(previousErrs) > 0 {
		prev = castToErrorStruct(previousErrs[0])
	}

	err := &errorStruct{
		prev:          prev,
		err:           errors.New(text),
		stackliteSkip: 3,
	}

	if prev != nil {
		prev.next = err
	}

	return err
}

// castToErrorStruct attempts to cast a generic error to *errorStruct.
// If the cast is unsuccessful, it wraps the error in a new *errorStruct.
func castToErrorStruct(err error) *errorStruct {
	if err == nil {
		return nil
	}

	if e, ok := err.(*errorStruct); ok {
		return e
	}
	return &errorStruct{err: err}
}

// Cast ensures that any error is converted into an Error interface. If the error
// is already an Error, it returns it directly; otherwise, it wraps the error.
func Cast(err error, previousErrs ...error) Error {
	if err == nil {
		return nil
	}

	var prev *errorStruct
	if len(previousErrs) > 0 {
		prev = castToErrorStruct(previousErrs[0])
	}

	if e, ok := err.(Error); ok {
		if es, ok := e.(*errorStruct); ok {
			es.prev = prev
			return es
		}

		return e
	}

	return &errorStruct{
		err:  err,
		prev: prev,
	}
}

// Error returns a string representation of the error chain from the current error
// back to the root cause, including the stack trace for each error.
func (e *errorStruct) Error() string {
	var b strings.Builder
	b.Grow(256) // Pre-allocate to avoid reallocations

	curr := e.Cause().(*errorStruct)
	for curr != nil {
		if sl := curr.StackLine(); sl != "" {
			b.WriteString(curr.StackLine())
			b.WriteString(": ")
		}

		if curr.err != nil {
			b.WriteString(curr.err.Error())
		}

		// Safely move to the next error if it exists, otherwise break the loop
		if curr.next != nil {
			curr = curr.next
			b.WriteString(" > ")
		} else {
			break
		}
	}

	return b.String()
}

// Is determines if the error matches or contains the specified error anywhere
// in the error chain. Following the chain, it checks each error for equality
// with the target error. Following the go standard.
func (e *errorStruct) Is(err error) bool {
	root := e
	for root != nil {
		switch err.(type) {
		case *errorStruct:
			if root == err {
				return true
			}

		case Error:
			if errors.Is(root.err, err) {
				return true
			}
		}

		if errors.Is(root.err, err) {
			return true
		}
		root = root.prev
	}
	return false
}

// Err returns the specific error held within this errorStruct. This return the
// go standard error and not the errorStruct.
func (e *errorStruct) Err() error {
	return e.err
}

// Unwrap provides compatibility with Go's error unwrapping scheme, returning
// the previous error in the chain if available. This return the go standard
// error and not the errorStruct.
func (e *errorStruct) Unwrap() error {
	if e.prev == nil {
		return nil
	}
	return e.prev.err
}

// Stack constructs a slice of strings that represents the error chain with
// each element corresponding to an error's stack line and message.
// The stack is constructed from the root cause to the current error.
// Each element in the slice contains the stack line and error message.
// The stack line is formatted as `[package.file#line function] error message`.
// If the stack line is not available, it is omitted from the output.
func (e *errorStruct) Stack() []string {
	// Estimate the depth of the error chain for initial slice capacity
	capacity := 0
	for curr := e; curr != nil; curr = curr.prev {
		capacity++
	}

	// Preallocate the slice with the exact needed capacity
	stack := make([]string, 0, capacity)
	for curr := e; curr != nil; curr = curr.prev {
		var b strings.Builder
		// Estimate the length of the final string to minimize reallocations
		line := curr.StackLine()
		err := ""
		if curr.err != nil {
			err = curr.err.Error()
		}
		b.Grow(len(line) + len(err) + 2) // Plus two for possible ": " separator

		if line != "" {
			b.WriteString(line)
			b.WriteString(": ")
		}
		b.WriteString(err)

		stack = append(stack, b.String())
	}

	slices.Reverse(stack)
	return stack
}

func (e *errorStruct) WithFrameSkip(skip int) Stackliteable {
	if skip < 2 {
		skip = 2
	}
	e.stackliteSkip = skip
	return e
}

// Cause returns the root error in the chain, providing access to the initial error
// that triggered the sequence of errors. This method traverses the chain to find
// the root cause.
func (e *errorStruct) Cause() ErrorChain {
	root := e
	for root != nil {
		if root.Prev() == nil {
			break
		}
		root = root.prev
	}
	return root
}

// Prev returns the previous error in the chain or nil if there is no previous
// error. This method allows traversing the chain in reverse order to access
// previous errors (until the root cause is reached).
func (e *errorStruct) Prev() ErrorChain {
	if e.prev == nil {
		return nil
	}

	return e.prev
}

// Next returns the next error in the chain or nil if there is no next error.
// This method allows traversing the chain in the forward direction to access
// subsequent errors.
func (e *errorStruct) Next() ErrorChain {
	if e.next == nil {
		return nil
	}
	return e.next
}

// StackLine constructs a formatted string containing the error's location in the code,
// including the package, file, and line number where the error occurred.
// The format is `[package.file#line function]`.
func (e *errorStruct) StackLine() string {
	if e.stacklite == nil {
		e.stacklite = errFuncCaller(e.stackliteSkip)
	}

	b := strings.Builder{}
	b.Grow(128) // Pre-allocate based on typical stack line length

	b.WriteRune('[')
	b.WriteString(e.stacklite.pack)
	b.WriteRune('.')
	b.WriteString(e.stacklite.file)
	b.WriteRune('#')
	b.WriteString(strconv.Itoa(e.stacklite.line))
	b.WriteRune(' ')
	b.WriteString(e.stacklite.fn)
	b.WriteRune(']')

	return b.String()
}

// errFuncCaller uses the runtime package to find the function that called it,
// allowing for detailed logging and error handling. 'skip' levels are bypassed
// to find the actual caller.
// It returns a *runtime.Func representing the caller, or nil if not found.
func errFuncCaller(skip int) *stacklite {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return nil
	}

	fn := runtime.FuncForPC(pc)
	fullFuncName := fn.Name()

	lastSlashIndex := strings.LastIndex(fullFuncName, "/")
	lastDotIndex := strings.LastIndex(fullFuncName, ".")

	var pack, funcName string
	if lastDotIndex > lastSlashIndex {
		pack = fullFuncName[lastSlashIndex+1 : lastDotIndex]
		funcName = fullFuncName[lastDotIndex+1:]
	} else {
		funcName = fullFuncName
	}

	file, line := fn.FileLine(pc)
	_, file = path.Split(file)

	return &stacklite{pack, funcName, file, line}
}
