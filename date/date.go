package date

import "time"

const (
	LAYOUT_YYYY = "2006"
	LAYOUT_MM   = "01"
	LAYOUT_DD   = "02"
)

var timeNow = time.Now()

func ToYYYYMMDD(t time.Time) string {
	layout := LAYOUT_YYYY + LAYOUT_MM + LAYOUT_DD
	return t.Format(layout)
}

func ToYYYY_MM_DD(t time.Time) string {
	layout := LAYOUT_YYYY + "-" + LAYOUT_MM + "-" + LAYOUT_DD
	return t.Format(layout)
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
	return t.AddDate(0, 0, -1)
}
