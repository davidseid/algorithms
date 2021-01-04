package clonegraphs

func dfsCloneGraph(node *Node) *Node {
	visited := map[int]*Node{}

	return dfs(node, visited)
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
