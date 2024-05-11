package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}
// <template>

func isValidBST(node *Node) (bool, int, int) {
	if node.left == nil && node.right == nil {
		return true, node.value, node.value
	}
	
	minValue := node.value
	maxValue := node.value
	
	isValidLeft := true
	var leftMin, leftMax int
	if node.left != nil {
		isValidLeft, leftMin, leftMax = isValidBST(node.left)
		if !isValidLeft {
			return false, 0, 0
		}
		if node.value <= leftMax {
			return false, 0, 0
		}
		minValue = leftMin
	}
	
	isValidRight := true
	var rightMin, rightMax int
	if node.right != nil {
		isValidRight, rightMin, rightMax = isValidBST(node.right)
		if !isValidRight {
			return false, 0, 0
		}
		if node.value >= rightMin {
			return false, 0, 0
		}
		maxValue = rightMax
	}
	return true, minValue, maxValue
}

func Solution(root *Node) bool {
	if root == nil {
		return true
	}
	isValid, _, _ := isValidBST(root)
	return isValid
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{8, nil, nil}
	node5 := Node{5, &node3, &node4}
	if !Solution(&node5) {
		panic("WA")
	}
	node2.value = 5
	if Solution(&node5) {
		panic("WA")
	}
}
