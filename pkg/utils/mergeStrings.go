package utils

import (
	"strconv"
	"strings"
)

func MergeString(s1 string, s2 string) string {
	result := strings.Builder{}
	result.WriteString(strconv.Itoa(len(s1)))
	result.WriteRune(' ')
	result.WriteString(s1)
	result.WriteString(s2)
	return result.String()
}

func UnmergeStrings(merged string) []string {
	s1Limit := strings.Builder{}
	i := 0
	for ; i < len(merged); i++ {
		if merged[i] == ' ' {
			i++
			break
		}
		s1Limit.WriteByte(merged[i])
	}

	s1Len, _ := strconv.Atoi(s1Limit.String())
	s1Len += i
	s1 := strings.Builder{}
	for ; i < s1Len; i++ {
		s1.WriteByte(merged[i])
	}
	result := []string{s1.String()}

	s2 := strings.Builder{}
	for ; i < len(merged); i++ {
		s2.WriteByte(merged[i])
	}
	return append(result, s2.String())
}
