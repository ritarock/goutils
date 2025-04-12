package slice

import (
	"cmp"
	"slices"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Max[T cmp.Ordered](array []T) T {
	if len(array) == 0 {
		var zero T
		return zero
	}
	return slices.Max(array)
}

func Min[T cmp.Ordered](array []T) T {
	if len(array) == 0 {
		var zero T
		return zero
	}
	return slices.Min(array)
}

func Sum[T Number](array []T) T {
	if len(array) == 0 {
		return 0
	}

	var sum T
	for _, v := range array {
		sum += v
	}

	return sum
}

func Average[T Number](array []T) float64 {
	if len(array) == 0 {
		return 0
	}

	sum := Sum(array)

	return float64(sum) / float64(len(array))
}

func Median[T Number](array []T) float64 {
	n := len(array)
	if n == 0 {
		return 0
	}

	sorted := Sort(array)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}

	a, b := float64(sorted[(n/2)-1]), float64(sorted[n/2])
	return (a + b) / 2
}

func Mode[T cmp.Ordered](array []T) []T {
	if len(array) == 0 {
		return array
	}

	freq := make(map[T]int)
	maxCount := 0
	for _, v := range array {
		freq[v]++
		if freq[v] > maxCount {
			maxCount = freq[v]
		}
	}

	var modes []T
	for v, count := range freq {
		if count == maxCount {
			modes = append(modes, v)
		}
	}

	return Sort(modes)
}
