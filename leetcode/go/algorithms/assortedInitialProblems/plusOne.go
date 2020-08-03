package algorithms

// Time Compelixity: O(n)
// Space Complexity: O(1)
// Further study: Is there a more efficient way to prepend an integer to a slice in golang?

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	digits = append([]int{1}, digits...)
	return digits
}
