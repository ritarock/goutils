package array

import (
	"sort"
)

func SortArray(array interface{}) interface{} {
	switch array.(type) {
	case []int:
		sort.Sort(sort.IntSlice(array.([]int)))
	case []float64:
		sort.Sort(sort.Float64Slice(array.([]float64)))
	}

	return array
}

func ReverseSortArray(array interface{}) interface{} {
	switch array.(type) {
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(array.([]int))))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(array.([]float64))))
	}
	return array
}

func UniqueArray(array interface{}) interface{} {
	var newList interface{}

	switch array.(type) {
	case []int:
		m := make(map[int]struct{})
		newList = make([]int, 0)

		for _, v := range array.([]int) {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newList = append(newList.([]int), v)
			}
		}
	case []float64:
		m := make(map[float64]struct{})
		newList = make([]float64, 0)

		for _, v := range array.([]float64) {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newList = append(newList.([]float64), v)
			}
		}
	}

	return newList
}

func MaxOfArray(array interface{}) interface{} {
	var max interface{}
	arr := ReverseSortArray(array)

	switch array.(type) {
	case []int:
		max = arr.([]int)[0]
	case []float64:
		max = arr.([]float64)[0]
	}

	return max
}

func MinOfArray(array interface{}) interface{} {
	var min interface{}
	arr := SortArray(array)

	switch array.(type) {
	case []int:
		min = arr.([]int)[0]
	case []float64:
		min = arr.([]float64)[0]
	}

	return min
}

func SumOfArray(array interface{}) interface{} {
	var sum interface{}

	switch array.(type) {
	case []int:
		var s int
		for _, v := range array.([]int) {
			s += v
		}
		sum = s
	case []float64:
		var s float64
		for _, v := range array.([]float64) {
			s += v
		}
		sum = s
	}

	return sum
}

func AvarageOfArray(array interface{}) interface{} {
	var avarage interface{}

	sumOfArray := SumOfArray(array)

	switch array.(type) {
	case []int:
		avarage = sumOfArray.(int) / len(array.([]int))
	case []float64:
		avarage = sumOfArray.(float64) / float64(len(array.([]float64)))
	}

	return avarage
}

func UnionArray(array1, array2 interface{}) interface{} {
	var unionArray interface{}

	switch array1.(type) {
	case []int:
		for _, v := range array2.([]int) {
			array1 = append(array1.([]int), v)
		}
		unionArray = UniqueArray(array1)
	case []float64:
		for _, v := range array2.([]float64) {
			array1 = append(array1.([]float64), v)
		}
		unionArray = UniqueArray(array1)
	}

	return unionArray
}

func IntersectArray(array1, array2 interface{}) interface{} {
	var intersectArray interface{}

	switch array1.(type) {
	case []int:
		tmp := []int{}
		for _, v := range array1.([]int) {
			for _, v2 := range array2.([]int) {
				if v == v2 {
					tmp = append(tmp, v)
				}
			}
		}
		intersectArray = tmp
	case []float64:
		tmp := []float64{}
		for _, v := range array1.([]float64) {
			for _, v2 := range array2.([]float64) {
				if v == v2 {
					tmp = append(tmp, v)
				}
			}
		}
		intersectArray = tmp
	}

	return intersectArray
}
