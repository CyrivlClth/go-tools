// package set
package set

import "github.com/CyrivlClth/go-tools/container/slice"

type integerSet slice.Integer

func NewIntegerSet(s slice.Integer) integerSet {
	return integerSet(s.Distinct())
}

func (t integerSet) Merge(s slice.Integer) integerSet {
	return integerSet(slice.Integer(append(t, s...)).Distinct())
}
