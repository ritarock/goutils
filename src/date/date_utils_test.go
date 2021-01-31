package date

import (
	"testing"
	"time"
)

func TestToYYYYMMDD(t *testing.T) {
	timeNow := time.Now()
	result := len(ToYYYYMMDD(timeNow))
	expect := 8

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestToHHMM(t *testing.T) {
	timeNow := time.Now()
	result := len(ToHHMM(timeNow))
	expect := 4

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestNowYYYYMMDD(t *testing.T) {
	timeNow := time.Now()
	result := len(NowYYYYMMDD(timeNow))
	expect := 8

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestTimeSpan(t *testing.T) {
	fromTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	toTime := time.Date(2021, 1, 1, 1, 0, 0, 0, time.Local)
	result := TimeSpan(fromTime, toTime)
	expect := time.Duration(3600000000000)

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}
