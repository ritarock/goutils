package date

import (
	"fmt"
	"time"
)

const LAYOUT_YYYYMMDD = "20060102"
const LAYOUT_HHMM = "1504"
const LAYOUT_YYYYMMDDHH = "20060102150405"

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
