package main

import "fmt"

/*
@Time : 2021/8/6 10:25
*/

func ReverseString(s []byte) {
	fmt.Println("s :", s)
	r := len(s) - 1
	l := 0
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	fmt.Println("s :", s)
}
