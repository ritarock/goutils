package date

import (
	"fmt"
	"time"
)

var nowTime = time.Now()

func ToYYYYMMDD(t time.Time) string {
	return t.Format(LAYOUT_YYYYMMDD)
}

func ToHHMM(t time.Time) string {
	return t.Format(LAYOUT_HHMM)
}

func TimeSpan(from, to time.Time) time.Duration {
	return to.Sub(from)
}

func ToYYYY_MM_DD(t time.Time) string {
	return fmt.Sprintf("%v-%v-%v",
		t.Format(LAYOUT_YYYYMMDD)[0:4],
		t.Format(LAYOUT_YYYYMMDD)[4:6],
		t.Format(LAYOUT_YYYYMMDD)[6:8])
}

func ToYYYY_MM_DD_HH_MM_SS(t time.Time) string {
	return fmt.Sprintf("%v-%v-%v %v:%v:%v",
		t.Format(LAYOUT_YYYYMMDDHH)[0:4],
		t.Format(LAYOUT_YYYYMMDDHH)[4:6],
		t.Format(LAYOUT_YYYYMMDDHH)[6:8],
		t.Format(LAYOUT_YYYYMMDDHH)[8:10],
		t.Format(LAYOUT_YYYYMMDDHH)[10:12],
		t.Format(LAYOUT_YYYYMMDDHH)[12:14])
}

func GetBeginningOfMonth() string {
	now := nowTime
	t := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	return t.Format(LAYOUT_YYYYMMDD)
}

func GetEndOfMonth() string {
	now := nowTime
	t := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local)
	t = t.Add(-time.Minute)
	return t.Format(LAYOUT_YYYYMMDD)
}
