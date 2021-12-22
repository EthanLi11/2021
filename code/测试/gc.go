package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

type struct2 struct {
	f int64
	g int64
	h float64
	i float64
} // 32 bytes

type struct1 struct {
	a int64
	b int64
	c float64
	d float64
	e *struct2
} // 40 bytes

// go:noinline
func allocStruct1() *struct1 {
	return &struct1{
		e: allocStruct2(),
	}
}

// go:noinline
func allocStruct2() *struct2 {
	return &struct2{}
}

// go run gc.go 2> trace.out
// go tool trace trace.out
func main() {

	_ = trace.Start(os.Stderr)
	defer trace.Stop()

	s1 := allocStruct1()
	s2 := allocStruct2()
	func() {
		_ = allocStruct2()
	}()
	runtime.GC()
	fmt.Printf("s1 = %X, s2 = %X\n", &s1, &s2)
}
