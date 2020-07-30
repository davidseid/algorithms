package algorithms

import "math"

/*
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/

// Find O(n) solution, then try a divide and conquer approach

func MaxSubArray(nums []int) int {

	var maxSoFar float64 = float64(nums[0])
	var currMax float64 = float64(nums[0])

	for i := 1; i < len(nums); i++ {
		currMax = math.Max(float64(nums[i]), currMax+float64(nums[i]))
		maxSoFar = math.Max(currMax, maxSoFar)
	}
	return int(maxSoFar)
}
