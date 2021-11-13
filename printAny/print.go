package main

import "fmt"

func printAny[G any] (arr []G) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func main() {
	strs := []string{"Hello", "World",  "Generics"}
	decs := []float64{3.14, 1.14, 1.618, 2.718 }
	nums := []int{2,4,6,8}
	printAny(strs)
	printAny(decs)
	printAny(nums)
}