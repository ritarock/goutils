package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToYYYYMMDD(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want string
	}{
		{
			name: "success yyyymmdd",
			arg:  time.Date(2023, 8, 1, 0, 0, 0, 0, time.Local),
			want: "20230801",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ToYYYYMMDD(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestToYYYY_MM_DD(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want string
	}{
		{
			name: "success yyyy-mm-dd",
			arg:  time.Date(2023, 8, 1, 0, 0, 0, 0, time.Local),
			want: "2023-08-01",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ToYYYY_MM_DD(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestTimeSpan(t *testing.T) {
	tests := []struct {
		name string
		from time.Time
		to   time.Time
		want time.Duration
	}{
		{
			name: "success timeSpan",
			from: time.Date(2023, 8, 1, 0, 0, 0, 0, time.Local),
			to:   time.Date(2023, 8, 1, 9, 0, 0, 0, time.Local),
			want: 9 * time.Hour,
		},
		{
			name: "failed timeSpan",
			from: time.Date(2023, 8, 1, 9, 0, 0, 0, time.Local),
			to:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.Local),
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TimeSpan(test.from, test.to)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetBeginningOfMonth(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want time.Time
	}{
		{
			name: "success get get beginning of month",
			arg:  time.Date(2023, 8, 15, 0, 0, 0, 0, time.Local),
			want: time.Date(2023, 8, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := GetBeginningOfMonth(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetEndMonth(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want time.Time
	}{
		{
			name: "success get get end month",
			arg:  time.Date(2023, 8, 15, 0, 0, 0, 0, time.Local),
			want: time.Date(2023, 8, 31, 0, 0, 0, 0, time.Local),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := GetEndMonth(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}
