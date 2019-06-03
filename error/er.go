package main

import (
	"errors"
	"fmt"
)

var ER_HELLO = errors.New("hello global var")

func main(){
	//.\er.go:10:13: e escapes to heap
	e := Tst()
	fmt.Println(e)

	fmt.Print(ER_HELLO)
}

func Tst()error{
	return errors.New("what happened!")
}
