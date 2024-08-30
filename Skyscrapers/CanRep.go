package main

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

func (rep *CanRep) RemoveAbsent(other *CanRep) (done bool) {
	done = false
	for c := range rep.list {
		if !other.DoesContain(c) {
			rep.Remove(c)
			done = true
		}
	}
	return
}
