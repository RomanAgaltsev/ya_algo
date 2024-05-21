package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}
// <template>

func successor(node *Node) *Node {
	curr := node
	for curr != nil && curr.left != nil {
		curr = curr.left
	}
	return curr
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}
	if key < node.value {
		node.left = remove(node.left, key)
	} else if key > node.value {
		node.right = remove(node.right, key)
	} else {
		if node.left == nil {
			right := node.right
			node = nil
			return right
		} else if node.right == nil {
			left := node.left
			node = nil
			return left
		}
		succ := successor(node.right)
		node.value = succ.value
		node.right = remove(node.right, succ.value)
	}
	return node
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}
	newHead := remove(&node7, 10)
	if newHead.value != 5 {
		panic("WA")
	}
	if newHead.right != &node5 {
		panic("WA")
	}
	if newHead.right.value != 8 {
		panic("WA")
	}
}

func main() {
	test()
}