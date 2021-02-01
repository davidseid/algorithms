package permutations

/*
46. Permutations
https://leetcode.com/problems/permutations/

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:

Input: nums = [1]
Output: [[1]]

Constraints:

    1 <= nums.length <= 6
    -10 <= nums[i] <= 10
    All the integers of nums are unique.
*/

func permute(nums []int) [][]int {

	return getPermutations(nums, []int{})
}

func getPermutations(nums, permutation []int) [][]int {
	permutations := [][]int{}

	if len(nums) == 0 {
		permutations = append(permutations, permutation)
		return permutations
	}

	for i, num := range nums {
		remaining := append([]int{}, nums[:i]...)
		if i+1 < len(nums) {
			remaining = append(remaining, nums[i+1:]...)
		}

		nextPermutation := append([]int{}, permutation...)
		nextPermutation = append(nextPermutation, num)
		permutations = append(permutations, getPermutations(remaining, nextPermutation)...)
	}

	return permutations
}
