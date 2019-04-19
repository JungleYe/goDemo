package foo

import(
	_ "unsafe"
	_ "goDemo/unexport/bar"
)

//go:linkname PrintBar bar.pb
func PrintBar()
