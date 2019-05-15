//读 github.com\smallnest\rpcx\util
package main

import (
	"unsafe"
	"log"
)

func main(){
	var sli = []byte{'a','b','c','d'}

	str1 := string(sli)	//拷贝了sli的内存数据，据为己有。
	str2 := SliceByteToString(sli)	//共享着sli的数据
	log.Printf("str1=%s,str2=%s",str1,str2)
	log.Printf("sli addr=%d, str1 addr=%d, str2 addr=%d",&sli[0],&str1,&str2)

	sli[0] = 'b' // 修改原始slice的值
	log.Printf("str1=%s,str2=%s",str1,str2)
}

//任何指针都可以转换为unsafe.Pointer
//unsafe.Pointer可以转换为任何指针
//uintptr可以转换为unsafe.Pointer
//unsafe.Pointer可以转换为uintptr

//高效的字节切片转换成string的方法
func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//高效的string转换成字节切片的方法
func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

