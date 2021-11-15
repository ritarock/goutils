package date

import (
	gs "ritarock/goparts/internal/string"
	"time"
)

const (
	LAYOUT_YYYYMMDD   = "20060102"
	LAYOUT_HHMM       = "1504"
	LAYOUT_YYYYMMDDHH = "20060102150405"
)

var timeNow = time.Now()

func ToYYYYMMDD(t time.Time) string {
	return t.Format(LAYOUT_YYYYMMDD)
}

func ToYYYY_MM_DD(t time.Time) string {
	return gs.Combine(
		t.Format(LAYOUT_YYYYMMDDHH)[0:4],
		"-",
		t.Format(LAYOUT_YYYYMMDDHH)[4:6],
		"-",
		t.Format(LAYOUT_YYYYMMDDHH)[6:8],
	)
}

func ToYYYY_MM_DD_HH_MM_SS(t time.Time) string {
	return gs.Combine(
		t.Format(LAYOUT_YYYYMMDDHH)[0:4],
		"-",
		t.Format(LAYOUT_YYYYMMDDHH)[4:6],
		"-",
		t.Format(LAYOUT_YYYYMMDDHH)[6:8],
		"-",
		t.Format(LAYOUT_YYYYMMDDHH)[8:10],
		"-",
		t.Format(LAYOUT_YYYYMMDDHH)[10:12],
		"-",
		t.Format(LAYOUT_YYYYMMDDHH)[12:14],
	)
}

func ToHHMM(t time.Time) string {
	return t.Format(LAYOUT_HHMM)
}

func TimeSpan(from, to time.Time) time.Duration {
	return to.Sub(from)
}

func GetBeginningOfMonth() time.Time {
	now := timeNow
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
}

func GetEndMonth() time.Time {
	now := timeNow
	t := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local)
	return t.Add(-24 * time.Hour)
}
