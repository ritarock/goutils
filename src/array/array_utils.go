package array

import "sort"

func ToFloat64(array interface{}) []float64 {
	var result []float64

	switch array.(type) {
	case []float64:
		result = array.([]float64)
	case []int:
		for _, v := range array.([]int) {
			result = append(result, float64(v))
		}
	default:
		result = append(result, 0)
	}
	return result
}

func MaxOfArray(array []float64) float64 {
	max := array[0]
	for _, v := range array {
		if max < v {
			max = v
		}
	}
	return max
}

func MinOfArray(array []float64) float64 {
	min := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
	}
	return min
}

func SumOfArray(array []float64) float64 {
	var sum float64 = 0
	for _, v := range array {
		sum += v
	}
	return sum
}

func AvarageOfArray(array []float64) float64 {
	sumOfArray := SumOfArray(array)
	return sumOfArray / float64(len(array))
}

func SortArray(array []float64) []float64 {
	sort.Sort(sort.Float64Slice(array))
	return array
}

func ReverseSortArray(array []float64) []float64 {
	sort.Sort(sort.Reverse(sort.Float64Slice(array)))
	return array
}

func UniqueArray(array []float64) []float64 {
	m := make(map[float64]struct{})
	newList := make([]float64, 0)

	for _, v := range array {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			newList = append(newList, v)
		}
	}

	return newList
}
