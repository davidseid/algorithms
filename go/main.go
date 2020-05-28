package main

import (
	"fmt"

	"./algorithms"
)

func main() {

	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(algorithms.MaxSubArray(arr))
}
