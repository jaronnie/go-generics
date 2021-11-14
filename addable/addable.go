package main

import (
	"fmt"
	"runtime"
)

// Addable
// go1.18 确定泛型约束方案
type Addable interface {
	int | int8 | int16 | int32 | string
}

func add[T Addable](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(runtime.Version())
	fmt.Println(add(1,2))
	fmt.Println(add("1", "2"))
}