package main

import "fmt"

type user struct{
	mid int32
}

func main(){
	var uObj *user
	fmt.Println(uObj)	//打印指针
	fmt.Println(fmt.Sprintf("the address of uObj = %p",uObj))
	fmt.Println(*uObj)	//打印指针指向的值
}