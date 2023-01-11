package linked_list

import (
	"fmt"
	"strings"
)

type Node struct {
	Key   string
	Value int
	Next  *Node
}

func NewNode(key string, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Next:  nil,
	}
}

func (n *Node) UpdateNext(node *Node) {
	n.Next = node
}

func (n Node) ToString() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[ %s: %d ]", n.Key, n.Value))
	return sb.String()
}
