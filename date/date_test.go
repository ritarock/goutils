package date

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	timeNow = time.Date(2023, 1, 15, 0, 0, 0, 0, time.Local)
	ret := m.Run()
	os.Exit(ret)
}

func TestToYYYYMMDD(t *testing.T) {
	dt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	got := ToYYYYMMDD(dt)
	want := "20230101"
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestToYYYY_MM_DD(t *testing.T) {
	dt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	got := ToYYYY_MM_DD(dt)
	want := "2023-01-01"
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestTimeSpan(t *testing.T) {
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2023, 1, 1, 9, 0, 0, 0, time.Local)
	got := TimeSpan(from, to)
	want := 9 * time.Duration(3600) * time.Second
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGetBeginningOfMonth(t *testing.T) {
	got := GetBeginningOfMonth()
	want := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGetEndMonth(t *testing.T) {
	got := GetEndMonth()
	want := time.Date(2023, 1, 31, 0, 0, 0, 0, time.Local)
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGetPast(t *testing.T) {
	got := GetPast(10)
	want := time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local)
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGetFuture(t *testing.T) {
	got := GetFuture(10)
	want := time.Date(2023, 1, 25, 0, 0, 0, 0, time.Local)
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
