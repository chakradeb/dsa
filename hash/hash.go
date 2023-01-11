package hash

import (
	"errors"
	"fmt"
	ll "github.com/chakradeb/dsa/linked_list"
)

var hasher = func(s string, limit int) int {
	var sum int
	for _, char := range s {
		sum += int(char)
	}
	return sum % limit
}

type Hash struct {
	Table    []*ll.LinkedList
	Capacity int
	Length   int
}

func NewHash(capacity int) *Hash {
	var table []*ll.LinkedList
	for i := 0; i < capacity; i++ {
		table = append(table, ll.NewLinkedList())
	}
	return &Hash{
		Table:    table,
		Capacity: capacity,
		Length:   0,
	}
}

func (h *Hash) Add(key string, value int) error {
	if h.Length >= h.Capacity {
		return errors.New("hashmap is already at full capacity")
	}
	index := hasher(key, cap(h.Table))
	n := ll.NewNode(key, value)
	h.Table[index].Append(n)
	h.Length++
	return nil
}

func (h *Hash) Remove(key string) error {
	index := hasher(key, cap(h.Table))
	nodeIndex := h.Table[index].FindNodeBy(key)
	if nodeIndex == -1 {
		return fmt.Errorf("key %s not present in hashmap", key)
	}

	err := h.Table[index].Remove(nodeIndex)
	if err == nil {
		h.Length--
	}
	return err
}
