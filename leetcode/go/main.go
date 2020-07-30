package main

import (
	"fmt"

	"./algorithms"
)

func main() {
	matrix := [][]int{
		[]int{-5, 7, 2147483647, 3},
		[]int{0, 3, 6, -2147483648},
		[]int{8, 3, -3, -6},
		[]int{-9, -9, 8, 0},
	}
	fmt.Println(matrix)
	algorithms.SetZeroes(matrix)
	fmt.Println(matrix)
}
