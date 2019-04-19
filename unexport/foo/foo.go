package foo

import(
	_ "unsafe"
	_ "glgo/testcode/unexport/bar"
)

//go:linkname PrintBar bar.pb
func PrintBar()
