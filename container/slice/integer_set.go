// package set
package slice

type integerSet Integer

func NewIntegerSet(s Integer) integerSet {
	return integerSet(s.Distinct())
}

func (t integerSet) Merge(s Integer) integerSet {
	return integerSet(Integer(append(t, s...)).Distinct())
}
