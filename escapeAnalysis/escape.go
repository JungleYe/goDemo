package main

func foo()*int32{
	var i int32
	return &i
}

func bar()int32{
	x := new(int32)
	*x = 123
	return *x
}

//逃逸分析：由runtime自动识别
func main(){
	//go run -gcflags "-m -l" escape.go

	//打印结果如下：
	//.\escape.go:5:9: &i escapes to heap
	//.\escape.go:4:6: moved to heap: i
	//.\escape.go:9:10: bar new(int32) does not escape
}
