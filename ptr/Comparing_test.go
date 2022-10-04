package ptr

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/Siiir/asserter/v2"
)

func TestCmpNBytes(t *testing.T) {
	const tcQuantity = 7
	argTab := [tcQuantity]struct {
		a [2][]byte
		n uintptr
	}{
		{[2][]byte{[]byte("a\bcde\f"), []byte("abcde")}, 4}, // 0
		{[2][]byte{[]byte("abc"), []byte("def")}, 3},
		{[2][]byte{[]byte("XHKr"), []byte("XHKr")}, 1}, // 2

		{[2][]byte{[]byte( // 3
			"ajklsdfasdfkljklasjkasdfflkasdfaklfiwe492104|X/epDDaLFDJAKLDAIWALwd"),
			[]byte("ajklsdfasdfkljklasjkasdfflkasdfaklfiwe492104|X/epDDaLFDJAKLDAIWALwd")},
			37,
		},
		{
			[2][]byte{
				[]byte("C¨\x9ap=\x80\x03\x0c\x0f\x83?\x12Ü\xd5&\x99w\\\xb9\xf8m`\xc6\xe8\x91f¤\x00}\xd0)t\x9e\x05q\xb6}ýßWx&\x11\x8c\xd9"),
				[]byte("C¨\x9ap=\x80\x03\x0c\x0f\x83?\x12Ü\xd5&\x99w\\\xb9\xf8m`\xc6\xe8\x91f¤\x00}\xd0)t\x9e\x05q\xb6}ýßWx&\x11\x8c\xd9"),
			},
			42,
		},
		{[2][]byte{[]byte("Î5_\x1fÝnD\xef"), []byte("Î5_\x1fÝnD\xef")}, 4}, // 5 −  false  ⇐ "Î" != "I"+"̂"

		{[2][]byte{[]byte(""), []byte("")}, 0},
	}
	expectTab := [tcQuantity]bool{
		false, false, true, // 0, 1, 2
		true, true, false, // 3, 4, 5
		true, // 6
	}

	a := asserter.PNewTiny(func(s string) { t.Error(s) })
	for tcInd, args := range argTab {
		n, sh1, sh2 := args.n, *(*reflect.SliceHeader)(unsafe.Pointer(&args.a[0])), *(*reflect.SliceHeader)(unsafe.Pointer(&args.a[1]))
		ex := expectTab[tcInd]
		got := CmpNBytes(n, unsafe.Pointer(sh1.Data), unsafe.Pointer(sh2.Data))
		a.AssertEq(ex, got)
	}
}
