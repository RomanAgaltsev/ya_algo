package main

import (
	"bufio"
	"errors"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
)

var (
	hashTableCapacity int = 1000001
	errValueIsAbsent = errors.New("value is absent")
)

type Command struct {
	name string // Имя команды
	key int // Ключ
	value int // Значение
}

type node struct {
	key int
	value int
	next *node
}

func newNode(key, value int, next *node) *node {
	return &node{
		key: key,
		value: value,
		next: next,
	}
}

type HashTable struct {
	capacity int
	table []*node
}

func (ht *HashTable) index(key int) int {
	h := fnv.New32()
	h.Write([]byte(strconv.Itoa(key)))
	hash := int(h.Sum32())
	return hash % ht.capacity
}

func (ht *HashTable) findNode(n *node, key int) *node {
	for n != nil {
		if n.key == key {
			return n
		}
		n = n.next
	}
	return nil
}

func (ht *HashTable) deleteNode(n *node, key int) (*node, *node) {
	if n == nil {
		return nil, nil
	}
	if n.key == key {
		return n, n.next
	}
	var prev *node
	head, curr := n, n
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			return curr, head
		} 
		prev = curr
		curr = curr.next
	}
	return nil, head
}

func (ht *HashTable) put(key, value int) {
	i := ht.index(key)
	node := ht.table[i]
	if node == nil {
		ht.table[i] = newNode(key, value, nil)
	} else if nodeByKey := ht.findNode(node, key); nodeByKey != nil {
		node.value = value
	} else {
		ht.table[i] = newNode(key, value, node)
	}
}

func (ht *HashTable) get(key int) (int, error) {
	i := ht.index(key)
	node := ht.findNode(ht.table[i], key)
	if node != nil {
		return node.value, nil
	}
	return 0, errValueIsAbsent
}

func (ht *HashTable) delete(key int) (int, error) {
	i := ht.index(key)
	node := ht.table[i]
	deletedNode, headNode := ht.deleteNode(node, key)
	ht.table[i] = headNode
	if deletedNode != nil {
		return deletedNode.value, nil
	}
	return 0, errValueIsAbsent
	
}

func NewHashTable() *HashTable {
	return &HashTable{
		capacity: hashTableCapacity,
		table: make([]*node, hashTableCapacity),
	}
}

func main() {
	scanner := makeScanner()
	commandsNumber := readInt(scanner)
	hashTable := NewHashTable()
	result := ""
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < commandsNumber; i++ {
		result = executeCommand(hashTable, readCommand(scanner))
		if result != "" {
			writer.WriteString(result)
			writer.WriteString("\n")
		}
	}
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 10 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readCommand(scanner *bufio.Scanner) Command {
	command := Command{}
	scanner.Scan()
	listCommand := strings.Split(scanner.Text(), " ")
	if len(listCommand) >= 2 {
		command.name = listCommand[0]
		command.key, _ = strconv.Atoi(listCommand[1])
	}
	if len(listCommand) == 3 {
		command.value, _ = strconv.Atoi(listCommand[2])
	}
	return command
}

func executeCommand(hashTable *HashTable, command Command) string {
	var err error
	var value int
	var hasValue bool
	switch command.name {
	case "put":
		hashTable.put(command.key, command.value)
	case "get":
		value, err = hashTable.get(command.key)
		hasValue = true
	case "delete":
		value, err = hashTable.delete(command.key)
		hasValue = true
	}
	if err != nil {
		return "None"
	}
	if hasValue {
		return strconv.Itoa(value)
	}
	return ""
}