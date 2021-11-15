package string

import "strconv"

func Combine(words ...string) string {
	b := make([]byte, 0, 10)
	for _, word := range words {
		b = append(b, word...)
	}
	return string(b)
}

func Repeat(str string, num int) string {
	b := make([]byte, 0, 10)
	for i := 0; i < num; i++ {
		b = append(b, str...)
	}
	return string(b)
}

func NumberToString(number int) string {
	return strconv.Itoa(number)
}
