package linked_list

import (
	"errors"
	"fmt"
	"strings"
)

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkedList(node *Node) *LinkedList {
	return &LinkedList{
		Head:   node,
		Tail:   node,
		Length: 1,
	}
}

func (l *LinkedList) Append(node *Node) {
	l.Tail.Next = node
	l.Tail = node
	l.Length++
}

func (l *LinkedList) Prepend(node *Node) {
	node.UpdateNext(l.Head)
	l.Head = node
	l.Length++
}

func (l *LinkedList) Insert(index int, node *Node) error {
	if index < 0 || index >= l.Length {
		return errors.New("invalid index for current linked list")
	}
	if index == 0 {
		l.Prepend(node)
		return nil
	}
	curr := l.traverseToIndex(index - 1)
	node.Next = curr.Next
	curr.Next = node
	l.Length++
	return nil
}

func (l *LinkedList) Remove(index int) error {
	if index < 0 || index >= l.Length {
		return errors.New("invalid index for current linked list")
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}
	curr := l.traverseToIndex(index - 1)
	curr.Next = curr.Next.Next
	if index == l.Length-1 {
		l.Tail = curr
	}
	l.Length--
	return nil
}

func (l LinkedList) traverseToIndex(index int) *Node {
	curr := l.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr
}

func (l LinkedList) ToString() string {
	var sb strings.Builder
	sb.WriteString(
		fmt.Sprintf("{\n\tHead: %+v,\n\tTail: %+v,\n\tLength: %d,\n\tNodes: ",
			l.Head, l.Tail, l.Length,
		),
	)
	for curr := l.Head; curr.Next != nil; curr = curr.Next {
		sb.WriteString(curr.ToString())
	}
	sb.WriteString(l.Tail.ToString())
	sb.WriteString("\n}")
	return sb.String()
}
