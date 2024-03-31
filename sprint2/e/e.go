/*
E. Всё наоборот

Вася решил запутать маму — делать дела в обратном порядке.
Список его дел теперь хранится в двусвязном списке.
Напишите функцию, которая вернёт список в обратном порядке.
Внимание: в этой задаче не нужно считывать входные данные.
Нужно написать только функцию, которая принимает на вход голову двусвязного списка и возвращает голову перевёрнутого списка.

Формат ввода
Функция принимает на вход единственный аргумент — голову двусвязного списка.
Длина списка не превосходит 1000 элементов. Список не бывает пустым.

Формат вывода
Функция должна вернуть голову развернутого списка.

*/

package main

// <template>
type ListNode struct {  
    data     string  
    next  *ListNode  
    prev  *ListNode  
}
// <template>

func Solution (head *ListNode) *ListNode {
	node := head
	var tmp, newHead *ListNode
	for node != nil {
		tmp = node.next
		node.next = node.prev
		if tmp != nil {
			node.prev = tmp
			node = tmp
		} else {
			newHead = node
			node.prev = nil
			node = nil
		}
	}
	return newHead
}

func test() {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0
	/*newHead :=*/ Solution(&node0)
	// result is : newHead == node3
	// node3.next == node2
	// node2.next == node1
	// node2.prev = node3
	// node1.next == node0
	// node1.prev == node2
	// node0.prev == node1
}