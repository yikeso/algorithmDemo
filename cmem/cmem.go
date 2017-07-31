package cmem

import "C"
import "unsafe"

func Alloc(size uintptr) *byte {
	return (*byte)(C.malloc(size))
}

func Free(ptr *byte) {
	C.free(unsafe.Pointer(ptr))
}
