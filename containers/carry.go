package containers

import "sort"

type Payload int

type ElementaryCarry struct {
	id		int
	value	Payload
}

func NewElementaryCarry(id int, val Payload) ElementaryCarry {
	return ElementaryCarry{
		id:id,
		value:val,
	}
}

type Carry struct {
	Id    int
	value []ElementaryCarry
}

func (carry *Carry) Size() int {
	return len(carry.value)
}

func (carry *Carry) Append(elem ElementaryCarry) {
	carry.value = append(carry.value, elem)
}

func NewCarries(args ...ElementaryCarry) []Carry {
	result := make([]Carry, len(args))
	for ind := 0; ind < len(args); ind++ {
		result[ind] = NewCarry(ind, args[ind])
	}

	return result
}

func NewCarriesN(number int) []Carry {
	result := make([]Carry, number)
	for ind := 0; ind < number; ind++ {
		result[ind] = NewCarry(ind, NewElementaryCarry(ind, Payload(number+1)))
	}

	return result
}

func NewCarry(id int, val ElementaryCarry) Carry {
	return Carry{
		Id:id,
		value:[]ElementaryCarry{val},
	}
}

func (carry *Carry) Equal(otherCarry Carry) bool {
	return carry.Id == otherCarry.Id
}

func (carry *Carry) NotEqual(otherCarry Carry) bool {
	return !carry.Equal(otherCarry)
}

type CarriesSet struct {
	StepId		StepId
	arrCarry	[]Carry
}

//type CarriesSet []Carry

func NewCarriesSet(args ...Carry) CarriesSet {
	ptr := new(CarriesSet)
	for _, val := range args {
		ptr.arrCarry = append(ptr.arrCarry, val)
	}

	sort.Sort(ById(*ptr))
	return *ptr
}

func (set *CarriesSet) Equal(otherSet CarriesSet) bool {
	if set.Size() != otherSet.Size() {
		return false
	}

	for ind := 0; ind < set.Size(); ind++ {
		if set.arrCarry[ind].NotEqual(otherSet.arrCarry[ind]) {
			return false
		}
	}

	return true
}

func (set CarriesSet) NotEqual(otherSet CarriesSet) bool {
	return !set.Equal(otherSet)
}

func (set *CarriesSet) AddSet(otherSet CarriesSet) {
	for _, val := range otherSet.arrCarry {
		set.Append(val)
	}
}

type ById []Carry

func (seq ById) Len() int {
	return len(seq)
}

func (seq ById) Less(i, j int) bool {
	return seq[i].Id < seq[j].Id
}

func (seq ById) Swap(i, j int) {
	seq[i], seq[j] = seq[j], seq[i]
}

func (set *CarriesSet) Consist(carry Carry) bool {
	for ind := 0; ind < set.Size(); ind++ {
		if set.arrCarry[ind].Equal(carry) {
			return true
		}
	}

	return false
}

func (set *CarriesSet) Append(carry Carry) {
	if set.Consist(carry) {
		return
	}

	set.arrCarry = append(set.arrCarry, carry)
	sort.Sort(ById(*set))
}

func (set CarriesSet) Size() int {
	return len(set.arrCarry)
}

func (set CarriesSet) Get(ind int) Carry {
	return set.arrCarry[ind]
}

func (set *CarriesSet) Clear() {
	set.arrCarry = make([]Carry, 0)
}
