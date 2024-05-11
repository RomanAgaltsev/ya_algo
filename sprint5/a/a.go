/*
A. Лампочки

Гоша повесил на стену гирлянду в виде бинарного дерева, в узлах которого находятся лампочки.
У каждой лампочки есть своя яркость. Уровень яркости лампочки соответствует числу, расположенному в узле дерева.
Помогите Гоше найти самую яркую лампочку в гирлянде, то есть такую, у которой яркость наибольшая.

Формат ввода
На вход подается корень дерева.

Формат вывода
Функция должна вернуть максимальное значение яркости в узле дерева.
*/

package main

// <template>
type Node struct {  
	value  int  
	left   *Node  
	right  *Node  
}
// <template>

func Solution(root *Node) int {
	maxValue := root.value
	maxLeft, maxRight := 0, 0
	if root.left != nil {
		maxLeft = Solution(root.left)
	}
	
	if root.right != nil {
		maxRight = Solution(root.right)
	}
	maxChild := max(maxLeft, maxRight)
	return max(maxValue, maxChild)
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, &node3, nil}
	if Solution(&node4) != 3 {
		panic("WA")
	}
}

func main() {
	test()
}