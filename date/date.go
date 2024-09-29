package date

import (
	"time"
)

func ToYYYYMMDD(t time.Time) string {
	return t.Format("20060102")
}

func ToYYYY_MM_DD(t time.Time) string {
	return t.Format(time.DateOnly)
}

func TimeSpan(from, to time.Time) time.Duration {
	span := to.Sub(from)
	if span <= -time.Second {
		return 0
	}

	return span
}

func GetBeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

func GetEndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 0, -1)
}

func GetDayOfNextMonth(t time.Time) time.Time {
	y1, m1, day := t.Date()
	first := time.Date(y1, m1, 1, 0, 0, 0, 0, time.Local)

	y2, m2, _ := first.AddDate(0, 1, 0).Date()
	next := time.Date(y2, m2, day, 0, 0, 0, 0, time.Local)
	if m2 != next.Month() {
		return first.AddDate(0, 2, -1)
	}
	return next
}
