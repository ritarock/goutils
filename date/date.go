package date

import "time"

const (
	LAYOUT_YYYYMMDD = "20060102"
)

var timeNow = time.Now()

func ToYYYYMMDD(t time.Time) string {
	return t.Format(LAYOUT_YYYYMMDD)
}

func ToYYYY_MM_DD(t time.Time) string {
	dt := t.Format(LAYOUT_YYYYMMDD)

	return dt[0:4] + "-" + dt[4:6] + "-" + dt[6:8]
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
