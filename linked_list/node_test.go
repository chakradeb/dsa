package linked_list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_ToString(t *testing.T) {
	key := "apple"
	value := 32
	n := NewNode(key, value)
	assert.Equal(t, fmt.Sprintf("[ %s: %d ]---> nil", key, value), n.ToString(), "unexpected toString representation")
}

func TestNode_UpdateNext(t *testing.T) {
	n := NewNode("apple", 32)
	assert.Nil(t, n.Next, "unexpected next")
	n2 := NewNode("banana", 28)
	n.UpdateNext(n2)
	assert.NotNil(t, n.Next, "unexpected next")
	assert.Equal(t, n2, n.Next, "unexpected next")
}
