// Original code from https://github.com/AlekSi/cgo-by-example
package main

import (
	"fmt"
	"unsafe"
)

/*
#include <stdlib.h>
#include "lib/lib.h"
#cgo LDFLAGS: -L . -l lib
*/
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func F(i C.int) {
	fmt.Println("Go F()")
}

func main() {
	fmt.Println("Cgo random:", Random())

	C.f0()

	C.f1(5)

	var i int = 8
	C.f1(C.int(i))

	// Convert Go string to C string (char*) and back.
	cs := C.CString("Go string")
	csRet := C.f2(cs)
	fmt.Printf("fmt: %s\n", C.GoString(csRet))
	C.free(unsafe.Pointer(cs))
	//defer C.free(unsafe.Pointer(csRet))

	// Pass C struct to C function by value.
	s1 := C.struct_s1{a: 5} // or cast with C.int: s1 := C.struct_s1{C.int(i)}
	s1Ret := C.f31(s1)
	fmt.Printf("f31: s1=%v, s1Ret=%v\n", s1, s1Ret)

	// Pass C struct to C function by pointer.
	s1 = C.struct_s1{a: 5}
	s1Ret = *C.f32(&s1)
	fmt.Printf("f32: s1=%v, s1Ret=%v\n", s1, s1Ret)

	// Pass C struct with int pointer to C function.
	// s2 := C.struct_s2{&C.int(i)} - compile error "cannot take the address of _Ctype_int(i)", create variable as below
	ci := C.int(i)
	s2 := C.struct_s2{p: &ci}
	fmt.Printf("s2=%v, *s2.p=%d, i=%d\n", s2, *s2.p, i)
	C.f4(s2)
	fmt.Printf("f4: s2=%v, *s2.p=%d, i=%d\n", s2, *s2.p, i)

	// Pass C function pointer to C function.
	// C.f5(C.f1)                 - compile error "must call C.f1"
	// fp := unsafe.Pointer(C.f1) - compile error "must call C.f1"
	// It is not possible to pass C function pointer via Go.

	// Pass Go function pointer to C function.
	f := func(i C.int) {
		fmt.Println("Go f()")
	}
	fp := unsafe.Pointer(&f)
	// C.f5((*[0]byte)(fp)) - Compiles but explodes. Do not do this!

	// Pass global Go function pointer to C function.
	f = F
	fp = unsafe.Pointer(&f)
	// C.f5((*[0]byte)(fp)) - Compiles but explodes. Do not do this!
	// It is possible to call Go function from C. See main2.go for solution.

	_ = fp
}
