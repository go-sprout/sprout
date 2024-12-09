package time_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/go-sprout/sprout/pesticide"
	rtime "github.com/go-sprout/sprout/registry/time"
)

func TestDate(t *testing.T) {
	t.Run("UTC", func(t *testing.T) {
		timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)
		tc := []pesticide.TestCase{
			{Name: "TestTimeObject", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": timeTest}},
			{Name: "TestTimeObjectPointer", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": &timeTest}},
		}

		pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
	})

	t.Run("New York timezone", func(t *testing.T) {
		local, err := time.LoadLocation("America/New_York")
		require.NoError(t, err)

		timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, local)

		tc := []pesticide.TestCase{
			{Name: "TestTimeObject", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 -0400", Data: map[string]any{"V": timeTest}},
			{Name: "TestTimeObjectPointer", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 -0400", Data: map[string]any{"V": &timeTest}},
		}

		pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
	})

	t.Run("New York offset", func(t *testing.T) {
		timeTest, err := time.Parse("02 Jan 06 15:04 -0700", "07 May 24 15:04 -0400")
		require.NoError(t, err)

		tc := []pesticide.TestCase{
			{Name: "TestTimeObject", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 -0400", Data: map[string]any{"V": timeTest}},
			{Name: "TestTimeObjectPointer", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 -0400", Data: map[string]any{"V": &timeTest}},
		}

		pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
	})

	t.Run("New York timezone", func(t *testing.T) {
		local, err := time.LoadLocation("America/New_York")
		require.NoError(t, err)

		timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, local)

		tc := []pesticide.TestCase{
			{Name: "TestTimeObject", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 -0400", Data: map[string]any{"V": timeTest}},
			{Name: "TestTimeObjectPointer", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 -0400", Data: map[string]any{"V": &timeTest}},
		}

		pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
	})

	t.Run("unixtime", func(t *testing.T) {
		t.Run("UTC", func(t *testing.T) {
			// temporarily force time.Local to UTC
			rtime.ForceTimeLocal(t, time.UTC)

			// here we are simulating a [gotime.Now]
			timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

			tc := []pesticide.TestCase{
				{Name: "TestTimeObjectUnix", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": timeTest.Unix()}},
				{Name: "TestTimeObjectUnixInt", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": int(timeTest.Unix())}},
			}

			pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
		})

		t.Run("New York timezone", func(t *testing.T) {
			local, err := time.LoadLocation("America/New_York")
			require.NoError(t, err)

			// temporarily force time.Local to New York
			rtime.ForceTimeLocal(t, local)

			// here we are simulating a [gotime.Now] call
			timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, local)

			tc := []pesticide.TestCase{
				{Name: "TestTimeObject", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 -0400", Data: map[string]any{"V": timeTest}},
				{Name: "TestTimeObjectPointer", Input: `{{ .V | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:04 -0400", Data: map[string]any{"V": &timeTest}},
			}

			pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
		})
	})
}

func TestDateInZone(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	tc := []pesticide.TestCase{
		{Name: "TestTimeObject", Input: `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": timeTest}},
		{Name: "TestTimeObjectPointer", Input: `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": &timeTest}},
		{Name: "TestTimeObjectUnix", Input: `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": timeTest.Unix()}},
		{Name: "TestTimeObjectUnixInt", Input: `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": int(timeTest.Unix())}},
		{Name: "TestTimeObjectUnixInt", Input: `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, ExpectedOutput: "07 May 24 15:04 +0000", Data: map[string]any{"V": int32(timeTest.Unix())}},
		{Name: "TestWithInvalidInput", Input: `{{ dateInZone "02 Jan 06 15:04 -0700" .V "UTC" }}`, ExpectedOutput: time.Now().UTC().Format("02 Jan 06 15:04 -0700"), Data: map[string]any{"V": "invalid"}},
		{Name: "TestWithInvalidZone", Input: `{{ dateInZone "02 Jan 06 15:04 -0700" .V "invalid" }}`, ExpectedErr: "unknown time zone invalid", Data: map[string]any{"V": timeTest}},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}

func TestDuration(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "InvalidInput", Input: `{{ .V | duration }}`, ExpectedOutput: "0s", Data: map[string]any{"V": "1h"}},
		{Name: "TestDurationWithInt", Input: `{{ .V | duration }}`, ExpectedOutput: "10s", Data: map[string]any{"V": int(10)}},
		{Name: "TestDurationWithInt64", Input: `{{ .V | duration }}`, ExpectedOutput: "10s", Data: map[string]any{"V": int64(10)}},
		{Name: "TestDurationWithFloat32", Input: `{{ .V | duration }}`, ExpectedOutput: "10s", Data: map[string]any{"V": float32(10)}},
		{Name: "TestDurationWithFloat64", Input: `{{ .V | duration }}`, ExpectedOutput: "10s", Data: map[string]any{"V": float64(10)}},
		{Name: "TestDurationWithString", Input: `{{ .V | duration }}`, ExpectedOutput: "26h3m4s", Data: map[string]any{"V": "93784"}},
		{Name: "TestDurationWithInvalidType", Input: `{{ .V | duration }}`, ExpectedOutput: "0s", Data: map[string]any{"V": make(chan int)}},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}

func TestDateAgo(t *testing.T) {
	timeTest := time.Now().Add(-time.Hour * 24)

	tc := []pesticide.TestCase{
		{Name: "TestTimeObject", Input: `{{ .V | dateAgo | substr 0 5 }}`, ExpectedOutput: "24h0m", Data: map[string]any{"V": timeTest}},
		{Name: "TestTimeObjectPointer", Input: `{{ .V | dateAgo | substr 0 5 }}`, ExpectedOutput: "24h0m", Data: map[string]any{"V": &timeTest}},
		{Name: "TestTimeObjectUnix", Input: `{{ .V | dateAgo | substr 0 5 }}`, ExpectedOutput: "24h0m", Data: map[string]any{"V": timeTest.Unix()}},
		{Name: "TestTimeObjectUnixInt", Input: `{{ .V | dateAgo | substr 0 5 }}`, ExpectedOutput: "24h0m", Data: map[string]any{"V": int(timeTest.Unix())}},
		{Name: "TestTimeObjectUnixInt32", Input: `{{ .V | dateAgo | substr 0 5 }}`, ExpectedOutput: "24h0m", Data: map[string]any{"V": int32(timeTest.Unix())}},
		{Name: "TestWithInvalidInput", Input: `{{ .V | dateAgo }}`, ExpectedOutput: "0s", Data: map[string]any{"V": "invalid"}},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}

func TestNow(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestNow", Input: `{{ now | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: time.Now().Format("02 Jan 06 15:04 -0700")},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}

func TestUnixEpoch(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	tc := []pesticide.TestCase{
		{Name: "TestUnixEpoch", Input: `{{ .V | unixEpoch }}`, ExpectedOutput: "1715094245", Data: map[string]any{"V": timeTest}},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}

func TestDateModify(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	tc := []pesticide.TestCase{
		{Name: "AddOneHour", Input: `{{ .V | mustDateModify "1h" | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 16:04 +0000", Data: map[string]any{"V": timeTest}},
		{Name: "AddOneHourWithPlusSign", Input: `{{ .V | mustDateModify "+1h" | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 16:04 +0000", Data: map[string]any{"V": timeTest}},
		{Name: "SubtractOneHour", Input: `{{ .V | mustDateModify "-1h" | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 14:04 +0000", Data: map[string]any{"V": timeTest}},
		{Name: "AddTenMinutes", Input: `{{ .V | mustDateModify "10m" | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:14 +0000", Data: map[string]any{"V": timeTest}},
		{Name: "SubtractTenSeconds", Input: `{{ .V | mustDateModify "-10s" | date "02 Jan 06 15:04 -0700" }}`, ExpectedOutput: "07 May 24 15:03 +0000", Data: map[string]any{"V": timeTest}},
		{Name: "WithEmptyInput", Input: `{{ .V | mustDateModify "" }}`, Data: map[string]any{"V": timeTest}, ExpectedErr: "invalid duration"},
		{Name: "WithInvalidInput", Input: `{{ .V | mustDateModify "zz" }}`, Data: map[string]any{"V": timeTest}, ExpectedErr: "invalid duration"},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}

func TestDurationRound(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "EmptyInput", Input: `{{ .V | durationRound }}`, ExpectedOutput: "0s", Data: map[string]any{"V": ""}},
		{Name: "RoundToHour", Input: `{{ .V | durationRound }}`, ExpectedOutput: "2h", Data: map[string]any{"V": "2h5s"}},
		{Name: "RoundToDay", Input: `{{ .V | durationRound }}`, ExpectedOutput: "1d", Data: map[string]any{"V": "24h5s"}},
		{Name: "RoundToMonth", Input: `{{ .V | durationRound }}`, ExpectedOutput: "3mo", Data: map[string]any{"V": "2400h5s"}},
		{Name: "RoundToMinute", Input: `{{ .V | durationRound }}`, ExpectedOutput: "45m", Data: map[string]any{"V": int64(45*time.Minute + 30*time.Second)}},
		{Name: "RoundToSecond", Input: `{{ .V | durationRound }}`, ExpectedOutput: "1s", Data: map[string]any{"V": int64(1*time.Second + 500*time.Millisecond)}},
		{Name: "RoundaDuration", Input: `{{ .V | durationRound }}`, ExpectedOutput: "2s", Data: map[string]any{"V": 2 * time.Second}},
		{Name: "RoundToYear", Input: `{{ .V | durationRound }}`, ExpectedOutput: "1y", Data: map[string]any{"V": int64(365*24*time.Hour + 12*time.Hour)}},
		{Name: "RoundToYearNegative", Input: `{{ .V | durationRound }}`, ExpectedOutput: "1y", Data: map[string]any{"V": time.Now().Add(-365*24*time.Hour - 72*time.Hour)}},
		{Name: "InvalidInput", Input: `{{ .V | durationRound }}`, ExpectedOutput: "0s", Data: map[string]any{"V": make(chan int)}},
		{Name: "RoundToHourNegative", Input: `{{ .V | durationRound }}`, ExpectedOutput: "-1h", Data: map[string]any{"V": "-1h01s"}},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}

func TestHtmlDate(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	tc := []pesticide.TestCase{
		{Name: "TestTimeObject", Input: `{{ .V | htmlDate }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": timeTest}},
		{Name: "TestTimeObjectPointer", Input: `{{ .V | htmlDate }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": &timeTest}},
		{Name: "TestTimeObjectUnix", Input: `{{ .V | htmlDate }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": timeTest.Unix()}},
		{Name: "TestTimeObjectUnixInt", Input: `{{ .V | htmlDate }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": int(timeTest.Unix())}},
		{Name: "TestTimeObjectUnixInt32", Input: `{{ .V | htmlDate }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": int32(timeTest.Unix())}},
		{Name: "TestZeroValue", Input: `{{ .V | htmlDate }}`, ExpectedOutput: "1970-01-01", Data: map[string]any{"V": 0}},
		{Name: "TestWithInvalidInput", Input: `{{ .V | htmlDate }}`, ExpectedOutput: time.Now().Format("2006-01-02"), Data: map[string]any{"V": make(chan int)}},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}

func TestHtmlDateInZone(t *testing.T) {
	timeTest := time.Date(2024, 5, 7, 15, 4, 5, 0, time.UTC)

	tc := []pesticide.TestCase{
		{Name: "TestTimeObject", Input: `{{ htmlDateInZone .V "UTC" }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": timeTest}},
		{Name: "TestTimeObjectPointer", Input: `{{ htmlDateInZone .V "UTC" }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": &timeTest}},
		{Name: "TestTimeObjectUnix", Input: `{{ htmlDateInZone .V "UTC" }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": timeTest.Unix()}},
		{Name: "TestTimeObjectUnixInt", Input: `{{ htmlDateInZone .V "UTC" }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": int(timeTest.Unix())}},
		{Name: "TestTimeObjectUnixInt32", Input: `{{ htmlDateInZone .V "UTC" }}`, ExpectedOutput: "2024-05-07", Data: map[string]any{"V": int32(timeTest.Unix())}},
		{Name: "TestWithInvalidInput", Input: `{{ htmlDateInZone .V "UTC" }}`, ExpectedOutput: time.Now().Format("2006-01-02"), Data: map[string]any{"V": make(chan int)}},
	}

	pesticide.RunTestCases(t, rtime.NewRegistry(), tc)
}
