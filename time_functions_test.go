package sprout

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	var tests = testCases{
		{"TestTimeObject", `{{ .V | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:04 +0000", map[string]any{"V": timeTest}},
		{"TestTimeObjectPointer", `{{ .V | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:04 +0000", map[string]any{"V": &timeTest}},
		{"TestTimeObjectUnix", `{{ .V | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:04 +0000", map[string]any{"V": timeTest.Unix()}},
		{"TestTimeObjectUnixInt", `{{ .V | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:04 +0000", map[string]any{"V": int(timeTest.Unix())}},
	}

	runTestCases(t, tests)
}

func TestDateInZone(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	var tests = testCases{
		{"TestTimeObject", `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, "07 May 24 15:04 +0000", map[string]any{"V": timeTest}},
		{"TestTimeObjectPointer", `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, "07 May 24 15:04 +0000", map[string]any{"V": &timeTest}},
		{"TestTimeObjectUnix", `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, "07 May 24 15:04 +0000", map[string]any{"V": timeTest.Unix()}},
		{"TestTimeObjectUnixInt", `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, "07 May 24 15:04 +0000", map[string]any{"V": int(timeTest.Unix())}},
		{"TestTimeObjectUnixInt", `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, "07 May 24 15:04 +0000", map[string]any{"V": int32(timeTest.Unix())}},
		{"TestWithInvalidInput", `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, time.Now().Format("02 Jan 06 15:04 -0700"), map[string]any{"V": "invalid"}},
		{"TestWithInvalidZone", `{{ dateInZone "02 Jan 06 15:04 -0700" .V "invalid" }}`, "07 May 24 15:04 +0000", map[string]any{"V": timeTest}},
	}

	runTestCases(t, tests)
}

func TestDuration(t *testing.T) {
	var tests = testCases{
		{"InvalidInput", `{{ .V | duration }}`, "0s", map[string]any{"V": "1h"}},
		{"TestDurationWithInt64", `{{ .V | duration }}`, "10s", map[string]any{"V": int64(10)}},
		{"TestDurationWithString", `{{ .V | duration }}`, "26h3m4s", map[string]any{"V": "93784"}},
		{"TestDurationWithInvalidType", `{{ .V | duration }}`, "0s", map[string]any{"V": make(chan int)}},
	}

	runTestCases(t, tests)
}

func TestDateAgo(t *testing.T) {
	timeTest := time.Now().Add(-time.Hour * 24)

	var tests = testCases{
		{"TestTimeObject", `{{ .V | dateAgo | substr 0 5 }}`, "24h0m", map[string]any{"V": timeTest}},
		{"TestTimeObjectPointer", `{{ .V | dateAgo | substr 0 5 }}`, "24h0m", map[string]any{"V": &timeTest}},
		{"TestTimeObjectUnix", `{{ .V | dateAgo | substr 0 5 }}`, "24h0m", map[string]any{"V": timeTest.Unix()}},
		{"TestTimeObjectUnixInt", `{{ .V | dateAgo | substr 0 5 }}`, "24h0m", map[string]any{"V": int(timeTest.Unix())}},
		{"TestTimeObjectUnixInt32", `{{ .V | dateAgo | substr 0 5 }}`, "24h0m", map[string]any{"V": int32(timeTest.Unix())}},
		{"TestWithInvalidInput", `{{ .V | dateAgo }}`, "0s", map[string]any{"V": "invalid"}},
	}

	runTestCases(t, tests)
}

func TestNow(t *testing.T) {
	var tests = testCases{
		{"TestNow", `{{ now | date "02 Jan 06 15:04 -0700" }}`, time.Now().Format("02 Jan 06 15:04 -0700"), nil},
	}

	runTestCases(t, tests)
}

func TestUnixEpoch(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	var tests = testCases{
		{"TestUnixEpoch", `{{ .V | unixEpoch }}`, "1715094245", map[string]any{"V": timeTest}},
	}

	runTestCases(t, tests)
}

func TestDateModify(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	var tests = testCases{
		{"AddOneHour", `{{ .V | dateModify "1h" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 16:04 +0000", map[string]any{"V": timeTest}},
		{"AddOneHourWithPlusSign", `{{ .V | dateModify "+1h" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 16:04 +0000", map[string]any{"V": timeTest}},
		{"SubtractOneHour", `{{ .V | dateModify "-1h" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 14:04 +0000", map[string]any{"V": timeTest}},
		{"AddTenMinutes", `{{ .V | dateModify "10m" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:14 +0000", map[string]any{"V": timeTest}},
		{"SubtractTenSeconds", `{{ .V | dateModify "-10s" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:03 +0000", map[string]any{"V": timeTest}},
		{"WithInvalidInput", `{{ .V | dateModify "zz" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:04 +0000", map[string]any{"V": timeTest}},
	}

	runTestCases(t, tests)
}

func TestDurationRound(t *testing.T) {
	var tests = testCases{
		{"", `{{ .V | durationRound }}`, "0s", map[string]any{"V": ""}},
		{"", `{{ .V | durationRound }}`, "2h", map[string]any{"V": "2h5s"}},
		{"", `{{ .V | durationRound }}`, "1d", map[string]any{"V": "24h5s"}},
		{"", `{{ .V | durationRound }}`, "3mo", map[string]any{"V": "2400h5s"}},
		{"", `{{ .V | durationRound }}`, "45m", map[string]any{"V": int64(45*time.Minute + 30*time.Second)}},
		{"", `{{ .V | durationRound }}`, "1s", map[string]any{"V": int64(1*time.Second + 500*time.Millisecond)}},
		{"", `{{ .V | durationRound }}`, "1y", map[string]any{"V": int64(365*24*time.Hour + 12*time.Hour)}},
		{"", `{{ .V | durationRound }}`, "1y", map[string]any{"V": time.Now().Add(-365*24*time.Hour - 72*time.Hour)}},
		{"", `{{ .V | durationRound }}`, "0s", map[string]any{"V": make(chan int)}},
		{"", `{{ .V | durationRound }}`, "-1h", map[string]any{"V": "-1h01s"}},
	}

	runTestCases(t, tests)
}

func TestHtmlDate(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	var tests = testCases{
		{"TestTimeObject", `{{ .V | htmlDate }}`, "2024-05-07", map[string]any{"V": timeTest}},
		{"TestTimeObjectPointer", `{{ .V | htmlDate }}`, "2024-05-07", map[string]any{"V": &timeTest}},
		{"TestTimeObjectUnix", `{{ .V | htmlDate }}`, "2024-05-07", map[string]any{"V": timeTest.Unix()}},
		{"TestTimeObjectUnixInt", `{{ .V | htmlDate }}`, "2024-05-07", map[string]any{"V": int(timeTest.Unix())}},
		{"TestTimeObjectUnixInt32", `{{ .V | htmlDate }}`, "2024-05-07", map[string]any{"V": int32(timeTest.Unix())}},
		{"TestZeroValue", `{{ .V | htmlDate }}`, "1970-01-01", map[string]any{"V": 0}},
		{"TestWithInvalidInput", `{{ .V | htmlDate }}`, time.Now().Format("2006-01-02"), map[string]any{"V": make(chan int)}},
	}

	runTestCases(t, tests)
}

func TestHtmlDateInZone(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	var tests = testCases{
		{"TestTimeObject", `{{ htmlDateInZone .V "UTC" }}`, "2024-05-07", map[string]any{"V": timeTest}},
		{"TestTimeObjectPointer", `{{ htmlDateInZone .V "UTC" }}`, "2024-05-07", map[string]any{"V": &timeTest}},
		{"TestTimeObjectUnix", `{{ htmlDateInZone .V "UTC" }}`, "2024-05-07", map[string]any{"V": timeTest.Unix()}},
		{"TestTimeObjectUnixInt", `{{ htmlDateInZone .V "UTC" }}`, "2024-05-07", map[string]any{"V": int(timeTest.Unix())}},
		{"TestTimeObjectUnixInt32", `{{ htmlDateInZone .V "UTC" }}`, "2024-05-07", map[string]any{"V": int32(timeTest.Unix())}},
		{"TestWithInvalidInput", `{{ htmlDateInZone .V "UTC" }}`, time.Now().Format("2006-01-02"), map[string]any{"V": make(chan int)}},
	}

	runTestCases(t, tests)
}

func TestMustDateModify(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	var tests = testCases{
		{"AddOneHour", `{{ .V | mustDateModify "1h" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 16:04 +0000", map[string]any{"V": timeTest}},
		{"AddOneHourWithPlusSign", `{{ .V | mustDateModify "+1h" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 16:04 +0000", map[string]any{"V": timeTest}},
		{"SubtractOneHour", `{{ .V | mustDateModify "-1h" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 14:04 +0000", map[string]any{"V": timeTest}},
		{"AddTenMinutes", `{{ .V | mustDateModify "10m" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:14 +0000", map[string]any{"V": timeTest}},
		{"SubtractTenSeconds", `{{ .V | mustDateModify "-10s" | date "02 Jan 06 15:04 -0700" }}`, "07 May 24 15:03 +0000", map[string]any{"V": timeTest}},
	}

	runTestCases(t, tests)

	var mustTests = mustTestCases{
		{testCase{"WithEmptyInput", `{{ .V | mustDateModify "" }}`, "", map[string]any{"V": timeTest}}, "invalid duration"},
		{testCase{"WithInvalidInput", `{{ .V | mustDateModify "zz" }}`, "", map[string]any{"V": timeTest}}, "invalid duration"},
	}

	runMustTestCases(t, mustTests)
}
