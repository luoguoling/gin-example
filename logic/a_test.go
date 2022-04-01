package logic

import (
	"fmt"
	"testing"
	"unsafe"
)

type A struct {
	a int
}

func TestA(t *testing.T) {
	a := A{}
	fmt.Println(unsafe.Sizeof(a))
}
