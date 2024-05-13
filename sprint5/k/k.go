/*
K. Выведи диапазон

Напишите функцию, которая будет выводить по неубыванию все ключи от L до R включительно в заданном бинарном дереве поиска.
Ключи в дереве могут повторяться. Решение должно иметь сложность O(h+k), где h — глубина дерева, k — число элементов в ответе.

В данной задаче если в узле содержится ключ x, то другие ключи, равные x, могут быть как в правом, так и в левом поддереве данного узла.
(Дерево строил стажёр, так что ничего страшного).

Формат ввода
На вход функции подаётся корень дерева и искомый ключ. Число вершин в дереве не превосходит 105.
Ключи – натуральные числа, не превосходящие 109. Гарантируется, что L ≤ R.

В итоговом решении не надо определять свою структуру / свой класс, описывающий вершину дерева.

Инструкцию по работе с Make вы можете найти в конце этого урока

Формат вывода
Функция должна напечатать по неубыванию все ключи от L до R по одному в строке.
*/

package main

import (
	"fmt"
)

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}
// <template>

func printRange(root *Node, left int, right int) {
	if root == nil {
		return
	}
	if root.value >= left{
		printRange(root.left, left, right)
	}
	if root.value >= left && root.value <= right {
		fmt.Print(root.value, "\n")
	}
	if root.value <= right {
		printRange(root.right, left, right)
	}
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}
	printRange(&node7, 2, 8)
	// expected output: 2 5 8 8
}

func main() {
	test()
}