package main

import "unsafe"

type struct2 struct {
	f int64
	g int64
	h float64
	i float64
} // 32 bytes

func main() {
	// var var_item float64 = 8
	s := &struct2{}
	a := unsafe.Sizeof(s)
	print(a)

	var t = 1
	t++
	print(t)
}
