/*
C. Дерево - анаграмма

Гоша и Алла играют в игру «Удивительные деревья».
Помогите ребятам определить, является ли дерево, которое им встретилось, деревом-анаграммой?
Дерево называется анаграммой, если оно симметрично относительно своего центра.

Формат ввода
Напишите функцию, которая определяет, является ли дерево анаграммой.
На вход подаётся корень дерева.

Формат вывода
Функция должна вернуть True если дерево является анаграммой. Иначе - False.
*/

package main

import "fmt"

// <template>
type Node struct {
    value int
    left  *Node
    right *Node
}
// <template>

func pathLeftToRight(node *Node) string {
    if node == nil {
        return ""
    }
    return pathLeftToRight(node.left) + fmt.Sprint(node.value) + pathLeftToRight(node.right)
}

func pathRightToLeft(node *Node) string {
    if node == nil {
        return ""
    }
    return pathRightToLeft(node.right) + fmt.Sprint(node.value) + pathRightToLeft(node.left)
}

func Solution(root *Node) bool {
    if root == nil {
        return false
    }
    if root.left == nil && root.right == nil {
        return true
    }
    if root.left == nil || root.right == nil {
        return false
    }
    return pathLeftToRight(root.left) == pathRightToLeft(root.right)
}

func test() {
    node1 := Node{3, nil, nil}
    node2 := Node{4, nil, nil}
    node3 := Node{4, nil, nil}
    node4 := Node{3, nil, nil}
    node5 := Node{2, &node1, &node2}
    node6 := Node{2, &node3, &node4}
    node7 := Node{1, &node5, &node6}

    if !Solution(&node7) {
        panic("WA")
    }
}

func main() {
	test()
}
