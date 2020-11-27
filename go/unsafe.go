// unsafe 可以突出指针的限制

package learn

import (
	"unsafe"
)

type Student struct {
	name string
	age  int
}

// UnsafeChangeStructValue 修改结构体的值
func UnsafeChangeStructValue() Student {

	s := Student{"shiniao", 25}
	// 获取name 地址
	// name := (*string)(unsafe.Pointer(&s))

	// 获取age的地址，并修改
	// unsafe.Sizeof获取s.name的大小，加上s的地址，就是age的地址
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s.name)))
	*age = 12

	return s

}
