package array

import (
	"sort"
)

func SortArray(array interface{}) interface{} {
	switch arr := array.(type) {
	case []int:
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	case []float64:
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}
	return array
}

func ReverseSortArray(array interface{}) interface{} {
	switch arr := array.(type) {
	case []int:
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
	case []float64:
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
	}
	return array
}

func UniueArray(array interface{}) interface{} {
	var newArray interface{}

	switch arr := array.(type) {
	case []int:
		m := make(map[int]struct{})
		newArray = make([]int, 0)

		for _, v := range arr {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newArray = append(newArray.([]int), v)
			}
		}
	case []float64:
		m := make(map[float64]struct{})
		newArray = make([]float64, 0)

		for _, v := range arr {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newArray = append(newArray.([]float64), v)
			}
		}
	}

	return newArray
}

func MaxOfArray(array interface{}) interface{} {
	var max interface{}

	switch arr := array.(type) {
	case []int:
		max = arr[0]
		for _, v := range arr {
			if max.(int) < v {
				max = v
			}
		}
	case []float64:
		max = arr[0]
		for _, v := range arr {
			if max.(float64) < v {
				max = v
			}
		}
	}
	return max
}

func MinOfArray(array interface{}) interface{} {
	var min interface{}

	switch arr := array.(type) {
	case []int:
		min = arr[0]
		for _, v := range arr {
			if min.(int) > v {
				min = v
			}
		}
	case []float64:
		min = arr[0]
		for _, v := range arr {
			if min.(float64) > v {
				min = v
			}
		}
	}
	return min
}

func SumOfArray(array interface{}) interface{} {
	var sum interface{}

	switch arr := array.(type) {
	case []int:
		var s int
		for _, v := range arr {
			s += v
		}
		sum = s
	case []float64:
		var s float64
		for _, v := range arr {
			s += v
		}
		sum = s
	}
	return sum
}

func AvarageOfArray(array interface{}) float64 {
	var avarage float64
	sumOfArray := SumOfArray(array)

	switch arr := array.(type) {
	case []int:
		avarage = float64(sumOfArray.(int)) / float64(len(arr))
	case []float64:
		avarage = float64(sumOfArray.(float64)) / float64(len(arr))
	}

	return avarage
}
