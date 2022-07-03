package date

import (
	"testing"
	"time"
)

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertTime(t *testing.T, got, want time.Time) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func setTestingTime() time.Time {
	return time.Date(2022, 1, 10, 0, 0, 0, 0, time.Local)
}

func TestToYYYYMMDD(t *testing.T) {
	dt := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	got := ToYYYYMMDD(dt)
	want := "20220101"
	assertString(t, got, want)
}

func TestToYYYY_MM_DD(t *testing.T) {
	dt := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	got := ToYYYY_MM_DD(dt)
	want := "2022-01-01"

	assertString(t, got, want)
}

func TestTimeSpan(t *testing.T) {
	from := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2022, 1, 1, 9, 0, 0, 0, time.Local)
	got := TimeSpan(from, to)
	want := 9 * time.Duration(3600) * time.Second

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetBeginningOfMonth(t *testing.T) {
	timeNow = setTestingTime()
	got := GetBeginningOfMonth()
	want := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)

	assertTime(t, got, want)
}

func TestGetEndMonth(t *testing.T) {
	timeNow = setTestingTime()
	got := GetEndMonth()
	want := time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local)

	assertTime(t, got, want)
}
