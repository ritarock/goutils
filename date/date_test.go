package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToYYYYMMDD(t *testing.T) {
	tests := []struct {
		arg  time.Time
		want string
	}{
		{
			arg:  time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
			want: "20241001",
		},
	}

	for _, test := range tests {
		got := ToYYYYMMDD(test.arg)
		assert.Equal(t, test.want, got)
	}
}

func TestToYYYY_MM_DD(t *testing.T) {
	tests := []struct {
		arg  time.Time
		want string
	}{
		{
			arg:  time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
			want: "2024-10-01",
		},
	}

	for _, test := range tests {
		got := ToYYYY_MM_DD(test.arg)
		assert.Equal(t, test.want, got)
	}
}

func TestTimeSpan(t *testing.T) {
	tests := []struct {
		from time.Time
		to   time.Time
		want time.Duration
	}{
		{
			from: time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
			to:   time.Date(2024, 10, 1, 9, 0, 0, 0, time.Local),
			want: 9 * time.Hour,
		},
		{
			from: time.Date(2024, 10, 1, 0, 0, 1, 0, time.Local),
			to:   time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
			want: 0,
		},
	}

	for _, test := range tests {
		got := TimeSpan(test.from, test.to)
		assert.Equal(t, test.want, got)
	}
}

func TestGetBeginningOfMonth(t *testing.T) {
	tests := []struct {
		arg  time.Time
		want time.Time
	}{
		{
			arg:  time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
			want: time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, test := range tests {
		got := GetBeginningOfMonth(test.arg)
		assert.Equal(t, test.want, got)
	}
}

func TestGetEndOfMonth(t *testing.T) {
	tests := []struct {
		arg  time.Time
		want time.Time
	}{
		{
			arg:  time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
			want: time.Date(2024, 10, 31, 0, 0, 0, 0, time.Local),
		},
	}

	for _, test := range tests {
		got := GetEndOfMonth(test.arg)
		assert.Equal(t, test.want, got)
	}
}

func TestGetDayOfNextMonth(t *testing.T) {
	tests := []struct {
		arg  time.Time
		want time.Time
	}{
		{
			arg:  time.Date(2024, 5, 30, 0, 0, 0, 0, time.Local),
			want: time.Date(2024, 6, 30, 0, 0, 0, 0, time.Local),
		},
		{
			arg:  time.Date(2024, 5, 31, 0, 0, 0, 0, time.Local),
			want: time.Date(2024, 6, 30, 0, 0, 0, 0, time.Local),
		},
	}

	for _, test := range tests {
		got := GetDayOfNextMonth(test.arg)
		assert.Equal(t, test.want, got)
	}
}
