/*
J. Добавь узел

Дано BST. Надо вставить узел с заданным ключом. Ключи в дереве могут повторяться.
На вход функции подаётся корень корректного бинарного дерева поиска и ключ, который надо вставить в дерево.
Осуществите вставку этого ключа. Если ключ уже есть в дереве, то его дубликаты уходят в правого сына.
Таким образом вид дерева после вставки определяется однозначно. Функция должна вернуть корень дерева после вставки вершины.
Ваше решение должно работать за O(h), где h — высота дерева.

Формат ввода
Ключи дерева – натуральные числа, не превосходящие 109.
Число вершин в дереве не превосходит 105.
*/

package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}
// <template>

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{value : key}
	}
	if key < root.value {
		root.left = insert(root.left, key)
	} else if key >= root.value {
		root.right = insert(root.right, key)
	}
	return root
}

func test() {
	node1 := Node{7, nil, nil}
	node2 := Node{8, &node1, nil}
	node3 := Node{7, nil, &node2}
	newHead := insert(&node3, 6)
	if newHead != &node3 {
		panic("WA")
	}
	if newHead.left.value != 6 {
		panic("WA")
	}
}

func main() {
	test()
}
