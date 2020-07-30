package manual

import (
	"fmt"
	"testing"
)

/*
Problem 1-28

Write a function to perform integer division without using either the / or * operators. Find a fast way to do it.
*/

func divideInteger(dividend, divisor int) int {

	if divisor == 0 {
		panic("Can't devide by zero")
	}

	quotient := 0

	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}

	return quotient
}

func TestDivideInteger(t *testing.T) {
	actual := divideInteger(1395871, 17)
	fmt.Println(actual)
}
