package removedupes

import (
	"testing"

	"github.com/go-test/deep"
)

/*
80. Remove Duplicates from Sorted Array II
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,0,1,1,1,1,2,3,3],

Your function should return length = 7, with the first seven elements of nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively.

It doesn't matter what values are set beyond the returned length.
Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

Rationale:
- Use two pointers to iterate through the input
The first pointer tracks the current position for setting
The second pointer tracks the current position for getting
If the current number to get is greater than the number two behind,
then it is a candidate to set, so we set it and increment the setting pointer
Otherwise just increment the getting pointer

Optimization:
- Optimized by returning early if the input has 2 or less numbers since we won't need to modify
- Optimized by starting our pointers at index 2, since we can ignore the first two
*/

func removeDuplicates(nums []int) int {
	i := 2

	len := len(nums)

	if len <= 2 {
		return len
	}

	for n := 2; n < len; n++ {
		if nums[n] > nums[i-2] {
			nums[i] = nums[n]
			i++
		}
	}

	return i
}

func TestRemoveDuplicates(t *testing.T) {
	t.Run("should remove duplicates", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 3}

		expected := []int{1, 1, 2, 2, 3}

		actual := removeDuplicates(nums)

		if diff := deep.Equal(nums[:len(expected)], expected); diff != nil {
			t.Error(diff)
		}

		if len(expected) != actual {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})

	t.Run("should remove duplicates 2", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

		expected := []int{0, 0, 1, 1, 2, 3, 3}

		actual := removeDuplicates(nums)

		if diff := deep.Equal(nums[:len(expected)], expected); diff != nil {
			t.Error(diff)
		}

		if len(expected) != actual {
			t.Errorf("Got %v, want %v", actual, expected)
		}
	})
}
