package models

type CanRep struct {
	list map[int]bool
}

func NewCanRep() *CanRep {
	return &CanRep{
		list: make(map[int]bool),
	}
}

func (rep *CanRep) Clear() {
	for k := range rep.list {
		delete(rep.list, k)
	}
}

func (rep *CanRep) Add(value int) {
	rep.list[value] = true
}

func (rep *CanRep) Remove(value int) {
	delete(rep.list, value)
}

func (rep *CanRep) DoesContain(value int) bool {
	_, ok := rep.list[value]
	return ok
}

func (rep *CanRep) GetFirst() int {
	if len(rep.list) > 0 {
		for c := range rep.list {
			return c
		}
	}
	return 0
}

func (rep *CanRep) ContainsSolution() bool {
	return len(rep.list) == 1
}

func (rep *CanRep) GetAll() []int {
	result := make([]int, 0)
	for v := range rep.list {
		result = append(result, v)
	}
	return result
}

func (rep *CanRep) SetAll(list []int) {
	rep.Clear()
	for _, v := range list {
		rep.Add(v)
	}
}
