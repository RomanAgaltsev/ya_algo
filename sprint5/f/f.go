/*
F. Максимальная глубина

Алла хочет побывать на разных островах архипелага Алгосы.
Она составила карту. Карта представлена в виде дерева: корень обозначает центр архипелага, узлы –— другие острова.
А листья —– это дальние острова, на которые Алла хочет попасть.
Помогите Алле определить максимальное число островов, через которые ей нужно пройти для совершения одной поездки
от стартового острова до места назначения, включая начальный и конечный пункты.

Формат ввода
На вход подается корень дерева.

Формат вывода
Функция должна вернуть число, равное максимальному числу островов в пути (включая начальный и конечный пункты).
*/

package main

// <template>
type Node struct {
    value int
    left  *Node
    right *Node
}
// <template>

func Solution(root *Node) int {
    if root == nil {
        return 0
    }
    if root.left == nil && root.right == nil {
        return 1
    }
    return 1 + max(Solution(root.left), Solution(root.right))
}

func test() {
    node1 := Node{1, nil, nil}
    node2 := Node{4, nil, nil}
    node3 := Node{3, &node1, &node2}
    node4 := Node{8, nil, nil}
    node5 := Node{5, &node3, &node4}
    if Solution(&node5) != 3 {
        panic("WA")
    }
}

func main() {
    test()
}
