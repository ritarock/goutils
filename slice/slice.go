package slice

import (
	"math"
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

func Max[T Number](array []T) T {
	if len(array) == 0 {
		return 0
	}

	max := array[0]
	for _, v := range array[1:] {
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
	for _, v := range array[1:] {
		if v < min {
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
		a := sorted[(length/2)-1]
		b := sorted[(length / 2)]
		return float64((a + b) / 2)
	} else {
		return float64(sorted[length/2])
	}
}

func Mode[T Number](array []T) []T {
	if len(array) == 0 {
		return []T{}
	}

	m := make(map[T]int, len(array))
	for _, v := range array {
		m[v]++
	}

	type kv struct {
		key   T
		value int
	}
	kvs := []kv{}
	for k, v := range m {
		kvs = append(kvs, kv{key: k, value: v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value > kvs[j].value
	})

	result := []T{kvs[0].key}
	for _, kv := range kvs[1:] {
		if kvs[0].value != kv.value {
			break
		}
		result = append(result, kv.key)
	}

	return result
}

func StandardDeviation[T Number](array []T) float64 {
	if len(array) == 0 {
		return 0
	}

	mean := Mean(array)
	deviation := make([]float64, len(array))
	for i, v := range array {
		deviation[i] = float64(v) - mean
	}

	variance := func(deviation []float64) float64 {
		var sum float64
		for _, v := range deviation {
			sum += math.Pow(v, 2)
		}
		return sum / float64(len(deviation))
	}(deviation)

	return math.Sqrt(variance)
}

func Some[T constraints.Ordered](array []T, f func(T) bool) bool {
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

func Every[T constraints.Ordered](array []T, f func(T) bool) bool {
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

func Filter[T constraints.Ordered](array []T, f func(T) bool) []T {
	if len(array) == 0 {
		return []T{}
	}

	result := []T{}
	for _, v := range array {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Unique[T constraints.Ordered](array []T) []T {
	result := make([]T, 0, len(array))
	m := make(map[T]struct{}, len(array))

	for _, v := range array {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func Include[T any](array []T, value T) bool {
	if len(array) == 0 {
		return false
	}

	m := make(map[any]struct{}, len(array))
	for _, v := range array {
		m[v] = struct{}{}
	}
	_, ok := m[value]
	return ok
}
