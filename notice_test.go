package sprout

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type noticeLoggerHandler struct {
	messages bytes.Buffer
}

func (h *noticeLoggerHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *noticeLoggerHandler) Handle(_ context.Context, r slog.Record) error {
	msg := fmt.Sprintf("[%s] %s\n", r.Level.String(), r.Message)
	h.messages.Write([]byte(msg))
	return nil
}

func (h *noticeLoggerHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *noticeLoggerHandler) WithGroup(name string) slog.Handler {
	return h
}

func TestWithNotice(t *testing.T) {
	handler := New()
	originalFunc := "originalFunc"
	notice := NewInfoNotice(originalFunc, "amazing")

	// Apply the WithNotices option with one notice.
	WithNotices(notice)(handler)

	// Check that the aliases were added.
	assert.Contains(t, handler.Notices(), *notice)
	assert.Len(t, handler.notices, 1, "there should be exactly 1 notice")

	// Apply the WithNotices option with multiple notices.
	notice2 := NewDeprecatedNotice(originalFunc, "oh no")
	WithNotices(notice, notice2)(handler)

	// Check that the aliases were added.
	assert.Contains(t, handler.Notices(), *notice)
	assert.Contains(t, handler.Notices(), *notice2)
	assert.Len(t, handler.notices, 3, "there should be exactly 3 notices")

	// Try to apply a notice with an empty function name.
	notice3 := &FunctionNotice{}
	WithNotices(notice3)(handler)

	// Check that the aliases were not added.
	assert.NotContains(t, handler.Notices(), *notice3)
	assert.Len(t, handler.notices, 3, "there should still be exactly 3 notices")
}

func TestAssignNotices(t *testing.T) {
	handler := New()
	originalFunc := "originalFunc"
	notice := NewInfoNotice(originalFunc, "amazing")

	// Mock a function for originalFunc and add it to funcsRegistry.
	mockFunc := func() string { return "cheese" }
	handler.cachedFuncsMap[originalFunc] = mockFunc

	// Assign the notices directly.
	handler.notices = []FunctionNotice{*notice}
	AssignNotices(handler)

	// Check that the aliases were added.
	assert.Contains(t, handler.Notices(), *notice)
	assert.Len(t, handler.notices, 1, "there should be exactly 1 notice")

	require.Contains(t, handler.Functions(), originalFunc)
	assert.NotEqual(t, reflect.ValueOf(mockFunc).Pointer(), reflect.ValueOf(handler.Functions()[originalFunc]).Pointer(), "the function should have been wrapped")
}

func TestCreateWrappedFunction(t *testing.T) {
	loggerHandler := &noticeLoggerHandler{}
	handler := New(WithLogger(slog.New(loggerHandler)))

	originalFunc := "originalFunc"
	mockFunc := func() string { return "cheese" }

	// Create a wrapped function.
	wrappedFunc := createWrappedFunction(handler, *NewInfoNotice(originalFunc, "amazing"), originalFunc, mockFunc)
	wrappedFunc2 := createWrappedFunction(handler, *NewDeprecatedNotice(originalFunc, "oh no"), originalFunc, mockFunc)
	wrappedFunc3 := createWrappedFunction(handler, *NewNotice(NoticeKindDebug, []string{originalFunc}, "Nice this function returns $out"), originalFunc, mockFunc)

	// Call the wrapped function.
	out, err := wrappedFunc()
	assert.NoError(t, err)
	assert.Equal(t, "cheese", out)
	assert.Contains(t, loggerHandler.messages.String(), "[INFO] amazing")

	out, err = wrappedFunc2()
	assert.NoError(t, err)
	assert.Equal(t, "cheese", out)
	assert.Contains(t, loggerHandler.messages.String(), "[WARN] Template function `originalFunc` is deprecated: oh no")

	out, err = wrappedFunc3()
	assert.NoError(t, err)
	assert.Equal(t, "cheese", out)
	assert.Contains(t, loggerHandler.messages.String(), "[DEBUG] Nice this function returns cheese")
}

func TestSafeCall(t *testing.T) {
	// Test a function that returns a string.
	fn := func() (string, error) { return "cheese", nil }
	out, err := safeCall(fn)
	assert.NoError(t, err)
	assert.Equal(t, "cheese", out)

	// Test a function that returns a string and an error.
	fn2 := func() (string, error) { return "cheese", fmt.Errorf("oh no") }
	out, err = safeCall(fn2)
	assert.Error(t, err)
	assert.Equal(t, "cheese", out)

	// Test a function that returns a string and an error.
	fn3 := func() (string, error) { return "", fmt.Errorf("oh no") }
	out, err = safeCall(fn3)
	assert.Error(t, err)
	assert.Empty(t, out)

	// Test a function that returns a string and an error.
	fn4 := func() (string, error) { return "", nil }
	out, err = safeCall(fn4)
	assert.NoError(t, err)
	assert.Empty(t, out)

	// Test a function that returns nothing.
	fn5 := func() {}
	out, err = safeCall(fn5)
	assert.NoError(t, err)
	assert.Nil(t, out)

	// Test a function that returns 3 values.
	a, b, c := "a", "b", "c"
	fn6 := func(a, b, c string) (string, string, string) { return a, b, c }
	out, err = safeCall(fn6, a, b, c)
	assert.NoError(t, err)
	assert.Equal(t, out, a, "the return should be the first argument")

	// Test a case where the function panics.
	fn7 := func() { panic("oh no") }
	out, err = safeCall(fn7)
	assert.ErrorContains(t, err, "recovered from panic: oh no")
	assert.Nil(t, out)

	// Test when fn is not a function.
	fn8 := "cheese"
	out, err = safeCall(fn8)
	assert.ErrorContains(t, err, "fn is not a function")
	assert.Nil(t, out)
}

func TestNoticeInTemplate(t *testing.T) {
	loggerHandler := &noticeLoggerHandler{}
	handler := New(WithLogger(slog.New(loggerHandler)))

	originalFunc := "originalFunc"
	mockFunc := func() string { return "cheese" }
	handler.cachedFuncsMap[originalFunc] = mockFunc

	// Add a notice to the handler.
	AddNotice(&handler.notices, NewInfoNotice(originalFunc, "amazing"))

	// Create a template with the function.
	tmpl, err := template.New("test").Funcs(handler.Build()).Parse("{{- originalFunc -}}")
	require.NoError(t, err)

	// Execute the template.
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	require.NoError(t, err)
	assert.Equal(t, "cheese", buf.String())
	assert.Equal(t, loggerHandler.messages.String(), "[INFO] amazing\n")
}
