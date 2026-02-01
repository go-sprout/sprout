package compatibility

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/go-sprout/sprout/sprigin"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

// skippedTests contains test codes that are intentionally skipped because they
// represent acceptable behavioral differences between Sprig and Sprout.
// Each entry documents why the difference is acceptable.
var skippedTests = map[string]string{
	// Unicode handling fixes: Sprout correctly handles Unicode characters
	// while Sprig has byte-level bugs with multi-byte characters.
	"CODE_000112": "Unicode fix: nospace correctly handles multi-byte Unicode characters",
	"CODE_000131": "Unicode fix: abbrev correctly handles multi-byte Unicode characters",
}

// data contains deterministic test data for all types used in compatibility testing.
// This map provides consistent, reproducible values across test runs to enable
// regression detection. All values are carefully chosen to be deterministic.
var data = map[string]any{
	// =====================================================================
	// STRING VALUES - Various string patterns for testing string functions
	// =====================================================================
	"emptyString":       "",
	"simpleString":      "hello",
	"spacedString":      "hello world",
	"trimString":        "   hello   ",
	"tabNewlineString":  "\t\nhello\t\n",
	"unicodeString":     "αβγδεζηθ",
	"mixedCaseString":   "HeLLo WoRLD",
	"upperString":       "HELLO WORLD",
	"lowerString":       "hello world",
	"snakeCaseString":   "hello_world_test",
	"camelCaseString":   "helloWorldTest",
	"kebabCaseString":   "hello-world-test",
	"pascalCaseString":  "HelloWorldTest",
	"pathString":        "/path/to/file.txt",
	"urlString":         "https://example.com/path?query=value#anchor",
	"jsonString":        `{"key": "value", "num": 123, "bool": true}`,
	"base64String":      "SGVsbG8gV29ybGQh",
	"base32String":      "JBSWY3DPEHPK3PXP",
	"regexPattern":      "[a-zA-Z]+[0-9]*",
	"specialCharsStr":   "hello!@#$%^&*()world",
	"multilineString":   "line1\nline2\nline3",
	"quotedString":      `"quoted"`,
	"singleQuotedStr":   `'single'`,
	"htmlString":        "<div>Hello &amp; World</div>",
	"numericString":     "12345",
	"floatString":       "3.14159",
	"semverString":      "1.2.3-alpha.1+build.456",
	"dateString":        "2023-06-15",
	"timeString":        "10:30:00",
	"datetimeString":    "2023-06-15T10:30:00Z",
	"durationString":    "2h30m45s",
	"ipv4String":        "192.168.1.100",
	"ipv6String":        "::1",
	"macString":         "00:1A:2B:3C:4D:5E",
	"uuidString":        "550e8400-e29b-41d4-a716-446655440000",
	"longString":        "This is a longer string that can be used for testing truncation, wrapping, and other string manipulation functions that work with longer text content.",

	// =====================================================================
	// INTEGER VALUES - Various integer patterns for testing math functions
	// =====================================================================
	"zeroInt":         0,
	"oneInt":          1,
	"negativeOneInt":  -1,
	"positiveInt":     42,
	"negativeInt":     -42,
	"largeInt":        999999,
	"smallInt":        -999999,
	"maxInt":          9223372036854775807,
	"minInt":          -9223372036854775808,
	"powerOfTwoInt":   1024,
	"primeInt":        17,
	"evenInt":         100,
	"oddInt":          101,
	"octalModeInt":    0755,
	"hexInt":          0xFF,
	"percentInt":      50,
	"countInt":        10,
	"indexInt":        5,
	"offsetInt":       3,
	"limitInt":        20,

	// =====================================================================
	// FLOAT VALUES - Various float patterns for testing float functions
	// =====================================================================
	"zeroFloat":       0.0,
	"oneFloat":        1.0,
	"negativeFloat":   -1.0,
	"piFloat":         3.14159265359,
	"eFloat":          2.71828182846,
	"goldenRatio":     1.61803398875,
	"halfFloat":       0.5,
	"quarterFloat":    0.25,
	"largeFloat":      999999.999999,
	"smallFloat":      0.000001,
	"negLargeFloat":   -999999.999999,
	"roundUpFloat":    1.9,
	"roundDownFloat":  1.1,
	"roundMidFloat":   1.5,
	"precisionFloat":  123.456789,
	"scientificFloat": 1.23e10,

	// =====================================================================
	// BOOLEAN VALUES - For testing logic and conditional functions
	// =====================================================================
	"trueVal":  true,
	"falseVal": false,

	// =====================================================================
	// NIL VALUE - For testing empty/default handling
	// =====================================================================
	"nilVal": nil,

	// =====================================================================
	// LIST/SLICE VALUES - Various list patterns for testing list functions
	// =====================================================================
	"emptyList":       []any{},
	"singleItemList":  []any{"only"},
	"intList":         []any{1, 2, 3, 4, 5},
	"stringList":      []any{"apple", "banana", "cherry"},
	"mixedList":       []any{1, "two", 3.0, true, nil},
	"nestedList":      []any{[]any{1, 2}, []any{3, 4}, []any{5, 6}},
	"duplicateList":   []any{1, 1, 2, 2, 3, 3},
	"boolList":        []any{true, false, true},
	"floatList":       []any{1.1, 2.2, 3.3, 4.4, 5.5},
	"unicodeList":     []any{"αβγ", "δεζ", "ηθι"},
	"sortableList":    []any{"zebra", "apple", "mango", "banana"},
	"numberedList":    []any{"item1", "item2", "item3", "item4", "item5"},
	"longList":        []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	"sparseList":      []any{1, "", 2, nil, 3, "", 4},
	"pathList":        []any{"path", "to", "file", "txt"},

	// =====================================================================
	// MAP/DICT VALUES - Various map patterns for testing dict functions
	// =====================================================================
	"emptyDict": map[string]any{},
	"simpleDict": map[string]any{
		"key1": "value1",
		"key2": "value2",
	},
	"intDict": map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
	},
	"nestedDict": map[string]any{
		"level1": map[string]any{
			"level2": map[string]any{
				"level3": "deep value",
			},
		},
	},
	"mixedDict": map[string]any{
		"string":  "hello",
		"int":     42,
		"float":   3.14,
		"bool":    true,
		"list":    []any{1, 2, 3},
		"nil":     nil,
		"unicode": "αβγ",
	},
	"personDict": map[string]any{
		"firstName": "John",
		"lastName":  "Doe",
		"age":       30,
		"email":     "john.doe@example.com",
		"active":    true,
	},
	"configDict": map[string]any{
		"host":     "localhost",
		"port":     8080,
		"debug":    false,
		"timeout":  30,
		"protocol": "https",
	},
	"unicodeDict": map[string]any{
		"αβγ": "delta",
		"δεζ": "epsilon",
	},

	// =====================================================================
	// FIXED TIME VALUE - Deterministic time for date/time function testing
	// Use a fixed time: 2023-06-15 10:30:00 UTC
	// =====================================================================
	"fixedTime": time.Date(2023, 6, 15, 10, 30, 0, 0, time.UTC),

	// =====================================================================
	// BYTE ARRAYS - For testing encoding functions
	// =====================================================================
	"emptyBytes":  []byte{},
	"simpleBytes": []byte("hello world"),
	"binaryBytes": []byte{0x00, 0x01, 0x02, 0xFF, 0xFE, 0xFD},

	// =====================================================================
	// CRYPTO TEST VALUES - Deterministic values for crypto testing
	// =====================================================================
	"aes16Key":      "0123456789abcdef",
	"aes24Key":      "0123456789abcdef01234567",
	"aes32Key":      "0123456789abcdef0123456789abcdef",
	"plaintext":     "secret message",
	"password":      "P@ssw0rd123!",
	"username":      "testuser",
	"salt":          "randomsalt",
	"caCommonName":  "Test CA",
	"certCommonName": "test.example.com",
	"certValidDays": 365,

	// =====================================================================
	// NETWORK VALUES - For testing network-related functions
	// =====================================================================
	"localhost":    "localhost",
	"localhostIP":  "127.0.0.1",
	"exampleHost":  "example.com",
	"internalHost": "internal.local",

	// =====================================================================
	// PATH VALUES - For testing path manipulation functions
	// =====================================================================
	"absolutePath": "/usr/local/bin/program",
	"relativePath": "relative/path/to/file.go",
	"dirtyPath":    "foo/bar/../baz/./qux",
	"unixPath":     "/home/user/documents",
	"fileName":     "document.pdf",
	"extension":    ".txt",
}

// codePattern matches lines containing test codes like "[CODE_000001]"
var codePattern = regexp.MustCompile(`\[CODE_(\d+)\]`)

func TestCompatibility(t *testing.T) {
	var bufSprig, bufSprigin bytes.Buffer
	eg, _ := errgroup.WithContext(t.Context())
	eg.Go(func() error {
		tmplSprig, err := template.New("compatibility").Funcs(sprig.FuncMap()).ParseFiles("compatibility.tmpl")
		if err != nil {
			return fmt.Errorf("could not parse sprig template: %w", err)
		}

		err = tmplSprig.ExecuteTemplate(&bufSprig, "compatibility.tmpl", data)
		if err != nil {
			return fmt.Errorf("executing sprig template: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		tmplSprigin, err := template.New("compatibility").Funcs(sprigin.FuncMap()).ParseFiles("compatibility.tmpl")
		if err != nil {
			return fmt.Errorf("could not parse sprigin template: %w", err)
		}
		err = tmplSprigin.ExecuteTemplate(&bufSprigin, "compatibility.tmpl", data)
		if err != nil {
			return fmt.Errorf("executing sprigin template: %w", err)
		}
		return nil
	})

	require.NoError(t, eg.Wait())

	// Compare line by line, skipping known acceptable differences
	sprigLines := strings.Split(bufSprig.String(), "\n")
	spriginLines := strings.Split(bufSprigin.String(), "\n")

	require.Equal(t, len(sprigLines), len(spriginLines), "Output line counts differ")

	var skippedCount int
	var failedCodes []string

	for i := range sprigLines {
		sprigLine := sprigLines[i]
		spriginLine := spriginLines[i]

		if sprigLine == spriginLine {
			continue
		}

		// Extract test code from the line
		matches := codePattern.FindStringSubmatch(sprigLine)
		if len(matches) < 2 {
			// No code found, this is a structural difference
			t.Errorf("Line %d differs (no code):\n  sprig:   %q\n  sprigin: %q", i+1, sprigLine, spriginLine)
			continue
		}

		code := "CODE_" + matches[1]
		if reason, skipped := skippedTests[code]; skipped {
			skippedCount++
			t.Logf("Skipped %s: %s", code, reason)
			continue
		}

		failedCodes = append(failedCodes, code)
		t.Errorf("[%s] Line %d differs:\n  sprig:   %q\n  sprigin: %q", code, i+1, sprigLine, spriginLine)
	}

	if skippedCount > 0 {
		t.Logf("Skipped %d tests with acceptable differences", skippedCount)
	}

	if len(failedCodes) > 0 {
		t.Errorf("Failed codes: %v", failedCodes)
	}
}
