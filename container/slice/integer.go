// package slice
package slice

type Integer []int

// Distinct distinct elem in slice
// only keep different elem in slice
// it will keep order of origin
// WARNING: this method will change origin slice sort and return part of it
//
// Example:
//     Integer([]int{2,2,1,3}).Distinct()
// >>> []int{2,1,3}
// origin slice: []int{2,1,3,3}
func (s Integer) Distinct() Integer {
	if len(s) == 0 {
		return s
	}

	m := make(map[int]bool)
	for _, u := range s {
		m[u] = true
	}
	l := len(m)
	i := 0
	for _, u := range s {
		if m[u] {
			s[i] = u
			m[u] = false
			i++
		}
		if i >= l {
			break
		}
	}

	return s[:i]
}

// Contains check if slice contains elem
func (s Integer) Contains(x int) bool {
	return contains(len(s), func(i int) bool { return s[i] == x })
}

func contains(n int, fn func(i int) bool) bool {
	for i := 0; i < n; i++ {
		if fn(i) {
			return true
		}
	}
	return false
}

// Filter filter elem in slice
// if fn(x) is true, it will keep it.
// Example:
//     Integer([]int{1,2,3,4}).Filter(func(x int) bool {return x%2==0})
// >>> []int{2,4}
func (s Integer) Filter(fn func(x int) bool) Integer {
	i := 0
	for _, it := range s {
		if fn(it) {
			s[i] = it
			i++
		}
	}

	return s[:i]
}
