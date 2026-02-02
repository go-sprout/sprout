package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDigIntoDictWithNoKeys(t *testing.T) {
	_, err := NewRegistry().digIntoDict(map[string]any{}, []string{})
	assert.ErrorContains(t, err, "unexpected termination of key traversal")
}

func TestSplitKeysWithEscapes(t *testing.T) {
	mr := NewRegistry()

	tests := []struct {
		name     string
		input    []string
		expected []string
		wantErr  string
	}{
		{
			name:     "simple key",
			input:    []string{"a"},
			expected: []string{"a"},
		},
		{
			name:     "dotted key splits",
			input:    []string{"a.b"},
			expected: []string{"a", "b"},
		},
		{
			name:     "multiple dotted keys",
			input:    []string{"a.b", "c.d"},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "escaped dot preserves literal dot",
			input:    []string{`a\.b`},
			expected: []string{"a.b"},
		},
		{
			name:     "escaped backslash",
			input:    []string{`a\\b`},
			expected: []string{`a\b`},
		},
		{
			name:     "mixed escaped and unescaped dots",
			input:    []string{`a\.b.c`},
			expected: []string{"a.b", "c"},
		},
		{
			name:     "multiple escaped dots",
			input:    []string{`a\.b\.c`},
			expected: []string{"a.b.c"},
		},
		{
			name:     "escaped backslash before dot",
			input:    []string{`a\\.b`},
			expected: []string{`a\`, "b"},
		},
		{
			name:     "complex case",
			input:    []string{`servers.api\.internal.port`},
			expected: []string{"servers", "api.internal", "port"},
		},
		{
			name:    "invalid escape sequence",
			input:   []string{`a\nb`},
			wantErr: "invalid escape sequence: \\n",
		},
		{
			name:    "trailing backslash",
			input:   []string{`abc\`},
			wantErr: "trailing backslash",
		},
		{
			name:     "empty key",
			input:    []string{""},
			expected: []string{""},
		},
		{
			name:     "fast path no backslash",
			input:    []string{"simple.dotted.path"},
			expected: []string{"simple", "dotted", "path"},
		},
		{
			name:    "fast path consecutive dots",
			input:   []string{"a..b"},
			wantErr: "empty key segment",
		},
		{
			name:    "fast path leading dot",
			input:   []string{".abc"},
			wantErr: "empty key segment",
		},
		{
			name:    "fast path trailing dot",
			input:   []string{"abc."},
			wantErr: "empty key segment",
		},
		{
			name:     "single empty key is valid",
			input:    []string{""},
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := mr.splitKeysWithEscapes(tt.input)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestSplitKeyOnUnescapedDots(t *testing.T) {
	mr := NewRegistry()

	tests := []struct {
		name     string
		input    string
		expected []string
		wantErr  string
	}{
		{
			name:     "no dots",
			input:    "abc",
			expected: []string{"abc"},
		},
		{
			name:     "single dot",
			input:    "a.b",
			expected: []string{"a", "b"},
		},
		{
			name:     "escaped dot",
			input:    `a\.b`,
			expected: []string{"a.b"},
		},
		{
			name:     "escaped backslash",
			input:    `a\\b`,
			expected: []string{`a\b`},
		},
		{
			name:     "escaped backslash then dot",
			input:    `a\\.b`,
			expected: []string{`a\`, "b"},
		},
		{
			name:     "multiple segments",
			input:    `a.b\.c.d`,
			expected: []string{"a", "b.c", "d"},
		},
		{
			name:    "invalid escape",
			input:   `a\nb`,
			wantErr: "invalid escape sequence: \\n",
		},
		{
			name:    "trailing backslash",
			input:   `abc\`,
			wantErr: "trailing backslash",
		},
		{
			name:    "consecutive dots",
			input:   "a..b",
			wantErr: "empty key segment",
		},
		{
			name:    "leading dot",
			input:   ".abc",
			wantErr: "empty key segment",
		},
		{
			name:    "trailing dot",
			input:   "abc.",
			wantErr: "empty key segment",
		},
		{
			name:    "only dots",
			input:   "..",
			wantErr: "empty key segment",
		},
		{
			name:     "single key no dots",
			input:    "single",
			expected: []string{"single"},
		},
		{
			name:     "empty string is valid single key",
			input:    "",
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := mr.splitKeyOnUnescapedDots(tt.input)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
