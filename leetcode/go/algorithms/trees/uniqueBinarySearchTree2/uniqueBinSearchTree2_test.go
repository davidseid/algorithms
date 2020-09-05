package uniqueBinarySearchTree2

import (
	"testing"

	"github.com/go-test/deep"
)

/*
95. Unique Binary Search Trees II
https://leetcode.com/problems/unique-binary-search-trees-ii/

Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:

Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3


Constraints:

0 <= n <= 8
*/

/*
Rationale:
Main algorithm that returns empty slice if n == 0, otherwise call
recursive inner algorithm with 1 and n

Base case of recursive algorithm is when start is greater than end,
this means there are no more nodes left and our remainder list should only include nil

For each index in start to end, recursively call helper on the left and right sides of the index to
generate a list of each side. Then do two nested loops (for each left, for each right) to create a
root, and then append left and right to it, then add it to the result and return at the end.

Followup:
Awesome algorithm but need to review and make sure I understand how to build this from scratch,
was pretty different from my original approach.

Time Complexity:
Not certain, but seems like O(2^n). For each value in n we make recursive calls to build lists for left and right,
which means we do recursive branching to the depth of n.

Space Complexity:

Takeaways:
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Credit:
https://leetcode.com/problems/unique-binary-search-trees-ii/discuss/31494/A-simple-recursive-solution
*/
func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return []*TreeNode{}
	}
	return genTreeList(1, n)
}

func genTreeList(start int, end int) []*TreeNode {
	list := []*TreeNode{}

	if start > end {
		list = append(list, nil)
	}

	for idx := start; idx <= end; idx++ {
		leftList := genTreeList(start, idx-1)
		rightList := genTreeList(idx+1, end)
		for _, left := range leftList {
			for _, right := range rightList {
				root := &TreeNode{Val: idx}
				root.Left = left
				root.Right = right
				list = append(list, root)
			}
		}
	}
	return list
}

func TestGenerateTrees(t *testing.T) {
	t.Run("should generate all unique BSTs", func(t *testing.T) {
		input := 3

		expected := []*TreeNode{
			&TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}},
			&TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
			&TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
			&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		}

		actual := generateTrees(input)

		if diff := deep.Equal(actual, expected); diff != nil {
			t.Error(diff)
		}
	})
}
