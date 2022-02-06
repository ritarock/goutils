package date

import (
	"testing"
	"time"
)

func TestToYYYYMMDD(t *testing.T) {
	got := ToYYYYMMDD(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))
	want := "20210101"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestToYYYY_MM_DD(t *testing.T) {
	got := ToYYYY_MM_DD(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))
	want := "2021-01-01"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTimeSpan(t *testing.T) {
	from := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2021, 1, 1, 1, 0, 0, 0, time.Local)
	got := TimeSpan(from, to)
	want := time.Duration(3600) * time.Second

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetBeginningOfMonth(t *testing.T) {
	now := time.Date(2021, 1, 10, 0, 0, 0, 0, time.Local)
	timeNow = func() time.Time {
		return now
	}()
	got := GetBeginningOfMonth()
	want := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetEndMonth(t *testing.T) {
	now := time.Date(2021, 1, 10, 0, 0, 0, 0, time.Local)
	timeNow = func() time.Time {
		return now
	}()
	got := GetEndMonth()
	want := time.Date(2021, 1, 31, 0, 0, 0, 0, time.Local)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
