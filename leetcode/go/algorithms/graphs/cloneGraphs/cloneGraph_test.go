package cloneGraphs

import "testing"

/*
133. Clone Graph
https://leetcode.com/problems/clone-graph/

Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a val (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}

Test case format:

For simplicity sake, each node's value is the same as the node's index (1-indexed). For example, the first node with val = 1, the second node with val = 2, and so on. The graph is represented in the test case using an adjacency list.

Adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.



Example 1:

Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

Example 2:

Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph consists of only one node with val = 1 and it does not have any neighbors.

Example 3:

Input: adjList = []
Output: []
Explanation: This an empty graph, it does not have any nodes.

Example 4:

Input: adjList = [[2],[1]]
Output: [[2],[1]]

Constraints:

    1 <= Node.val <= 100
    Node.val is unique for each node.
    Number of Nodes will not exceed 100.
    There is no repeated edges and no self-loops in the graph.
    The Graph is connected and all nodes can be visited starting from the given node.
*/

func TestCloneGraph(t *testing.T) {
	t.Run("should return clone of original populated graph", func(t *testing.T) {

	})

	t.Run("should return clone of graph with only one node", func(t *testing.T) {

	})

	t.Run("should return clone of empty graph", func(t *testing.T) {

	})

	t.Run("should return clone of two-node graph", func(t *testing.T) {

	})
}

// Queue Implementation with Doubly Linked List
type Queue struct {
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Val      *Node
	Next     *ListNode
	Previous *ListNode
}

func (q *Queue) Enqueue(node *Node) {
	listNode := &ListNode{
		Val: node,
	}

	if q.Head == nil {
		q.Head = listNode
		q.Tail = listNode
		return
	}

	q.Tail.Next = listNode
	q.Tail = listNode
}

func (q *Queue) Dequeue() *Node {
	if q.Head == nil {
		return nil
	}

	node := q.Head
	q.Head = q.Head.Next
	return node.Val
}

// Graph Search Implementation

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := map[int]*Node{}

	return bfs(node, visited)
}

func bfs(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	queue := []*Node{node}

	visited[node.Val] = &Node{Val: node.Val}

	for len(queue) > 0 {
		curr := queue[len(queue)-1]
		queue = queue[1:]
		for _, neighbor := range curr.Neighbors {
			if _, isVisited := visited[neighbor.Val]; !isVisited {
				visited[neighbor.Val] = &Node{
					Val: neighbor.Val,
				}
				queue = append(queue, neighbor)
			}
			visited[curr.Val].Neighbors = append(visited[curr.Val].Neighbors, visited[neighbor.Val])
		}
	}

	return visited[node.Val]
}

func dfs(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, ok := visited[node.Val]; ok {
		return visited[node.Val]
	}

	root := &Node{
		Val: node.Val,
	}

	visited[node.Val] = root

	for _, neighbor := range node.Neighbors {
		root.Neighbors = append(root.Neighbors, dfs(neighbor, visited))
	}

	return root
}
