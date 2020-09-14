// package convert
// TODO: improve performance
// TODO: add test cases
// Some simple method of string slice
package convert

import (
	"encoding/json"
	"strconv"
	"strings"
	"unsafe"
)

// ToJSONString convert object to json string
// Example:
// []int{1,2,3} --> "[1,2,3]"
func ToJSONString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return *((*string)(unsafe.Pointer(&bytes)))
}

// StrToIntSlice string slice convert to int slice
// Example:
// []string{"1","2","3"} --> []int{1,2,3}
func StrToIntSlice(s []string) ([]int, error) {
	r := make([]int, len(s))
	for i, t := range s {
		atoi, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		r[i] = atoi
	}

	return r, nil
}

// IntToStrSlice int slice to string slice
// Example:
// []int{1,2,3} --> []string{"1","2","3"}
func IntToStrSlice(s []int) []string {
	z := make([]string, len(s))
	for i, j := range s {
		z[i] = strconv.Itoa(j)
	}

	return z
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func camelToSnake(s string) string {
	var b strings.Builder
	b.Grow(len(s) * 2)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '_' || asciiSpace[c] == 1 {
			continue
		}

		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
			if i != 0 {
				b.WriteByte('_')
			}
		}
		b.WriteByte(c)
	}

	return b.String()
}

func snakeToCamel(s string, first bool) string {
	var b strings.Builder
	b.Grow(len(s))
	n := first
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '_' || asciiSpace[c] == 1 {
			n = true
			continue
		}
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if n {
			c -= 'a' - 'A'
			n = false
		}
		b.WriteByte(c)
	}

	return b.String()
}
