package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Max[T Number](array []T) T {
	if len(array) == 0 {
		return 0
	}

	sorted := ReverseSort(array)
	return sorted[0]
}

func Min[T Number](array []T) T {
	if len(array) == 0 {
		return 0
	}

	sorted := Sort(array)
	return sorted[0]
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
		b := sorted[length/2]
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

	return Sort(result)
}
