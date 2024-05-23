package main

import (
	"fmt"
	"strings"
)

func main() {

}

// Node Узел двусвязного списка
type Node struct {
	data string
	prev *Node
	next *Node
}

// DualList Двусвязный список
type DualList struct {
	head *Node
	tail *Node
}

// NewDualList Возвращает новый пустой двусвязный список
func NewDualList() *DualList {

	return &DualList{}
}

// Append Добавление данных в список
func (dl *DualList) Append(data string) {

	newNode := &Node{
		data: data,
	}
	newNode.prev = dl.head

	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
		return
	}

	dl.head.next = newNode
	dl.head = newNode
}

// ToStringForward От начала до конца
func (dl *DualList) ToStringForward() string {

	if dl.tail == nil {
		return ""
	}

	var sb strings.Builder

	sb.WriteString(dl.tail.data)
	tmpNode := dl.tail.next

	for tmpNode != nil {

		tmpStr := fmt.Sprintf(" -> %s", tmpNode.data)
		sb.WriteString(tmpStr)
		tmpNode = tmpNode.next
	}

	return sb.String()
}

// ToStringBackward С конца до начала
func (dl *DualList) ToStringBackward() string {

	if dl.head == nil {
		return ""
	}

	var sb strings.Builder

	sb.WriteString(dl.head.data)
	tmpNode := dl.head.prev

	for tmpNode != nil {

		tmpStr := fmt.Sprintf(" -> %s", tmpNode.data)
		sb.WriteString(tmpStr)
		tmpNode = tmpNode.prev
	}

	return sb.String()
}

var ErrOutOfRangeIndex = fmt.Errorf("index out of range")
var ErrSearchInEmptyList = fmt.Errorf("search in empty list")

// DataByIndex Индексация с 0
func (dl *DualList) DataByIndex(idx int) (string, error) {

	if dl.tail == nil {
		return "", ErrSearchInEmptyList
	}

	idxCnt := 0
	tmpNode := dl.tail

	for idxCnt < idx {

		if tmpNode.next == nil {
			return "", ErrOutOfRangeIndex
		}
		tmpNode = tmpNode.next
		idxCnt++
	}

	return tmpNode.data, nil
}

func (dl *DualList) AppendWithIdx(data string, idx int) {

}
