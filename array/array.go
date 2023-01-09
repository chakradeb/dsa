package array

import (
	"errors"
	"math"
)

type Array struct {
	Length int
	Data   map[int]int
}

func NewArray() *Array {
	return &Array{
		Length: 0,
		Data:   make(map[int]int),
	}
}

func (a *Array) Add(element int) {
	a.Data[a.Length] = element
	a.Length++
}

func (a Array) Get(index int) (int, error) {
	if a.Length == 0 {
		return a.emptyArrayAlert()
	}
	if index >= a.Length {
		return a.invalidArrayIndexAlert()
	}
	return a.Data[index], nil
}

func (a *Array) RemoveLast() (int, error) {
	if a.Length == 0 {
		return a.emptyArrayAlert()
	}
	last := a.Data[a.Length-1]
	delete(a.Data, a.Length-1)
	a.Length--
	return last, nil
}

func (a *Array) Remove(index int) (int, error) {
	if a.Length == 0 {
		return a.emptyArrayAlert()
	}
	if index >= a.Length {
		return a.invalidArrayIndexAlert()
	}
	element := a.Data[index]
	for index+1 < a.Length {
		a.Data[index] = a.Data[index+1]
		index++
	}
	a.RemoveLast()
	return element, nil
}

func (a Array) emptyArrayAlert() (int, error) {
	return -math.MaxInt, errors.New("array is empty")
}

func (a Array) invalidArrayIndexAlert() (int, error) {
	return -math.MaxInt, errors.New("invalid array index")
}
