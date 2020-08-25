// package convert
// TODO: improve performance
// TODO: add test cases
// Some simple method of string slice
package convert

import (
	"encoding/json"
	"strconv"
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
