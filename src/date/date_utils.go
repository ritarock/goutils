package date

import "time"

const LAYOUT_YYYYMMDD = "20060102"
const LAYOUT_HHMM = "1504"

func ToYYYYMMDD(t time.Time) string {
	return t.Format(LAYOUT_YYYYMMDD)
}

func ToHHMM(t time.Time) string {
	return t.Format(LAYOUT_HHMM)
}

func NowYYYYMMDD(t time.Time) string {
	return ToYYYYMMDD(t)
}

func TimeSpan(from, to time.Time) time.Duration {
	return to.Sub(from)
}
