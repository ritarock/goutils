package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](array []T) []T {
	copied := make([]T, len(array))
	copy(copied, array)

	sort.Slice(copied, func(i, j int) bool {
		return copied[i] < copied[j]
	})

	return copied
}

func ReverseSort[T constraints.Ordered](array []T) []T {
	copied := make([]T, len(array))
	copy(copied, array)

	sort.Slice(copied, func(i, j int) bool {
		return copied[i] > copied[j]
	})

	return copied
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
	if len(array) == 0 {
		return []T{}
	}

	m := make(map[T]struct{}, len(array))
	result := []T{}

	for _, v := range array {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func Include[T constraints.Ordered](array []T, value T) bool {
	if len(array) == 0 {
		return false
	}

	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}
