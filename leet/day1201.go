package leet

import (
	"fmt"
	"unsafe"
)

type Man struct {
}

func PrintStructLen() {
	m := Man{}
	fmt.Println(unsafe.Sizeof(m))
	fmt.Println(unsafe.Sizeof(Man{}))
}
