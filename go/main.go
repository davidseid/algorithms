package main

import (
	"fmt"

	"./algorithms"
)

func main() {
	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	algorithms.SetZeroes(matrix)
	fmt.Println(matrix)
}
