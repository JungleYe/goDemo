package main

import "fmt"

type Client struct{
	ip string
	port int32
	option options
}

type options struct{
	heartbeat bool
	readTimeOut int32
}

//创建一个client的方法：一
func NewClient1(ip string, option options)*Client{
	obj := new(Client)
	obj.option = option
	return obj
}

//创建option
func NewOpts(heartbeat bool, readTimeout int32)options{
	return options{heartbeat:heartbeat,readTimeOut:readTimeout}
}


////方式二：另外一种设置配置的方式
////看起来不是很顺利，忘记了是哪里有记录这种用法的。很奇怪
//func setHeartbeat(c *Client,v interface{})error{
//	val,ok := v.(bool)
//	if ok{
//		c.option.heartbeat = val
//	}
//	return nil
//}
//
////方式二：另外一种设置配置的方式
//func setReadTimeout(c *Client,v interface{})error{
//	val,ok := v.(int32)
//	if ok{
//		c.option.readTimeOut = val
//	}
//	return nil
//}
//
////方式二：设置配置的统一入口
//func (c *Client)SetOptions(fn func(x *Client,v interface{})error, vs interface{})error{
//	return fn(c,vs)
//}

//=============方式三：
//方式三的好处：参数无序，且耦合度低
//假设由于业务需要，添加了新的Options有了扩展，按这种方式只用加上对应的WithXXX方法就可以了。非常方便
type OptionFn func(o *options)

//创建option
func NewOptions (opts ...OptionFn)options{
	obj := options{}
	for _,o := range opts{
		o(&obj)
	}
	return obj
}

//设置是否开启心跳
func WithHeartbeat(flag bool)OptionFn{
	return func(o *options){
		o.heartbeat = flag
	}
}
//设置读超时时间
func WithReadTimeout(t int32)OptionFn{
	return func(o *options){
		o.readTimeOut = t
	}
}

func main(){
	//方式一
	opt1 := NewOpts(true,2)
	obj := new(Client)
	obj.option = opt1
	fmt.Printf("obj = %v",obj)

	////方式二：
	//obj2 := new(Client)
	//obj2.SetOptions(setHeartbeat,true)	//设置是否心跳
	//obj2.SetOptions(setReadTimeout,int32(35))	//设置读超时时间
	//fmt.Printf("obj2 = %v",obj2)

	//方式三：
	opt2 := NewOptions(WithHeartbeat(true),WithReadTimeout(12))
	obj3 := new(Client)
	obj3.option = opt2
	fmt.Printf("obj3 = %v",obj3)
}