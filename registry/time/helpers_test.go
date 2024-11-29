package time_test

import (
	"testing"
	goTime "time"

	"github.com/stretchr/testify/assert"

	"github.com/go-sprout/sprout/registry/time"
)

func TestComputeTimeFromFormat(t *testing.T) {
	now := goTime.Now()

	tests := []struct {
		name string
		date any
		want goTime.Time
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
			want: goTime.Unix(1643723900, 0),
		},
		{
			name: "int",
			date: 1643723900,
			want: goTime.Unix(int64(1643723900), 0),
		},
		{
			name: "int32",
			date: int32(1643723900),
			want: goTime.Unix(int64(1643723900), 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := time.ComputeTimeFromFormat(tt.date)
			assert.Equal(t, tt.want, got)
		})
	}

	t.Run("invalid format", func(t *testing.T) {
		// ComputeTimeFromFormat generates the current time if the format is invalid
		got := time.ComputeTimeFromFormat("invalid date")

		// so we can only guess the date is close to the current time
		assert.Less(t, goTime.Since(got), 10*goTime.Millisecond)
	})
}
