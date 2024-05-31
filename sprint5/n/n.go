/*
N. Разбиение дерева

Дано бинарное дерево поиска, в котором хранятся целые числа. От этого дерева надо отделить k самых маленьких элементов.
Реализуйте функцию, которая принимает корень дерева и число k, а возвращает два BST — в первом k наименьших элементов из исходного дерева, а во втором — оставшиеся вершины BST.
В вершинах дерева уже записаны корректные размеры поддеревьев (точное название поля смотрите в заготовках кода).
После разбиения размеры должны остаться корректными — вам придётся пересчитывать их на ходу.
Ваше решение должно иметь асимптотику O(h), где h — высота исходного дерева.

Формат ввода
Числа, записанные в вершинах дерева, лежат в диапазоне от 0 до 109. Дерево не содержит одинаковых ключей.
Число вершин в дереве не превосходит 105.
*/

package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
	size  int
}
// <template>

func split(node *Node, k int) (*Node, *Node) {
	leftSize := 0
	if node.left != nil {
		leftSize = node.left.size
	}
	
	if leftSize == k {
		
	}
	
	if leftSize < k {
		return split(node.right, k-leftSize-1)
	}
	
	return split(node.left, k)
}

func test() {
	node1 := &Node{3, nil, nil, 1}
	node2 := &Node{2, nil, node1, 2}
	node3 := &Node{8, nil, nil, 1}
	node4 := &Node{11, nil, nil, 1}
	node5 := &Node{10, node3, node4, 3}
	node6 := &Node{5, node2, node5, 6}
	left, right := split(node6, 4)
	if left.size != 4 {
		panic("WA")
	}
	if right.size != 2 {
		panic("WA")
	}
}
