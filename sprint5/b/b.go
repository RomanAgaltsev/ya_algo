/*
B. Сбалансированное дерево

Гоше очень понравилось слушать рассказ Тимофея про деревья.
Особенно часть про сбалансированные деревья.
Он решил написать функцию, которая определяет, сбалансировано ли дерево.
Дерево считается сбалансированным, если левое и правое поддеревья каждой вершины отличаются по высоте не больше, чем на единицу.

Реализуйте функцию, определяющую, является ли дерево сбалансированным.

Формат ввода
На вход функции подаётся корень бинарного дерева.

Формат вывода
Функция должна вернуть True, если дерево сбалансировано в соответствии с критерием из условия, иначе - False.
*/

package main

// <template>
type Node struct {  
	value  int  
	left   *Node  
	right  *Node  
}
// <template>

func getTreeHeight(root *Node, height int) int {
	if root == nil {
		return height
	}
	return max(getTreeHeight(root.left, height+1), getTreeHeight(root.right, height+1))
}

func Solution(root *Node) bool {
	if root == nil {
		return true
	}
	var leftHeight, rightHeight int
	if root.left != nil {
		leftHeight = getTreeHeight(root.left, 0)
	}
	if root.right != nil {
		rightHeight = getTreeHeight(root.right, 0)
	}
	
	heightDiff := leftHeight - rightHeight
	
	return heightDiff >= -1 && heightDiff <= 1 && Solution(root.left) && Solution(root.right)
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{10, nil, nil}
	node5 := Node{2, &node3, &node4}
	if (!Solution(&node5)) {
		panic("WA")
	}
}

func main() {
	test()
}