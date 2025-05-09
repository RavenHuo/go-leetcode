package main

import (
	"fmt"
	"unsafe"
)

func main() {
	struct1 := struct{}{}
	struct2 := struct{}{}
	fmt.Printf("struct1:%d struct2:%d, struct1P:%p, struct2P:%p\n", unsafe.Sizeof(struct1), unsafe.Sizeof(struct2), &struct1, &struct2)
	fmt.Printf("%+v", &struct1 == &struct2)
	// struct1:0 struct2:0, struct1P:0x8083e0, struct2P:0x8083e0
}
