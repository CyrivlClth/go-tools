package helper

import (
	"strings"
)

func VersionCompare(v1, v2 string) int {
	v1 = strings.TrimPrefix(v1, "v")
	v2 = strings.TrimPrefix(v2, "v")
	s1 := strings.Split(v1, ".")
	s2 := strings.Split(v2, ".")
	ml := len(s1)
	if ml > len(s2) {
		ml = len(s2)
	}
	for i := 0; i < ml; i++ {
		if s1[i] > s2[i] {
			return 1
		}
		if s1[i] < s2[i] {
			return -1
		}
	}
	if len(s1) > len(s2) {
		return 1
	}
	if len(s1) < len(s2) {
		return -1
	}
	return 0
}
