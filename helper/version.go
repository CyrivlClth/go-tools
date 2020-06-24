package helper

import (
	"strconv"
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
		l, _ := strconv.Atoi(s1[i])
		m, _ := strconv.Atoi(s2[i])
		if l > m {
			return 1
		}
		if l < m {
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
