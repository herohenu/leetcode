package leet

import (
	"container/list"
)

func isValid(s string) bool {
	list := list.New()
	for idx, v := range s {
		char := string(v)
		if char == "{" || char == "[" || char == "(" {
			list.PushBack(char)
		}

		if char == "}" {
			e := list.Back()
			if e == nil {
				return false
			}
			if e.Value == "{" {
				list.Remove(e)
			} else {
				return false
			}
		}

		if char == "]" {
			e := list.Back()
			if e == nil {
				return false
			}
			if e.Value == "[" {
				list.Remove(e)
			} else {
				return false
			}
		}
		if char == ")" {
			e := list.Back()
			if e == nil {
				return false
			}
			if e.Value == "(" {
				list.Remove(e)
			} else {
				return false
			}
		}

		if idx == len(s)-1 && list.Len() != 0 {
			return false
		}

	}

	return true
}
