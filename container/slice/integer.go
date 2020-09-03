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
