package ptr

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

// Uses `unsafe.Pointer` & `unsafe.Add` to return
//   pointer `ptr` offset with `offset` .
func Offset[T any, I constraints.Integer](ptr *T, offset I) *T {
	return (*T)(
		unsafe.Add(
			unsafe.Pointer(ptr),
			offset,
		),
	)
}