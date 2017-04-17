package ptr_unsafe

import (
	"fmt"
	"unsafe"
)

type Empty struct {
}

func main() {
	fmt.Println(unsafe.Sizeof(Empty{})) // Output: 0
	var numbs [5]int = [...]int{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(numbs)) // Output: 40 (8*5)
	p := unsafe.Pointer(&numbs)
	fmt.Println(p) //0xc4200142d0
	fmt.Println(uintptr(p))

	p = unsafe.Pointer(uintptr(p) + 8) // offset 8 (8*1)
	pint := (*int)(p)
	fmt.Println(*pint) // Output: 2

	data := []byte("李谊恒")
	slicePtr := unsafe.Pointer(&data)
	strPtr := (*string)(slicePtr)
	fmt.Println(*strPtr) // Output: 李谊恒

}
