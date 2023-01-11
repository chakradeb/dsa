package linked_list

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)

	ll := NewLinkedList(node1)
	assert.Equal(t, node1, ll.Head, "unexpected head")
	assert.Equal(t, node1, ll.Tail, "unexpected tail")
	assert.Equal(t, 1, ll.Length, "unexpected length")

	ll.Append(node2)
	assert.Equal(t, 2, ll.Length, "unexpected length")
	assert.Equal(t, node1.ToString(), ll.Head.ToString(), "unexpected key for head")
	assert.Equal(t, node2.ToString(), ll.Tail.ToString(), "unexpected key for tail")
	assert.Equal(t, node2.ToString(), ll.Head.Next.ToString(), "unexpected key for next to head")
	assert.Nil(t, ll.Tail.Next, "unexpected next for tail")
}

func TestLinkedList_Prepend(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	ll.Prepend(node2)
	assert.Equal(t, 2, ll.Length, "unexpected length")
	assert.Equal(t, node2.ToString(), ll.Head.ToString(), "unexpected key for head")
	assert.Equal(t, node1.ToString(), ll.Tail.ToString(), "unexpected key for tail")
	assert.Equal(t, node1.ToString(), ll.Head.Next.ToString(), "unexpected key for next to head")
	assert.Nil(t, ll.Tail.Next, "unexpected next for tail")
}

func TestLinkedList_Insert(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)
	node3 := NewNode("cherry", 76)

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	ll.Append(node3)
	assert.Equal(t, 2, ll.Length, "unexpected length")

	err := ll.Insert(1, node2)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 3, ll.Length, "unexpected length")
	assert.Equal(t, node1.ToString(), ll.Head.ToString(), "unexpected key for head")
	assert.Equal(t, node3.ToString(), ll.Tail.ToString(), "unexpected key for tail")
	assert.Equal(t, node2.ToString(), ll.Head.Next.ToString(), "unexpected key for next to head")
	assert.Nil(t, ll.Tail.Next, "unexpected next for tail")
}

func TestLinkedList_Insert_IndexZero(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	err := ll.Insert(0, node2)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 2, ll.Length, "unexpected length")
	assert.Equal(t, node2.ToString(), ll.Head.ToString(), "unexpected key for head")
	assert.Equal(t, node1.ToString(), ll.Tail.ToString(), "unexpected key for tail")
	assert.Equal(t, node1.ToString(), ll.Head.Next.ToString(), "unexpected key for next to head")
	assert.Nil(t, ll.Tail.Next, "unexpected next for tail")
}

func TestLinkedList_Insert_WhenInvalidIndex(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)

	ll := NewLinkedList(node1)
	err := ll.Insert(5, node2)
	assert.EqualError(t, err, "invalid index for current linked list", "unexpected error string")
	assert.Equal(t, 1, ll.Length, "unexpected length")
	assert.Equal(t, node1.ToString(), ll.Head.ToString(), "unexpected key for head")
	assert.Equal(t, node1.ToString(), ll.Tail.ToString(), "unexpected key for tail")
}

func TestLinkedList_Remove(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)
	node3 := NewNode("cherry", 76)

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	ll.Append(node2)
	ll.Append(node3)
	assert.Equal(t, 3, ll.Length, "unexpected length")

	err := ll.Remove(1)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 2, ll.Length, "unexpected length")
	assert.Equal(t, node1.ToString(), ll.Head.ToString(), "unexpected key for head")
	assert.Equal(t, node3.ToString(), ll.Tail.ToString(), "unexpected key for tail")
	assert.Equal(t, node3.ToString(), ll.Head.Next.ToString(), "unexpected key for next to head")
	assert.Nil(t, ll.Tail.Next, "unexpected next for tail")
}

func TestLinkedList_Remove_WhenFirstElement(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	ll.Append(node2)
	assert.Equal(t, 2, ll.Length, "unexpected length")
	assert.Equal(t, node2.ToString(), ll.Tail.ToString(), "unexpected key for tail")

	err := ll.Remove(0)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 1, ll.Length, "unexpected length")
	assert.Equal(t, node2.ToString(), ll.Head.ToString(), "unexpected key for head")
	assert.Equal(t, node2.ToString(), ll.Tail.ToString(), "unexpected key for tail")
	assert.Nil(t, ll.Head.Next, "unexpected key for next to head")
	assert.Nil(t, ll.Tail.Next, "unexpected next for tail")
}

func TestLinkedList_Remove_WhenLastElement(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	ll.Append(node2)
	assert.Equal(t, 2, ll.Length, "unexpected length")
	assert.Equal(t, node2.ToString(), ll.Tail.ToString(), "unexpected key for tail")

	err := ll.Remove(1)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 1, ll.Length, "unexpected length")
	assert.Equal(t, node1.ToString(), ll.Head.ToString(), "unexpected key for head")
	assert.Equal(t, node1.ToString(), ll.Tail.ToString(), "unexpected key for tail")
	assert.Nil(t, ll.Head.Next, "unexpected key for next to head")
	assert.Nil(t, ll.Tail.Next, "unexpected next for tail")
}

func TestLinkedList_Remove_WhenInvalidIndex(t *testing.T) {
	node1 := NewNode("apple", 32)

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	err := ll.Remove(4)
	assert.EqualError(t, err, "invalid index for current linked list", "unexpected error")
	assert.Equal(t, 1, ll.Length, "unexpected length")
}

func TestLinkedList_ToString(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)
	var sb strings.Builder

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	ll.Append(node2)
	assert.Equal(t, 2, ll.Length, "unexpected length")

	sb.WriteString(node1.ToString())
	sb.WriteString(node2.ToString())

	assert.True(t, strings.Contains(ll.ToString(), sb.String()), "unexpected to string value")
}

func TestLinkedList_traverseToIndex(t *testing.T) {
	node1 := NewNode("apple", 32)
	node2 := NewNode("banana", 28)
	node3 := NewNode("cherry", 76)

	ll := NewLinkedList(node1)
	assert.Equal(t, 1, ll.Length, "unexpected length")

	ll.Append(node2)
	ll.Append(node3)
	assert.Equal(t, 3, ll.Length, "unexpected length")

	expected := ll.traverseToIndex(1)
	assert.Equal(t, node2.ToString(), expected.ToString(), "unexpected node in traversal")
}
