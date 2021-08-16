package handy

import (
	"unsafe"
)

const MaxUint = ^uint(0)
const MaxIny = int(MaxUint >> 1)

const UintptrSize = unsafe.Sizeof(uintptr(0))

func Is64() bool {
	return ^uint(0)>>63 == 1
}
