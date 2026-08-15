package uniqueid_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/uniqueid"
)

const (
	// uuidv7Fixture is a version 7 UUID generated at 2024-05-07T15:04:05Z, with
	// every random bit set to zero to keep it deterministic.
	uuidv7Fixture = "018f5395-4e88-7000-8000-000000000000"
	// uuidv1Fixture is a version 1 UUID generated at 2024-05-07T15:04:05Z, with
	// every random bit set to zero to keep it deterministic.
	uuidv1Fixture = "0bcce080-0c83-11ef-8000-000000000000"
	// uuidv5Fixture is the UUID of the name "python.org" in the dns namespace.
	uuidv5Fixture = "886313e1-3b8a-5372-9b90-0c9aee199e5d"
)

func TestUuidv4(t *testing.T) {
	tc := []pesticide.RegexpTestCase{
		{Name: "TestUuidv4", Template: `{{uuidv4}}`, Regexp: `^[\da-f]{8}-[\da-f]{4}-4[\da-f]{3}-[\da-f]{4}-[\da-f]{12}$`, Length: 36},
	}

	pesticide.RunRegexpTestCases(t, uniqueid.NewRegistry(), tc)
}

func TestUuidv7(t *testing.T) {
	tc := []pesticide.RegexpTestCase{
		{Name: "TestUuidv7", Template: `{{uuidv7}}`, Regexp: `^[\da-f]{8}-[\da-f]{4}-7[\da-f]{3}-[\da-f]{4}-[\da-f]{12}$`, Length: 36},
	}

	pesticide.RunRegexpTestCases(t, uniqueid.NewRegistry(), tc)
}

// TestUuidv7IsSortable ensures the defining property of the version 7: two
// UUIDs generated in a row are ordered by generation time.
func TestUuidv7IsSortable(t *testing.T) {
	first, err := pesticide.TestTemplate(t, uniqueid.NewRegistry(), `{{ uuidv7 }}`, nil)
	require.NoError(t, err)

	second, err := pesticide.TestTemplate(t, uniqueid.NewRegistry(), `{{ uuidv7 }}`, nil)
	require.NoError(t, err)

	require.Less(t, first, second)
}

// failingReader is a random source always failing, used to check the error of
// the generators is propagated instead of being swallowed.
type failingReader struct{}

func (failingReader) Read(p []byte) (int, error) {
	return 0, errors.New("no entropy available")
}

func TestUuidv7WithoutEntropy(t *testing.T) {
	uuid.SetRand(failingReader{})
	t.Cleanup(func() { uuid.SetRand(nil) })

	tc := []pesticide.TestCase{
		{Name: "TestUuidv7WithoutEntropy", Input: `{{ uuidv7 }}`, ExpectedErr: "no entropy available"},
	}

	pesticide.RunTestCases(t, uniqueid.NewRegistry(), tc)
}

func TestUuidv5(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestPredefinedNamespace", Input: `{{ "python.org" | uuidv5 "dns" }}`, ExpectedOutput: uuidv5Fixture},
		{Name: "TestPredefinedNamespaceUppercase", Input: `{{ "python.org" | uuidv5 "DNS" }}`, ExpectedOutput: uuidv5Fixture},
		{Name: "TestUrlNamespace", Input: `{{ "python.org" | uuidv5 "url" }}`, ExpectedOutput: "7af94e2b-4dd9-50f0-9c9a-8a48519bdef0"},
		{Name: "TestOidNamespace", Input: `{{ "python.org" | uuidv5 "oid" }}`, ExpectedOutput: "cd5d0bff-2444-5d26-ab53-4f7db1cb733d"},
		{Name: "TestX500Namespace", Input: `{{ "python.org" | uuidv5 "x500" }}`, ExpectedOutput: "e9246a06-296f-5b50-be57-0519806c97e8"},
		{Name: "TestCustomNamespace", Input: `{{ "python.org" | uuidv5 "6ba7b810-9dad-11d1-80b4-00c04fd430c8" }}`, ExpectedOutput: uuidv5Fixture},
		{Name: "TestEmptyName", Input: `{{ "" | uuidv5 "dns" }}`, ExpectedOutput: "4ebd0208-8328-5d69-8c44-ec50939c0967"},
		{Name: "TestIsDeterministic", Input: `{{ eq ("python.org" | uuidv5 "dns") ("python.org" | uuidv5 "dns") }}`, ExpectedOutput: "true"},
		{Name: "TestNamesAreDistinct", Input: `{{ ne ("python.org" | uuidv5 "dns") ("golang.org" | uuidv5 "dns") }}`, ExpectedOutput: "true"},
		{Name: "TestNamespacesAreDistinct", Input: `{{ ne ("python.org" | uuidv5 "dns") ("python.org" | uuidv5 "url") }}`, ExpectedOutput: "true"},
		{Name: "TestWithInvalidNamespace", Input: `{{ "python.org" | uuidv5 "invalid" }}`, ExpectedErr: `invalid namespace "invalid"`},
		{Name: "TestWithEmptyNamespace", Input: `{{ "python.org" | uuidv5 "" }}`, ExpectedErr: `invalid namespace ""`},
	}

	pesticide.RunTestCases(t, uniqueid.NewRegistry(), tc)
}

func TestUuidv3(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestPredefinedNamespace", Input: `{{ "python.org" | uuidv3 "dns" }}`, ExpectedOutput: "6fa459ea-ee8a-3ca4-894e-db77e160355e"},
		{Name: "TestCustomNamespace", Input: `{{ "python.org" | uuidv3 "6ba7b810-9dad-11d1-80b4-00c04fd430c8" }}`, ExpectedOutput: "6fa459ea-ee8a-3ca4-894e-db77e160355e"},
		{Name: "TestEmptyName", Input: `{{ "" | uuidv3 "dns" }}`, ExpectedOutput: "c87ee674-4ddc-3efe-a74e-dfe25da5d7b3"},
		{Name: "TestIsDeterministic", Input: `{{ eq ("python.org" | uuidv3 "dns") ("python.org" | uuidv3 "dns") }}`, ExpectedOutput: "true"},
		{Name: "TestDiffersFromVersion5", Input: `{{ ne ("python.org" | uuidv3 "dns") ("python.org" | uuidv5 "dns") }}`, ExpectedOutput: "true"},
		{Name: "TestWithInvalidNamespace", Input: `{{ "python.org" | uuidv3 "invalid" }}`, ExpectedErr: `invalid namespace "invalid"`},
		{Name: "TestWithEmptyNamespace", Input: `{{ "python.org" | uuidv3 "" }}`, ExpectedErr: `invalid namespace ""`},
	}

	pesticide.RunTestCases(t, uniqueid.NewRegistry(), tc)
}

func TestUuidNil(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestUuidNil", Input: `{{ uuidNil }}`, ExpectedOutput: "00000000-0000-0000-0000-000000000000"},
	}

	pesticide.RunTestCases(t, uniqueid.NewRegistry(), tc)
}

func TestIsUUID(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestValidUUID", Input: `{{ "886313e1-3b8a-5372-9b90-0c9aee199e5d" | isUUID }}`, ExpectedOutput: "true"},
		{Name: "TestValidNilUUID", Input: `{{ uuidNil | isUUID }}`, ExpectedOutput: "true"},
		{Name: "TestValidGeneratedUUID", Input: `{{ uuidv4 | isUUID }}`, ExpectedOutput: "true"},
		{Name: "TestUppercaseUUID", Input: `{{ "886313E1-3B8A-5372-9B90-0C9AEE199E5D" | isUUID }}`, ExpectedOutput: "true"},
		{Name: "TestUrnPrefixedUUID", Input: `{{ "urn:uuid:886313e1-3b8a-5372-9b90-0c9aee199e5d" | isUUID }}`, ExpectedOutput: "true"},
		{Name: "TestBracedUUID", Input: `{{ "{886313e1-3b8a-5372-9b90-0c9aee199e5d}" | isUUID }}`, ExpectedOutput: "true"},
		{Name: "TestUUIDWithoutHyphen", Input: `{{ "886313e13b8a53729b900c9aee199e5d" | isUUID }}`, ExpectedOutput: "true"},
		{Name: "TestInvalidUUID", Input: `{{ "not-a-uuid" | isUUID }}`, ExpectedOutput: "false"},
		{Name: "TestEmptyString", Input: `{{ "" | isUUID }}`, ExpectedOutput: "false"},
		{Name: "TestTruncatedUUID", Input: `{{ "886313e1-3b8a-5372-9b90" | isUUID }}`, ExpectedOutput: "false"},
		{Name: "TestInvalidUrnPrefix", Input: `{{ "urn:uuiE:886313e1-3b8a-5372-9b90-0c9aee199e5d" | isUUID }}`, ExpectedOutput: "false"},
		{Name: "TestUnbalancedBraces", Input: `{{ "{886313e1-3b8a-5372-9b90-0c9aee199e5d[" | isUUID }}`, ExpectedOutput: "false"},
		{Name: "TestNonHexCharacters", Input: `{{ "zzzzzzzz-3b8a-5372-9b90-0c9aee199e5d" | isUUID }}`, ExpectedOutput: "false"},
	}

	pesticide.RunTestCases(t, uniqueid.NewRegistry(), tc)
}

func TestUuidVersion(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestVersion5", Input: `{{ .V | uuidVersion }}`, ExpectedOutput: "5", Data: map[string]any{"V": uuidv5Fixture}},
		{Name: "TestVersion7", Input: `{{ .V | uuidVersion }}`, ExpectedOutput: "7", Data: map[string]any{"V": uuidv7Fixture}},
		{Name: "TestVersion1", Input: `{{ .V | uuidVersion }}`, ExpectedOutput: "1", Data: map[string]any{"V": uuidv1Fixture}},
		{Name: "TestVersion4", Input: `{{ uuidv4 | uuidVersion }}`, ExpectedOutput: "4"},
		{Name: "TestUrnPrefixedUUID", Input: `{{ "urn:uuid:886313e1-3b8a-5372-9b90-0c9aee199e5d" | uuidVersion }}`, ExpectedOutput: "5"},
		{Name: "TestNilUUID", Input: `{{ uuidNil | uuidVersion }}`, ExpectedOutput: "0"},
		{Name: "TestEmptyString", Input: `{{ "" | uuidVersion }}`, ExpectedErr: "invalid UUID"},
		{Name: "TestWithInvalidInput", Input: `{{ "not-a-uuid" | uuidVersion }}`, ExpectedErr: "invalid UUID"},
	}

	pesticide.RunTestCases(t, uniqueid.NewRegistry(), tc)
}

func TestUuidTime(t *testing.T) {
	// temporarily force time.Local to UTC to keep the output deterministic
	pesticide.ForceTimeLocal(t, time.UTC)

	tc := []pesticide.TestCase{
		{Name: "TestVersion7", Input: `{{ .V | uuidTime }}`, ExpectedOutput: "2024-05-07 15:04:05 +0000 UTC", Data: map[string]any{"V": uuidv7Fixture}},
		{Name: "TestVersion1", Input: `{{ .V | uuidTime }}`, ExpectedOutput: "2024-05-07 15:04:05 +0000 UTC", Data: map[string]any{"V": uuidv1Fixture}},
		{Name: "TestVersion4HasNoTime", Input: `{{ uuidv4 | uuidTime }}`, ExpectedErr: "uuid version 4 does not embed a time"},
		{Name: "TestVersion5HasNoTime", Input: `{{ .V | uuidTime }}`, ExpectedErr: "uuid version 5 does not embed a time", Data: map[string]any{"V": uuidv5Fixture}},
		{Name: "TestNilUUIDHasNoTime", Input: `{{ uuidNil | uuidTime }}`, ExpectedErr: "uuid version 0 does not embed a time"},
		{Name: "TestWithInvalidInput", Input: `{{ "not-a-uuid" | uuidTime }}`, ExpectedErr: "invalid UUID"},
	}

	pesticide.RunTestCases(t, uniqueid.NewRegistry(), tc)
}

// TestUuidTimeOnVersion6 covers the version 6, which cannot be pinned to a
// fixture: the version bits overwrite the low bits of its timestamp, so a
// hand-written fixture would assert that quirk instead of our own code.
func TestUuidTimeOnVersion6(t *testing.T) {
	id, err := uuid.NewV6()
	require.NoError(t, err)

	got, err := uniqueid.NewRegistry().UuidTime(id.String())
	require.NoError(t, err)
	require.WithinDuration(t, time.Now(), got, time.Minute)
}
