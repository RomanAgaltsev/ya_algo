/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 3
Задача - B. Эффективная быстрая сортировка

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/23815/run-report/112777191/

-- ПРИНЦИП РАБОТЫ --
Из постановки задачи сразу было понятно, что нужно сделать и как решать:
- На вход подается массив участников, у каждого три поля;
- Надо отсортировать быстрой сортировкой;
- Быстрая сортировка не должна использовать дополнительную память для промежуточных данных;
- Для сортировки потребуется функция-компаратор, так как участников надо сравнивать по трем параметрам.

Сначала читаем входные данные участников и складываем в массив структур.
Для хранения данных участника используется структура Participant.
Возможно, если бы решал не на Go, можно было бы использовать массив массивов.
Но Go не позволяет смешивать в массиве (слайсе) значения разных типов. Поэтому - массив структур.

Так как требуется отсортировать полностью все входные данные, то пришлось сразу всё прочитать в слайс.
Значит, весь слайс сразу был загружен в память.

Далее, слайс структур сортируется при помощи модифицированной быстрой сортировки, без использования дополнительнной памяти для промежуточных данных.

Функция сортировки quickSort от варианта с использованием дополнительной памяти отличается тем,
что в ней не используются промежуточные массивы данных для элементов меньше опорного, равного опорному и больше опорного.
Используя один исходный слайс и указатели left и right, мы как-будто "гуляем" указателями по этому одному слайсу, никуда его не выгружая и ничего не загружая в него.
Просто дробим диапазон индексов указателями и вызываем рекурсивно вложенные quickSort.

Вся "магия" происходит в функции partition:
- Выбирается опорный элемент;
- Дополнительные слайсы для элементов меньше, равных и больше опорного не создаются;
- Пока указатели left и right не встретились, во вложенных циклах они двигаются с концов слайса навстречу друг другу;
- При нахождении пары элементов, которые надо поменять местами, выполняется обмен;
- В конце возвращается указатель right.

Именно при движении указателей left и right в функции partition используется функция-компаратор less.
При движении указателей элементы слайса, на которые они указывают, сравниваются с опорным элементом.
Но так как в нашем слайсе не примитивные типа (число, строка), которые являются сравниваемымми, а структуры, требуется компаратор.

Принцип работы функции less следующий:
- Сначала сравниваются количества решенных задач;
- Потом сравниваются размеры штрафов;
- Далее сравниваются логины участников в лексикографическом порядке;
- Чем больше задач решил участник, тем меньшей должна быть его позиция в выходном списке.
	Участник считается меньше, если количество задач у него больше;
- Чем больше штрафа получил участник, тем большей должна быть его позиция в выходном списке.
	Участник считается больше, если размер штрафа у него больше (ну или наоборот - меньше/меньше);
- Если оба предыдущих параметра сравниваемых участников равны, всё решается сравнением логинов.
	По условию задачи логины уникальны, поэтому, совпадений быть не может - обязательно один участник будет меньше, другой больше.

Для сравнений количества задач и размеров штрафов использовал удобную функцию cmp.Compare - подсмотрел в пакете sort.
Её удобство в том, что для сравнения при решении задачи, необходимо определять три результата сравнения - больше, меньше и равно.
cmp.Compare как раз возвращает нужные результаты в виде -1/0/+1.

Вот и всё решение.
Рекурсивные вызовы дойдя до базового случая прекращаются, получаем слайс структур, которые отсортированы быстрой сортировкой по логике функции-компаратора less.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Могу выделить следующие критерии корректности решения:
1. Решение возвращает список логинов, отсортированных по требуемой логике сравнения участников - так и есть, тесты на Контесте пройдены;
2. Решение отрабатывает в рамках ограничений по времени и памяти - так и есть, опять же, тесты на Контесте пройдены;
3. В решении используется быстрая сортировка (Quicksort) - так и есть, используются функции quickSort и partition, реализующие быструю сортировку;
4. Быстрая сортировка, используемая в решении, не должна потреблять дополнительную память для промежуточных данных - так и есть, решение оперирует исходным слайсом и указателями, дополнительные слайсы не создаются.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Чтение входящего массива, его обработка до поиска и вывод результата не оценивается.

Фактически, решение задачи, это быстрая сортировка, только без использования дополнительной памяти для промежуточных данных.
Поэтому, в худшем случае - если массив уже отсортирован или при неудачном выборе опорного элемента (хотя, в решении первый и последний не выбираются) - это O(n^2).
В среднем случае (так и будем считать) - это O(n * log n) - O(n) на операции обработки элементов слайса и O(log n) на рекурсивные вызовы.
Возьму на себя смелость сказать, что сравнение структур-участников - это O(1).

В итоге, считаю, что решение работает за O(n * log n), свойственное среднему случаю быстрой сортировки.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Чтение входящего массива, его обработка до поиска и вывод результата не оценивается.

В решении используется память для:
- Массив (слайс) - O(n);
- Стек рекурсивных вызовов быстрой сортировки - O(log n);
- Несколько переменных в коде - O(1).

В итоге, имеем O(n + log n) всего. Из них O(log n) дополнительной памяти.
*/

package main

import (
	"bufio"
	"cmp"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Participant - структура для хранения данных участника
type Participant struct {
	login string // Логин участника
	solve int // Количество решенных задач
	penalty int // Штраф
}

// less - компаратор, сравнивает двух участников в соответствии с условиями задачи:
// - сначала по количеству решенных задач, больше задач - меньше позиция участника
// - затем по размеру штрафа, больше штраф - больше позиция участника
// - в конце по логину в лексикографическом порядке
// Возвращает:
//	-1, если участник "a" меньше (меньше позиция в списке) участника "b"
//	+1, если участник "a" больше (больше позиция в списке) участника "b"
//	0, если равны, что маловероятно...
func less(a, b Participant) int {
	// Больше задач - меньше позиция участника
	if res := cmp.Compare(b.solve, a.solve); res != 0 {
		return res
	}
	// Больше штраф - больше позиция участника
	if res := cmp.Compare(a.penalty, b.penalty); res != 0 {
		return res
	}
	// Логины просто сравниваются
	return strings.Compare(a.login, b.login)
}

// partition - разбивает переданный слайс участников в рамках переданных границ left и right,
//	с определением опорного элемента и раскидыванием остальных элементов влево и вправо от опорного
func partition(participants []Participant, left, right int) int {
	// Определяем индекс опорного элемента - отдельной строкой, чтобы удобнее было в отладчике смотреть его значение
	pivotIndex := rand.Intn(right + 1 - left) + left
	// Сам опорный элемент
	pivot := participants[pivotIndex]
	// Двигаем указатели left и right, пока они не встретились
	for left < right {
		// Двигаем вправо левый указатель, пока левый элемент меньше опорного
		for less(participants[left], pivot) == -1 {
			left++
		}
		// Двигаем влево правый указатель, пока опорный элемент, меньше правого
		for less(pivot, participants[right]) == -1 {
			right--
		}
		// Если указатели не сошлись в один, меняем местами участников, на которых смотрят left и right
		// Если указатели сошлись, менять элемент сам на себя нет смысла
		if left < right {
			participants[left], participants[right] = participants[right], participants[left]
		}
	}
	// Возвращаем правый указатель
	return right
}

// quickSort - сортирует участников при помощи алгоритма быстрой сортировки
func quickSort(participants []Participant, left, right int) {
	// Проверяем на базовый случай
	if left >= right {
		// Базовый случай, возврат
		return
	}
	// Вызываем разбивку диапазона left-right с перекидыванием участников влево-вправо от опорного
	// Возвращается индекс участника, который уже на "своем" месте
	p := partition(participants, left, right)
	// Проваливаемся в быструю сортировку диапазона от left, до элемента перед p
	quickSort(participants, left, p-1)
	// Проваливаемся в быструю сортировку диапазона от элемента после p, до right
	quickSort(participants, p+1, right)
}

func main() {
	// Создаем сканер
	scanner := makeScanner()
	// Читаем количество участников
	n := readInt(scanner)
	// Читаем участников в слайс структур
	participants := readParticipants(scanner, n)
	// Сортируем слайс участников
	quickSort(participants, 0, len(participants)-1)
	// Выводим участников после сортировки - только логины
	printParticipants(participants)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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
func readParticipants(scanner *bufio.Scanner, n int) []Participant {
	participants := make([]Participant, n)
	for i := 0; i < n; i++ {
		participants[i] = readParticipant(scanner)
	}
	return participants
}

func readParticipant(scanner *bufio.Scanner) Participant {
	scanner.Scan()
	listParticipant := strings.Split(scanner.Text(), " ")
	login := listParticipant[0]
	solve, _ := strconv.Atoi(listParticipant[1]) 
	penalty, _ := strconv.Atoi(listParticipant[2])
	return Participant{
		login,
		solve,
		penalty,
	}
}

func printParticipants(participants []Participant) {
	writer := bufio.NewWriter(os.Stdout)
	for _, participant := range participants {
		writer.WriteString(participant.login)
		writer.WriteString("\n")
	}
	writer.Flush()
}