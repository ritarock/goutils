package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToYYYYMMDD(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  time.Time
		want string
	}{
		{
			name: "succeed",
			arg:  time.Date(2025, 10, 10, 0, 0, 0, 0, time.Local),
			want: "20251010",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := ToYYYYMMDD(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestToYYYY_MM_DD(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  time.Time
		want string
	}{
		{
			name: "succeed",
			arg:  time.Date(2025, 10, 10, 0, 0, 0, 0, time.Local),
			want: "2025-10-10",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := ToYYYY_MM_DD(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestTimeSpan(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		from time.Time
		to   time.Time
		want time.Duration
	}{
		{
			name: "succeed: to < from",
			from: time.Date(2025, 10, 1, 0, 0, 0, 0, time.Local),
			to:   time.Date(2025, 10, 1, 9, 0, 0, 0, time.Local),
			want: 9 * time.Hour,
		},
		{
			name: "succeed: to > from",
			from: time.Date(2025, 10, 1, 0, 1, 0, 0, time.Local),
			to:   time.Date(2025, 10, 1, 0, 0, 0, 0, time.Local),
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := TimeSpan(test.from, test.to)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetMonthStartDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  time.Time
		want time.Time
	}{
		{
			name: "succeed",
			arg:  time.Date(2025, 10, 10, 0, 0, 0, 0, time.Local),
			want: time.Date(2025, 10, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := GetMonthStartDate(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetMonthEndDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  time.Time
		want time.Time
	}{
		{
			name: "succeed",
			arg:  time.Date(2025, 10, 10, 0, 0, 0, 0, time.Local),
			want: time.Date(2025, 10, 31, 0, 0, 0, 0, time.Local),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := GetMonthEndDate(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetNextMonthDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  time.Time
		want time.Time
	}{
		{
			name: "succeed: 30th to 30th of next month",
			arg:  time.Date(2025, 5, 30, 0, 0, 0, 0, time.Local),
			want: time.Date(2025, 6, 30, 0, 0, 0, 0, time.Local),
		},
		{
			name: "succeed: 31st to 30th of next month",
			arg:  time.Date(2025, 5, 31, 0, 0, 0, 0, time.Local),
			want: time.Date(2025, 6, 30, 0, 0, 0, 0, time.Local),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := GetNextMonthDate(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}
