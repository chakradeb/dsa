package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray_Add(t *testing.T) {
	arr := NewArray()
	assert.Equal(t, 0, arr.Length, "unexpected length")
	arr.Add(5)
	assert.Equal(t, 1, arr.Length, "unexpected length")
	assert.Equal(t, 5, arr.Data[0], "unexpected element")
}

func TestArray_Get(t *testing.T) {
	arr := NewArray()
	arr.Add(5)
	arr.Add(9)
	assert.Equal(t, 5, arr.Get(0), "unexpected element")
	assert.Equal(t, 9, arr.Get(1), "unexpected element")
}

func TestArray_Remove(t *testing.T) {
	arr := NewArray()
	arr.Add(3)
	arr.Add(6)
	arr.Add(9)
	assert.Equal(t, 3, arr.Length, "unexpected length")
	assert.Equal(t, 6, arr.Get(1), "unexpected element")
	arr.Remove(1)
	assert.Equal(t, 2, arr.Length, "unexpected length")
	assert.Equal(t, 9, arr.Get(1), "unexpected element")
}

func TestArray_RemoveLast(t *testing.T) {
	arr := NewArray()
	arr.Add(3)
	arr.Add(6)
	arr.Add(9)
	assert.Equal(t, 3, arr.Length, "unexpected length")
	assert.Equal(t, 9, arr.Get(arr.Length-1), "unexpected element")
	arr.RemoveLast()
	assert.Equal(t, 2, arr.Length, "unexpected length")
	assert.Equal(t, 6, arr.Get(arr.Length-1), "unexpected element")
}
