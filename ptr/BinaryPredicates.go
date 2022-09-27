package ptr

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

// Returns true if both bytes arrays of size `n`, pointed by p1 & p2 are equal.
// ptr.CmpBitsAs[[n]uint8](p1,p2) is safe & static alternive for this function.
func CmpNBytes(n uintptr, p1, p2 unsafe.Pointer) bool {
	if n > 0 {
		n--
		for {
			if *(*uint8)(p1) != *(*uint8)(p2) {
				return false
			}
			if n == 0 {
				break
			}
			n--
			p1 = unsafe.Add(p1, 1)
			p2 = unsafe.Add(p2, 1)
		}
	}
	return true
}

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

/*
type integer interface {
	uint8 | int8 |
		uint16 | int16 |
		uint32 | int32 |
		uint64 | int64 |
		int | uint |
		uintptr
}
*/
