package date

import "time"

var timeNow = time.Now()

func ToYYYYMMDD(t time.Time) string {
	return t.Format("20060102")
}

func ToYYYY_MM_DD(t time.Time) string {
	return t.Format(time.DateOnly)
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

func GetPast(day int) time.Time {
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-day, 0, 0, 0, 0, time.Local)
}

func GetFuture(day int) time.Time {
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()+day, 0, 0, 0, 0, time.Local)
}
