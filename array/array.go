package array

import "sort"

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sort[T Number](array []T) []T {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	return array
}

func ReverseSort[T Number](array []T) []T {
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})
	return array
}

func Unique[T Number | string](array []T) []T {
	newArray := make([]T, 0)
	m := make(map[T]struct{})

	for _, v := range array {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func MaxOfArray[T Number](array []T) T {
	max := array[0]
	for _, v := range array {
		if max < v {
			max = v
		}
	}
	return max
}

func MinOfArray[T Number](array []T) T {
	min := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
	}
	return min
}

func SumOfArray[T Number](array []T) T {
	var sum T
	for _, v := range array {
		sum += v
	}
	return sum
}

func AvarageOfArray[T Number](array []T) float64 {
	sum := SumOfArray(array)

	return float64(sum) / float64(len(array))
}
