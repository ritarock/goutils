package array

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sort[T Number](data []T) []T {
	copied := make([]T, len(data))
	copy(copied, data)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i] < copied[j]
	})
	return copied
}

func ReverseSort[T Number](data []T) []T {
	copied := make([]T, len(data))
	copy(copied, data)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i] > copied[j]
	})
	return copied
}

func Unique[T Number | string](data []T) []T {
	newdata := make([]T, 0, len(data))
	m := make(map[T]struct{}, len((data)))

	for _, v := range data {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			newdata = append(newdata, v)
		}
	}
	return newdata
}

func MaxOfArray[T Number](data []T) T {
	max := data[0]
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	return max
}

func MinOfArray[T Number](data []T) T {
	min := data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
	}
	return min
}

func SumOfArray[T Number](data []T) T {
	var sum T
	for _, v := range data {
		sum += v
	}
	return sum
}

func MeanOfArray[T Number](data []T) float64 {
	sum := SumOfArray(data)

	return float64(sum) / float64(len(data))
}

func Median[T Number](data []T) float64 {
	sorted := Sort(data)

	if len(sorted) == 0 {
		return 0
	}

	if len(sorted)%2 == 0 {
		return (float64(sorted[len(sorted)/2]) + float64(sorted[len(sorted)/2]-1)) / 2
	} else {
		return float64(sorted[len(sorted)/2])
	}
}
