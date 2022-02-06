package string

import (
	"strconv"
	"strings"
)

func Repeat(str string, num int) string {
	b := make([]byte, 0, 10)
	for i := 0; i < num; i++ {
		b = append(b, str...)
	}

	return string(b)
}

func NumberToString(num int) string {
	return strconv.Itoa(num)
}

func Capitalize(str string) string {
	word := strings.Split(str, "")

	for i, v := range word {
		word[i] = strings.ToLower(v)
	}
	word[0] = strings.ToUpper(word[0])

	return strings.Join(word, "")
}
