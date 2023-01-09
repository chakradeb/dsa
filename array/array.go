package array

type Array struct {
	Length int64
	Data   map[int64]int64
}

func NewArray() *Array {
	return &Array{
		Length: 0,
		Data:   make(map[int64]int64),
	}
}

func (a *Array) Add(element int64) {
	a.Data[a.Length] = element
	a.Length++
}

func (a Array) Get(index int64) int64 {
	return a.Data[index]
}

func (a *Array) RemoveLast() int64 {
	last := a.Data[a.Length-1]
	delete(a.Data, a.Length-1)
	a.Length--
	return last
}

func (a *Array) Remove(index int64) int64 {
	element := a.Data[index]
	for index+1 < a.Length {
		a.Data[index] = a.Data[index+1]
		index++
	}
	a.RemoveLast()
	return element
}
