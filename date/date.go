package date

import "time"

const (
	FORMAT_YYYY = "2006"
	FORMAT_MM   = "01"
	FORMAT_DD   = "02"
)

var timeNow = time.Now()

func ToYYYYMMDD(t time.Time) string {
	format := FORMAT_YYYY + FORMAT_MM + FORMAT_DD
	return t.Format(format)
}

func ToYYYY_MM_DD(t time.Time) string {
	format := FORMAT_YYYY + "-" + FORMAT_MM + "-" + FORMAT_DD
	return t.Format(format)
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
