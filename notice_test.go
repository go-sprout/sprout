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

	// Apply the WithNotices option with an empty message
	notice3 := NewDebugNotice(originalFunc, "")
	WithNotices(notice3)(handler)

	assert.Contains(t, handler.Notices(), *notice)
	assert.Contains(t, handler.Notices(), *notice2)
	assert.Contains(t, handler.Notices(), *notice3)
	assert.Len(t, handler.notices, 4, "there should be exactly 3 notices")

	// Try to apply a notice with an empty function name.
	notice4 := &FunctionNotice{}
	WithNotices(notice4)(handler)

	// Check that the aliases were not added.
	assert.NotContains(t, handler.Notices(), *notice4)
	assert.Len(t, handler.notices, 4, "there should still be exactly 3 notices")
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
	require.NoError(t, err)
	assert.Equal(t, "cheese", out)
	assert.Contains(t, loggerHandler.messages.String(), "[INFO] amazing")

	out, err = wrappedFunc2()
	require.NoError(t, err)
	assert.Equal(t, "cheese", out)
	assert.Contains(t, loggerHandler.messages.String(), "[WARN] Template function `originalFunc` is deprecated: oh no")

	out, err = wrappedFunc3()
	require.NoError(t, err)
	assert.Equal(t, "cheese", out)
	assert.Contains(t, loggerHandler.messages.String(), "[DEBUG] Nice this function returns cheese")
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
	assert.Equal(t, "[INFO] amazing\n", loggerHandler.messages.String())
}
