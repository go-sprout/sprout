package sprout

import (
	"fmt"
	"strings"

	"github.com/go-sprout/sprout/internal/runtime"
)

// NoticeKind represents the type of notice that can be applied to a function.
// It is an enumeration with different possible values that dictate how the
// notice should behave.
type NoticeKind int

type FunctionNotice struct {
	// FunctionNames is a list of function names to which the notice should be
	// applied. The function names are case-sensitive.
	FunctionNames []string

	// Kind is the kind of the notice
	Kind NoticeKind

	// Message is the message of the notice
	Message string
}

const (
	// NoticeKindDeprecated indicates that the function is deprecated.
	NoticeKindDeprecated NoticeKind = iota + 1
	// NoticeKindInfo indicates that the notice is informational.
	NoticeKindInfo
	// NoticeKindDebug indicates that the notice is for debugging purposes.
	// When using this kind, the notice message can contain the "$out" placeholder
	// which will be replaced with the output of the function.
	NoticeKindDebug
)

// NewNotice creates a new function notice with the given function names, kind,
// and message. The function names are case-sensitive. The kind should be one of
// the predefined NoticeKind values. The message is a string that describes the
// notice.
//
// Example:
//
//	notice := NewNotice(NoticeKindDeprecated, []string{"myFunc"}, "please use myNewFunc instead")
//
// This example creates a new notice that indicates the function "myFunc" is
// deprecated and should be replaced with "myNewFunc" during template rendering.
func NewNotice(kind NoticeKind, functionNames []string, message string) *FunctionNotice {
	return &FunctionNotice{
		FunctionNames: functionNames,
		Kind:          kind,
		Message:       message,
	}
}

// NewDeprecatedNotice creates a new deprecated function notice with the given
// function name and message. The function name is case-sensitive. The message
// is a string that describes what the user should do instead of using the
// deprecated function.
func NewDeprecatedNotice(functionName, message string) *FunctionNotice {
	return NewNotice(NoticeKindDeprecated, []string{functionName}, message)
}

// NewInfoNotice creates a new informational function notice with the given
// function name and message. The function name is case-sensitive. The message
// is a string that provides additional information.
func NewInfoNotice(functionName, message string) *FunctionNotice {
	return NewNotice(NoticeKindInfo, []string{functionName}, message)
}

// NewDebugNotice creates a new debug function notice with the given function
// name and message. The function name is case-sensitive. The message is a
// string that provides additional information for debugging purposes. The
// message can contain the "$out" placeholder, which will be replaced with the
// output of the function.
func NewDebugNotice(functionName, message string) *FunctionNotice {
	return NewNotice(NoticeKindDebug, []string{functionName}, message)
}

// AssignNotices assigns all notices defined in the handler to their original
// functions. This function is used to ensure that all notices are properly
// associated with their original functions in the handler instance.
//
// It should be called after all functions and notices have been added and
// inside the Build function in case of using a custom handler.
func AssignNotices(h Handler) {
	funcs := h.RawFunctions()
	for _, notice := range h.Notices() {
		for _, functionName := range notice.FunctionNames {
			if fn, ok := funcs[functionName]; ok {
				wrappedFn := noticeWrapper(h, notice, functionName, fn)
				funcs[functionName] = wrappedFn
			}
		}
	}
}

// noticeWrapper creates a wrapped function that logs a notice after
// calling the original function. The notice is logged using the handler's
// logger instance. The wrapped function is returned as a wrappedFunc, which
// is a type alias for a function that takes a variadic list of arguments
// and returns an `any` result and an `error`.
func noticeWrapper(h Handler, notice FunctionNotice, functionName string, fn any) wrappedFunction {
	return func(args ...any) (any, error) {
		out, err := runtime.SafeCall(fn, args...)
		switch notice.Kind {
		case NoticeKindDebug:
			h.Logger().With("function", functionName, "notice", "debug").Debug(strings.ReplaceAll(notice.Message, "$out", fmt.Sprint(out)))
		case NoticeKindInfo:
			h.Logger().With("function", functionName, "notice", "info").Info(notice.Message)
		case NoticeKindDeprecated:
			h.Logger().With("function", functionName, "notice", "deprecated").Warn(fmt.Sprintf("Template function `%s` is deprecated: %s", functionName, notice.Message))
		}
		return out, err
	}
}

// WithNotices is used to add one or more function notices to the handler.
// This option allows you to associate a notice with a function, providing
// information about the function's deprecation or other special handling.
//
// The notices are applied to the original function name and its aliases.
// You can use the ApplyOnAliases method on the FunctionNotice to control
// whether the notice should be applied to aliases.
func WithNotices(notices ...*FunctionNotice) HandlerOption[*DefaultHandler] {
	return func(p *DefaultHandler) error {
		// Preallocate the slice if we expect to append multiple notices
		if cap(p.notices) < len(p.notices)+len(notices) {
			newNotices := make([]FunctionNotice, len(p.notices), len(p.notices)+len(notices))
			copy(newNotices, p.notices)
			p.notices = newNotices
		}

		for _, notice := range notices {
			// Skip if the function name is empty or the kind is not valid
			if len(notice.FunctionNames) == 0 || notice.Kind <= 0 {
				continue
			}

			// Append the notice directly without dereferencing
			p.notices = append(p.notices, *notice)
		}

		return nil
	}
}
