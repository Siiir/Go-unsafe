package ptr

import (
	"unsafe"
)

/*
	Returns true if both bytes arrays of size `n`, pointed by p1 & p2 are equal.

	Alternatives:

	Function CmpBitsAs[[n]uint8](p1,p2) from https://github.com/Siiir/ptr v1.0.0,
	is a safe & more compile-time alternative for this unsafe one.
*/
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
