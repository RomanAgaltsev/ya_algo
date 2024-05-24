/*
D. Деревья - близнецы

Гоше на день рождения подарили два дерева.
Тимофей сказал, что они совершенно одинаковые. Но, по мнению Гоши, они отличаются.
Помогите разрешить этот философский спор!

Формат ввода
На вход подаются корни двух деревьев.

Формат вывода
Функция должна вернуть True если деревья являются близнецами. Иначе - False.
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

func Solution(root1 *Node, root2 *Node) bool {
    return pathLeftToRight(root1) == pathLeftToRight(root2)
}

func test() {
    node1 := Node{1, nil, nil}
    node2 := Node{2, nil, nil}
    node3 := Node{3, &node1, &node2}
    node4 := Node{1, nil, nil}
    node5 := Node{2, nil, nil}
    node6 := Node{3, &node4, &node5}

    if !Solution(&node3, &node6) {
        panic("WA")
    }
}

func main() {
    test()
}
