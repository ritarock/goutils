package array

import (
	"sort"

	"github.com/golang-collections/collections/set"
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
		aSet := set.New()
		bSet := set.New()
		for i, arr := range array {
			if i == 0 {
				for _, v := range arr.([]int) {
					aSet.Insert(v)
				}
			}
			for _, v := range arr.([]int) {
				bSet.Insert(v)
			}
		}
		tmp := aSet.Union(bSet)
		tmp2 := []int{}
		tmp.Do(func(i interface{}) { tmp2 = append(tmp2, i.(int)) })
		unionArray = tmp2
	case []float64:
		aSet := set.New()
		bSet := set.New()
		for i, arr := range array {
			if i == 0 {
				for _, v := range arr.([]float64) {
					aSet.Insert(v)
				}
			}
			for _, v := range arr.([]float64) {
				bSet.Insert(v)
			}
		}
		tmp := aSet.Union(bSet)
		tmp2 := []float64{}
		tmp.Do(func(i interface{}) { tmp2 = append(tmp2, i.(float64)) })
		unionArray = tmp2
	}

	return SortArray(unionArray)
}

func IntersectArray(array ...interface{}) interface{} {
	var intersectArray interface{}

	switch array[0].(type) {
	case []int:
		switch length := len(array); length {
		case 2:
			aSet := set.New()
			bSet := set.New()
			for i, arr := range array {
				if i == 0 {
					for _, v := range arr.([]int) {
						aSet.Insert(v)
					}
					for _, v := range arr.([]int) {
						bSet.Insert(v)
					}
				}
			}
			tmp := aSet.Intersection(bSet)
			tmp2 := []int{}
			tmp.Do(func(i interface{}) { tmp2 = append(tmp2, i.(int)) })
			intersectArray = tmp2

		default:
			firstArr := IntersectArray(array[0], array[1])
			tmp := []int{}
			for i := 0; i < len(array)-2; i++ {
				func(arr1, arr2 []int) {
					aSet := set.New()
					bSet := set.New()
					for _, v := range arr1 {
						aSet.Insert(v)
					}
					for _, v := range arr2 {
						bSet.Insert(v)
					}
					fTmp := aSet.Intersection(bSet)
					fTmp.Do(func(i interface{}) { tmp = append(tmp, i.(int)) })
				}(firstArr.([]int), array[i+2].([]int))
			}
			intersectArray = UniqueArray(tmp)
		}

	case []float64:
		switch length := len(array); length {
		case 2:
			aSet := set.New()
			bSet := set.New()
			for i, arr := range array {
				if i == 0 {
					for _, v := range arr.([]float64) {
						aSet.Insert(v)
					}
					for _, v := range arr.([]float64) {
						bSet.Insert(v)
					}
				}
			}
			tmp := aSet.Intersection(bSet)
			tmp2 := []float64{}
			tmp.Do(func(i interface{}) { tmp2 = append(tmp2, i.(float64)) })
			intersectArray = tmp2

		default:
			firstArr := IntersectArray(array[0], array[1])
			tmp := []float64{}
			for i := 0; i < len(array)-2; i++ {
				func(arr1, arr2 []float64) {
					aSet := set.New()
					bSet := set.New()
					for _, v := range arr1 {
						aSet.Insert(v)
					}
					for _, v := range arr2 {
						bSet.Insert(v)
					}
					fTmp := aSet.Intersection(bSet)
					fTmp.Do(func(i interface{}) { tmp = append(tmp, i.(float64)) })
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
		aSet := set.New()
		bSet := set.New()
		for _, v := range array1.([]int) {
			aSet.Insert(v)
		}
		for _, v := range array2.([]int) {
			bSet.Insert(v)
		}
		tmp := aSet.Difference(bSet)
		tmp2 := []int{}
		tmp.Do(func(i interface{}) { tmp2 = append(tmp2, i.(int)) })
		differenceArray = tmp2
	case []float64:
		aSet := set.New()
		bSet := set.New()
		for _, v := range array1.([]float64) {
			aSet.Insert(v)
		}
		for _, v := range array2.([]float64) {
			bSet.Insert(v)
		}
		tmp := aSet.Difference(bSet)
		tmp2 := []float64{}
		tmp.Do(func(i interface{}) { tmp2 = append(tmp2, i.(float64)) })
		differenceArray = tmp2
	}

	return differenceArray
}
