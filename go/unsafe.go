// unsafe 可以突破指针的限制

package learn

import (
	"unsafe"
)

// Student is a student
type Student struct {
	name string
	age  int
}

// UnsafeChangeStructValue 修改结构体的值
func UnsafeChangeStructValue() Student {

	s := Student{"shiniao", 25}
	// 获取name 地址
	name := (*string)(unsafe.Pointer(&s))
	*name = "chaojie"

	// 获取age的地址，并修改
	// unsafe.Offsetof获取 s.age 的偏移量，加上 s 的地址，就是 age 的地址
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.age)))
	*age = 12

	return s

}
