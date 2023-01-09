package array

import (
	"math"
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
	actual, err := arr.Get(0)
	assert.Equal(t, 5, actual, "unexpected element")
	assert.NoError(t, err, "unexpected error")
}

func TestArray_Get_WhenEmptyArray(t *testing.T) {
	arr := NewArray()
	actual, err := arr.Get(0)
	assert.Equal(t, -math.MaxInt, actual, "invalid element")
	assert.EqualError(t, err, "array is empty", "invalid error message")
}

func TestArray_Get_WhenInvalidIndex(t *testing.T) {
	arr := NewArray()
	arr.Add(5)
	actual, err := arr.Get(3)
	assert.Equal(t, -math.MaxInt, actual, "invalid element")
	assert.EqualError(t, err, "invalid array index", "invalid error message")
}

func TestArray_Remove(t *testing.T) {
	arr := NewArray()
	arr.Add(3)
	arr.Add(6)
	arr.Add(9)
	assert.Equal(t, 3, arr.Length, "unexpected length")
	assert.Equal(t, 6, arr.Data[1], "unexpected element")
	element, err := arr.Remove(1)
	assert.Equal(t, 2, arr.Length, "unexpected length")
	assert.Equal(t, 6, element, "unexpected length")
	assert.Equal(t, 9, arr.Data[1], "unexpected element")
	assert.NoError(t, err, "unexpected error")
}

func TestArray_Remove_WhenEmptyArray(t *testing.T) {
	arr := NewArray()
	assert.Equal(t, 0, arr.Length, "unexpected length")
	element, err := arr.Remove(1)
	assert.Equal(t, 0, arr.Length, "unexpected length")
	assert.Equal(t, -math.MaxInt, element, "unexpected length")
	assert.EqualError(t, err, "array is empty", "unexpected error")
}

func TestArray_Remove_WhenInvalidIndex(t *testing.T) {
	arr := NewArray()
	arr.Add(3)
	assert.Equal(t, 1, arr.Length, "unexpected length")
	assert.Equal(t, 3, arr.Data[0], "unexpected element")
	element, err := arr.Remove(5)
	assert.Equal(t, 1, arr.Length, "unexpected length")
	assert.Equal(t, -math.MaxInt, element, "unexpected length")
	assert.EqualError(t, err, "invalid array index", "unexpected error")
}

func TestArray_RemoveLast(t *testing.T) {
	arr := NewArray()
	arr.Add(3)
	arr.Add(6)
	arr.Add(9)
	assert.Equal(t, 3, arr.Length, "unexpected length")
	assert.Equal(t, 9, arr.Data[arr.Length-1], "unexpected element")
	actual, err := arr.RemoveLast()
	assert.Equal(t, 2, arr.Length, "unexpected length")
	assert.Equal(t, 9, actual, "unexpected length")
	assert.Equal(t, 6, arr.Data[arr.Length-1], "unexpected element")
	assert.NoError(t, err, "unexpected error")
}

func TestArray_RemoveLast_WhenEmptyArray(t *testing.T) {
	arr := NewArray()
	assert.Equal(t, 0, arr.Length, "unexpected length")
	actual, err := arr.RemoveLast()
	assert.Equal(t, -math.MaxInt, actual, "unexpected length")
	assert.EqualError(t, err, "array is empty", "unexpected error")
}
