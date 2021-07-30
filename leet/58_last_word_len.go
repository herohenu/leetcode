package leet

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	l := len(arr)
	if l == 0 {
		return 0
	}
	return len(arr[l-1])
}

func LengthOfLastWord(s string) int {
	return lengthOfLastWord(s)
}
