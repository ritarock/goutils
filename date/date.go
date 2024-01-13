package date

import "time"

func ToYYYYMMDD(t time.Time) string {
	return t.Format("20060102")
}

func ToYYYY_MM_DD(t time.Time) string {
	return t.Format(time.DateOnly)
}

func TimeSpan(from, to time.Time) time.Duration {
	span := to.Sub(from)
	if span < -time.Hour {
		return 0
	}
	return span
}

func GetBeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

func GetEndMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 0, -1)
}
