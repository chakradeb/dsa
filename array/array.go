package array

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

func (a Array) Get(index int) int {
	return a.Data[index]
}

func (a *Array) RemoveLast() int {
	last := a.Data[a.Length-1]
	delete(a.Data, a.Length-1)
	a.Length--
	return last
}

func (a *Array) Remove(index int) int {
	element := a.Data[index]
	for index+1 < a.Length {
		a.Data[index] = a.Data[index+1]
		index++
	}
	a.RemoveLast()
	return element
}
