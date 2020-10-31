package main

import (
	"fmt"
	"strconv"
)

type node struct {
	val  int
	next *node
}

type List struct {
	firstNode *node
	lastNode  *node
}

func NewList() *List {
	return &List{}
}

func (l *List) Add(val int) {
	if l.firstNode == nil {
		l.firstNode = &node{val: val}
		l.lastNode = l.firstNode
		return
	}
	l.lastNode.next = &node{val: val}
	l.lastNode = l.lastNode.next
}

func (l *List) String() string {
	currentNode := l.firstNode
	res := ""
	for currentNode != nil {
		res += strconv.Itoa(currentNode.val) + ", "
		currentNode = currentNode.next
	}
	res = "LIST: [" + res[:len(res)-2] + "]"
	return res
}

// Функция определения зацикленности списка алгоритмом Флойда
func isCycledList(l *List) bool {
	// Если текущая node nil либо следующая за ней nil
	if l.firstNode == nil || l.firstNode.next == nil {
		// список не зациклен
		// завершаем сопрограмму
		return false
	}

	// инициализируем переменную для +1 итерации (наша черепашка)
	turtleNode := l.firstNode
	// инициализируем переменную для +2 итераций (зайчик)
	hareNode := l.firstNode.next

	// До тех пор пока переменная +1 итерации не равна nil выполняем:
	for turtleNode != nil {
		// если указатель первой переменной равен указателю 2 переменной:
		if turtleNode == hareNode {
			// список является зацикленным
			return true
		}
		// если следующее значение быстрой переменной равно nil или следующее>следующее значение nil:
		if hareNode.next == nil || hareNode.next.next == nil {
			// Завершаем функцию, список не является зацикленным
			return false
		}
		// Присваиваем в переменную для +1 итерации значение следующего узла
		turtleNode = turtleNode.next
		// Присваиваем в переменную для +2 итерации значение следующего узла за следующим
		hareNode = hareNode.next.next
	}
	// Цикл пройден, список не зациклен
	return false
}

func main() {
	l := &List{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.firstNode.next.next = l.firstNode.next
	fmt.Println(isCycledList(l))
}
