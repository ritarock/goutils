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

func UnionArray(array ...interface{}) interface{} {
	var unionArray interface{}

	switch array[0].(type) {
	case []int:
		m := make(map[int]struct{})
		for _, arr := range array {
			for _, v := range arr.([]int) {
				m[v] = struct{}{}
			}
		}
		tmp := []int{}
		for k := range m {
			tmp = append(tmp, k)
		}
		unionArray = tmp

	case []float64:
		m := make(map[float64]struct{})
		for _, arr := range array {
			for _, v := range arr.([]float64) {
				m[v] = struct{}{}
			}
		}

		tmp := []float64{}
		for k := range m {
			tmp = append(tmp, k)
		}
		unionArray = tmp
	}

	return SortArray(unionArray)
}

func IntersectArray(array ...interface{}) interface{} {
	var intersectArray interface{}

	switch array[0].(type) {
	case []int:
		switch length := len(array); length {
		case 1:
			intersectArray = array[0].([]int)
		case 2:
			tmp := []int{}
			m := make(map[int]struct{})
			for _, v := range array[0].([]int) {
				m[v] = struct{}{}
			}
			for _, v := range array[1].([]int) {
				if _, ok := m[v]; !ok {
					continue
				}
				tmp = append(tmp, v)
			}
			intersectArray = tmp
		default:
			tmp := []int{}
			firstArr := IntersectArray(array[0], array[1])

			for i := 0; i < len(array)-2; i++ {
				func(arr1, arr2 []int) {
					m := make(map[int]struct{})
					for _, v := range arr1 {
						m[v] = struct{}{}
					}
					for _, v := range arr2 {
						if _, ok := m[v]; !ok {
							continue
						}
						tmp = append(tmp, v)
					}
				}(firstArr.([]int), array[i+2].([]int))
			}
			intersectArray = UniqueArray(tmp)
		}

	case []float64:
		switch length := len(array); length {
		case 1:
			intersectArray = array[0].([]float64)
		case 2:
			tmp := []float64{}
			m := make(map[float64]struct{})
			for _, v := range array[0].([]float64) {
				m[v] = struct{}{}
			}
			for _, v := range array[1].([]float64) {
				if _, ok := m[v]; !ok {
					continue
				}
				tmp = append(tmp, v)
			}
			intersectArray = tmp
		default:
			tmp := []float64{}
			firstArr := IntersectArray(array[0], array[1])

			for i := 0; i < len(array)-2; i++ {
				func(arr1, arr2 []float64) {
					m := make(map[float64]struct{})
					for _, v := range arr1 {
						m[v] = struct{}{}
					}
					for _, v := range arr2 {
						if _, ok := m[v]; !ok {
							continue
						}
						tmp = append(tmp, v)
					}
				}(firstArr.([]float64), array[i+2].([]float64))
			}

			intersectArray = UniqueArray(tmp)
		}
	}

	return SortArray(intersectArray)
}

func DifferenceArray(array1, array2 interface{}) interface{} {
	var differenceArray interface{}

	switch array1.(type) {
	case []int:
		m := make(map[int]struct{})
		for _, v := range array2.([]int) {
			m[v] = struct{}{}
		}

		tmp := []int{}
		for _, v := range array1.([]int) {
			if _, ok := m[v]; ok {
				continue
			}
			tmp = append(tmp, v)
		}
		differenceArray = tmp
	case []float64:
		m := make(map[float64]struct{})
		for _, v := range array2.([]float64) {
			m[v] = struct{}{}
		}

		tmp := []float64{}
		for _, v := range array1.([]float64) {
			if _, ok := m[v]; ok {
				continue
			}
			tmp = append(tmp, v)
		}
		differenceArray = tmp

	}

	return differenceArray
}
