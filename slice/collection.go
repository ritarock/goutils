package slice

import (
	"cmp"
	"slices"
)

func Sort[T cmp.Ordered](array []T) []T {
	if len(array) == 0 {
		return array
	}
	copied := make([]T, len(array))
	copy(copied, array)
	itr := slices.Values(copied)
	return slices.Sorted(itr)
}

func ReverseSort[T cmp.Ordered](array []T) []T {
	if len(array) == 0 {
		return array
	}
	copied := make([]T, len(array))
	copy(copied, array)
	itr := slices.Values(copied)
	sorted := slices.Sorted(itr)
	slices.Reverse(sorted)
	return sorted
}

func Some[T any](array []T, f func(T) bool) bool {
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

func Every[T any](array []T, f func(T) bool) bool {
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

func Filter[T any](array []T, f func(T) bool) []T {
	if len(array) == 0 {
		return array
	}

	result := []T{}
	for _, v := range array {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func Unique[T any](array []T) []T {
	if len(array) == 0 {
		return array
	}

	m := make(map[any]struct{}, len(array))
	result := []T{}

	for _, v := range array {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func Include[T cmp.Ordered](array []T, value T) bool {
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
