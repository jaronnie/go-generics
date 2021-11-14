package main

import "fmt"

type ComputeAble interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func mapFunc[T ComputeAble, M ComputeAble](a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func main() {
	ii := []int{1, 2, 3, 4}
	miSquareRes := mapFunc(ii, func(t int) int {
		return t * t
	})
	fmt.Println(miSquareRes)
	miAddRes := mapFunc(ii, func(t int) int {
		return t + t
	})
	fmt.Println(miAddRes)

	fi := []float64{1.1, 2.2, 3.3}
	mfSquareRes := mapFunc(fi, func(t float64) float64 {
		return t * t
	})
	fmt.Println(mfSquareRes)
	mfAddRes := mapFunc(fi, func(t float64) float64 {
		return t + t
	})
	fmt.Println(mfAddRes)

}