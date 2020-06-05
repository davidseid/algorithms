package main

import (
	"fmt"

	"./algorithms"
)

func main() {
	matrix := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	fmt.Println(matrix)
	algorithms.SetZeroes(matrix)
	fmt.Println(matrix)
}
