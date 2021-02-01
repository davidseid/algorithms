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
	result := [][]int{}
	path := []int{}
	visited := map[int]bool{}
	getPermutations(nums, result, path, visited)
	return result
}

func getPermutations(nums []int, result [][]int, path []int, visited map[int]bool) {
	if len(path) == len(nums) {
		result = append(result, path)
		return
	}

	for _, num := range nums {
		if _, ok := visited[num]; !ok {
			path = append(path, num)
			visited[num] = true
			getPermutations(nums, result, path, visited)
			visited[num] = false
			path = path[:len(path)-1]
		}
	}
}
