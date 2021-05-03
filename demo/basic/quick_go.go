package basic

import (
	"errors"
	"fmt"
	"unsafe"
)

func DemoIf() bool {
	fmt.Println(">>> DemoIf")

	tmp := map[string]string {
		"a":"a",
		"b":"b",
		"c":"c",
	}

	if ele, ok := tmp["t"]; !ok {
		fmt.Println("t not exist")
	} else {
		fmt.Println(ele)
	}
	if ele, ok := tmp["d"]; !ok {
		fmt.Println("d not exist")
	} else {
		fmt.Println(ele)
	}

	return true
}

type DemoS1 struct {
	x int
	y int
}

func DemoStruct() bool {
	fmt.Println(">>> DemoStruct")

	var t  = DemoS1 {
		x: 1,
		y: 2,
	}
	fmt.Println(t)
	return true
}

func DemoInterface() bool {
	fmt.Println(">>> DemoInterface")

	var t = DemoS1{}

	var a, b interface{}
	a = t
	fmt.Println(unsafe.Sizeof(a))
	b = &t
	fmt.Println(unsafe.Sizeof(b))
	var p = b.(*DemoS1)
	fmt.Println(unsafe.Sizeof(p))

	return true
}

type CyError struct {
	Err string
}

func (e CyError) Error() string {
	return e.Err
}

func DemoError() {
	fmt.Println(">>> DemoError")
	var err CyError = CyError{"a mistake"}
	fmt.Println(err)

	var err2 = errors.New("haha")
	fmt.Println(err2)
}

