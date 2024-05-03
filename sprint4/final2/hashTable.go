/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 4
Задача - B. Хеш-таблица

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/24414/run-report/113449962/
427ms/48.09Mb

-- ПРИНЦИП РАБОТЫ --

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --


*/

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
	hashTableCapacity = 1000001 // Размер хеш-таблицы для инициализации
	errValueIsAbsent = errors.New("value is absent") // Ошибка отсутствия значения
)

// сommand - структура команды
type сommand struct {
	name string // Имя команды
	key int // Ключ
	value int // Значение
}

// node - структура связного списка для разрешения коллизий методом цепочек
type node struct {
	key int // Ключ
	value int // Значение
	next *node // Указатель на следующий элемент
}

// newNode - конструктор элемента связного списка
func newNode(key, value int, next *node) *node {
	return &node{
		key: key, // Ключ
		value: value, // Значение
		next: next, // Указатель на следующий элемент
	}
}

// HashTable - структура хеш-таблицы
type HashTable struct {
	capacity int // Размер таблицы
	table []*node // Таблица - слайс указателей на головы связных списков
}

// index - формирует и возращает индекс таблицы по переданному ключу
func (ht *HashTable) index(key int) int {
	// Сначала формируем хеш ключа
	h := fnv.New32()
	h.Write([]byte(strconv.Itoa(key)))
	hash := int(h.Sum32())
	// Затем возвращаем остаток от деления хеша на размер таблицы
	return hash % ht.capacity
}

// findNode - метод хеш-таблицы, выполняющий поиск элемента связного списка
//	по переданным указателю на голову этого списка и ключу.
//	Возвращает указатель на найденный элемент или nil
func (ht *HashTable) findNode(n *node, key int) *node {
	// Обходим связный список, пока не дошли до конца или не нашли нужный элемент
	for n != nil {
		// Если нашли элемент списка с переданным ключом
		if n.key == key {
			// Возвращаем этот элемент
			return n
		}
		// Переходим к следующему элементу
		n = n.next
	}
	// Не нашли элемент связного списка по переданному ключу
	return nil
}

// deleteNode - метод хеш-таблицы, выполняющий удаление элемента связного списка
//	по переданным указателю на голову этого списка и ключу.
//	Возвращает два указателя - на удаленный элемент списка и на голову списка (оба могут быть nil, голова может поменяться)
func (ht *HashTable) deleteNode(n *node, key int) (*node, *node) {
	// Проверяем, есть ли вообще список
	if n == nil {
		// Список пустой (отсутствует), возвращать нечего
		return nil, nil
	}
	// Проверяем, возможно нужный нам элемент - это голова
	if n.key == key {
		// Так и есть, возвращаем текущую и новую головы
		return n, n.next
	}
	// Искомый элемент не голова, надо искать и подменять указатели
	// Предыдущий элемент
	var prev *node
	// Сохраняем голову для возврата и текущий элемент списка
	head, curr := n, n
	// Обходим связный список, пока не дошли до конца или не нашли нужный элемент
	for curr != nil {
		// // Если нашли элемент списка с переданным ключом
		if curr.key == key {
			// Подменяем указатели
			prev.next = curr.next
			// И возвращаем текущий элемент и голову списка
			return curr, head
		}
		// Переходим дальше по связному списку
		prev = curr
		curr = curr.next
	}
	// Нужный элемент списка не нашли, но голову надо вернуть в любом случае
	return nil, head
}

// put - метод хеш-таблицы, выполняющий добавление переданного значения по переданному ключу
func (ht *HashTable) put(key, value int) {
	// По переданному ключу получаем индекс таблицы
	i := ht.index(key)
	// Берем из таблицы по индексу ссылку на голову связного списка
	node := ht.table[i]
	// Проверяем, что получили
	if node == nil {
		// Если получили nil, эта ячейка вообще пустая, связного списка нет
		// Создаем элемент списка и указатель на него записываем в таблицу
		ht.table[i] = newNode(key, value, nil) // Создаваемый элемент ссылается в никуда
	} else if nodeByKey := ht.findNode(node, key); nodeByKey != nil {
		// Если по переданному ключу нашли элемент в списке, обновляем значение - так по условию задачи
		node.value = value
	} else {
		// Иначе добавляем новый элемент в голову списка
		ht.table[i] = newNode(key, value, node) // Создаваемый элемент ссылается на текущую голову
	}
}

// get - метод хеш-таблицы, выполняющий поиск значения по переданному ключу
//	Возвращает значение и nil, если значение по ключу найдено
//	Возвращает 0 и ошибку, если значение по ключу не найдено
func (ht *HashTable) get(key int) (int, error) {
	// По переданному ключу получаем индекс таблицы
	i := ht.index(key)
	// По указателю на голову и ключу ищем элемент в списке
	node := ht.findNode(ht.table[i], key)
	// Проверяем, что нашли
	if node != nil {
		// Нашли элемент, возвращаем его значение и nil в качестве ошибки
		return node.value, nil
	}
	// Не нашли элемент, возвращаем 0 и ошибку
	return 0, errValueIsAbsent
}

// delete - метод хеш-таблицы, выполняющий удаление значения по переданному ключу
//	Возвращает значение и nil, если значение по ключу найдено
//	Возвращает 0 и ошибку, если значение по ключу не найдено
func (ht *HashTable) delete(key int) (int, error) {
	// По переданному ключу получаем индекс таблицы
	i := ht.index(key)
	// Берем из таблицы по индексу ссылку на голову связного списка
	node := ht.table[i]
	// Вызываем метод удаления элемента списка по ключу
	// Получаем указатели на найденный элемент и на голову связного списка (могла измениться)
	deletedNode, headNode := ht.deleteNode(node, key)
	// Указатель на голову записываем в таблицу - если удаляемый элемент является головой, nil
	ht.table[i] = headNode
	// Проверяем, был ли по ключу найден элемент
	if deletedNode != nil {
		// Элемент найден, возвращаем значение и nil
		return deletedNode.value, nil
	}
	// Элемент не найден, возвращаем 0 и ошибку
	return 0, errValueIsAbsent
	
}

// NewHashTable - конструкто хеш-таблицы
func NewHashTable() *HashTable {
	return &HashTable{
		capacity: hashTableCapacity, // Размер таблицы
		table: make([]*node, hashTableCapacity), // Сама таблица
	}
}

func main() {
	// Создаем сканер
	scanner := makeScanner()
	// Считываем с ввода количество команд
	commandsNumber := readInt(scanner)
	// Создаем новую хеш-таблицу
	hashTable := NewHashTable()
	// Переменная для результата выполнения команды
	result := ""
	// Создаем писатель для вывода
	writer := bufio.NewWriter(os.Stdout)
	// В цикле по считанному количеству читаем команды с ввода, выполняем, результат кладем в писатель
	for i := 0; i < commandsNumber; i++ {
		// readCommand - читаем команду с ввода сканером
		// executeCommand - выполняем команду, получаем результат
		result = executeCommand(hashTable, readCommand(scanner))
		// Выводим только непустые результаты
		if result != "" {
			writer.WriteString(result)
			writer.WriteString("\n")
		}
	}
	writer.Flush()
}

// makeScanner - создает новый сканер
func makeScanner() *bufio.Scanner {
	const maxCapacity = 10 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

// readInt - читает целое число сканером с ввода
func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

// readCommand - читает команду
func readCommand(scanner *bufio.Scanner) сommand {
	// Создаем новую структур команды
	command := сommand{}
	// Читаем строку команды
	scanner.Scan()
	// Разбиваем строку команды на имя, ключ и значение
	listCommand := strings.Split(scanner.Text(), " ")
	// Проверяем, что получили
	if len(listCommand) >= 2 {
		// Имя команды и ключ берем из первого и второго элементов
		command.name = listCommand[0]
		command.key, _ = strconv.Atoi(listCommand[1])
	}
	// Проверяем, есть ли третий элемент - в нем должно лежать значение
	if len(listCommand) == 3 {
		// Третий элемент есть, берем его
		command.value, _ = strconv.Atoi(listCommand[2])
	}
	// Возвращаем структур команды
	return command
}

func executeCommand(hashTable *HashTable, command сommand) string {
	// Служебные переменные
	var err error // Ошибка выполнения команды
	var value int // Результат выполнения команды
	var hasValue bool // Признак того, что выполнялась команда, возвращающая значение
	// По переданному имени команды вызываем метод хеш-таблицы
	switch command.name {
	case "put":
		// Добавление значения по ключу
		hashTable.put(command.key, command.value)
	case "get":
		// Получение значения по ключу
		value, err = hashTable.get(command.key)
		// Команда, возвращающая значение
		hasValue = true
	case "delete":
		// Получение значения по ключу с удалением
		value, err = hashTable.delete(command.key)
		// Команда, возвращающая значение
		hasValue = true
	}
	// Проверим, есть ли ошибка - все равно, какая команда выполнялась
	if err != nil {
		// Ошибка есть, возвращаем "None"
		return "None"
	}
	// Проверим, выполнялась ли команда, возвращающая значение
	if hasValue {
		// Выполнялась, вернем значение, преобразованное к строке
		return strconv.Itoa(value)
	}
	// Без ошибок выполнилась команда, которая не возвращает значение
	return ""
}