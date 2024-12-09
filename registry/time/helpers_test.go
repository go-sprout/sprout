package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestComputeTimeFromFormat(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name string
		date any
		want time.Time
	}{
		{
			name: "time.Time",
			date: now,
			want: now,
		},
		{
			name: "*time.Time",
			date: &now,
			want: now,
		},
		{
			name: "int64",
			date: int64(1643723900),
			want: time.Unix(1643723900, 0),
		},
		{
			name: "int",
			date: 1643723900,
			want: time.Unix(int64(1643723900), 0),
		},
		{
			name: "int32",
			date: int32(1643723900),
			want: time.Unix(int64(1643723900), 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := computeTimeFromFormat(tt.date)
			assert.Equal(t, tt.want, got)
		})
	}

	t.Run("invalid format", func(t *testing.T) {
		// computeTimeFromFormat generates the current time if the format is invalid
		got := computeTimeFromFormat("invalid date")

		// so we can only guess the date is close to the current time
		assert.Less(t, time.Since(got), 10*time.Millisecond)
	})
}
