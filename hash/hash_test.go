package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/chakradeb/dsa/linked_list"
)

func TestHash_Add(t *testing.T) {
	hasher = func(s string, limit int) int {
		return 1
	}
	n := ll.NewNode("apple", 35)
	l := ll.NewLinkedList()
	l.Append(n)

	h := NewHash(3)
	assert.Zero(t, h.Length, "unexpected length")
	assert.True(t, h.Table[0].IsEmpty(), "unexpected key value pair")
	assert.True(t, h.Table[1].IsEmpty(), "unexpected key value pair")
	assert.True(t, h.Table[2].IsEmpty(), "unexpected key value pair")
	err := h.Add("apple", 35)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 1, h.Length, "unexpected length")
	assert.True(t, h.Table[0].IsEmpty(), "unexpected key value pair")
	assert.Equal(t, l, h.Table[1], "unexpected key value pair")
	assert.True(t, h.Table[2].IsEmpty(), "unexpected key value pair")
}

func TestHash_Add_WhenCollision(t *testing.T) {
	hasher = func(s string, limit int) int {
		return 1
	}

	n1 := ll.NewNode("apple", 35)
	n2 := ll.NewNode("banana", 28)
	l := ll.NewLinkedList()
	l.Append(n1)
	l.Append(n2)

	h := NewHash(3)
	assert.Zero(t, h.Length, "unexpected length")

	err := h.Add("apple", 35)
	assert.NoError(t, err, "unexpected error")

	err = h.Add("banana", 28)
	assert.NoError(t, err, "unexpected error")
	assert.True(t, h.Table[0].IsEmpty(), "unexpected key value pair")
	assert.True(t, h.Table[2].IsEmpty(), "unexpected key value pair")
	assert.False(t, h.Table[1].IsEmpty(), "unexpected key value pair")
	assert.Equal(t, l.ToString(), h.Table[1].ToString(), "unexpected key value pair")
}

func TestHash_Add_WhenExceedCapacity(t *testing.T) {
	hasher = func(s string, limit int) int {
		return 0
	}

	h := NewHash(1)
	err := h.Add("apple", 35)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 1, h.Length, "unexpected length")

	err = h.Add("banana", 28)
	assert.EqualError(t, err, "hashmap is already at full capacity")
	assert.Equal(t, 1, h.Length, "unexpected length")
}

func TestHash_Remove(t *testing.T) {
	hasher = func(s string, limit int) int {
		return 1
	}

	n1 := ll.NewNode("apple", 35)
	n2 := ll.NewNode("cherry", 76)
	l := ll.NewLinkedList()
	l.Append(n1)
	l.Append(n2)

	h := NewHash(3)
	err := h.Add("apple", 35)
	assert.NoError(t, err, "unexpected error")

	err = h.Add("banana", 28)
	assert.NoError(t, err, "unexpected error")
	err = h.Add("cherry", 76)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 3, h.Length, "unexpected length")

	err = h.Remove("banana")
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 2, h.Length, "unexpected length")
	assert.Equal(t, l.ToString(), h.Table[1].ToString(), "unexpected key value pair")
}

func TestHash_Remove_WhenInvalidKey(t *testing.T) {
	hasher = func(s string, limit int) int {
		return 0
	}

	h := NewHash(1)
	err := h.Add("apple", 35)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 1, h.Length, "unexpected length")

	err = h.Remove("banana")
	assert.EqualError(t, err, "key banana not present in hashmap", "unexpected error")
	assert.Equal(t, 1, h.Length, "unexpected length")
}
