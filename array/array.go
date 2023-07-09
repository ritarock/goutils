package array

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sort[T Number](array []T) []T {
	copied := make([]T, len(array))
	copy(copied, array)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i] < copied[j]
	})
	return copied
}

func ReverseSort[T Number](array []T) []T {
	copied := make([]T, len(array))
	copy(copied, array)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i] > copied[j]
	})
	return copied
}

func Unique[T Number | string](array []T) []T {
	newArray := make([]T, 0, len(array))
	m := make(map[T]struct{}, len(array))

	for _, v := range array {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func Max[T Number](array []T) T {
	if len(array) == 0 {
		return 0
	}

	max := array[0]
	for _, v := range array {
		if max < v {
			max = v
		}
	}
	return max
}

func Min[T Number](array []T) T {
	if len(array) == 0 {
		return 0
	}

	min := array[0]
	for _, v := range array {
		if min > v {
			min = v
		}
	}
	return min
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

func Mean[T Number](array []T) float64 {
	if len(array) == 0 {
		return 0
	}

	sum := Sum(array)
	return float64(sum) / float64(len(array))
}

func Median[T Number](array []T) float64 {
	length := len(array)
	if length == 0 {
		return 0
	}

	sorted := Sort(array)

	if length%2 == 0 {
		return (float64(sorted[(length/2)-1]) + float64(sorted[length/2])) / 2
	} else {
		return float64(sorted[length/2])
	}
}

func Mode[T Number](array []T) []T {
	if len(array) == 0 {
		return []T{0}
	}

	m := make(map[T]int, len(array))

	for _, v := range array {
		m[v] += 1
	}

	type kv struct {
		key   T
		value int
	}
	var kvs []kv
	for k, v := range m {
		kvs = append(kvs, kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value > kvs[j].value
	})
	newArray := []T{kvs[0].key}

	for _, kv := range kvs[1:] {
		if kvs[0].value != kv.value {
			break
		}
		newArray = append(newArray, kv.key)
	}

	return newArray
}

func Some[T Number | string](array []T, f func(T) bool) bool {
	if len(array) == 0 {
		return false
	}
	for _, v := range array {
		if f(v) {
			return true
		}
	}
	return false
}

func Ever[T Number | string](array []T, f func(T) bool) bool {
	if len(array) == 0 {
		return false
	}
	for _, v := range array {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter[T Number | string](array []T, f func(T) bool) []T {
	result := []T{}
	for _, v := range array {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
