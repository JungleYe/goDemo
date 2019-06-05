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


//方式二：另外一种设置配置的方式
//看起来不是很顺利，忘记了是哪里有记录这种用法的。很奇怪
func setHeartbeat(c *Client,v interface{})error{
	val,ok := v.(bool)
	if ok{
		c.option.heartbeat = val
	}
	return nil
}

//方式二：另外一种设置配置的方式
func setReadTimeout(c *Client,v interface{})error{
	val,ok := v.(int32)
	if ok{
		c.option.readTimeOut = val
	}
	return nil
}

//方式二：设置配置的统一入口
func (c *Client)SetOptions(fn func(x *Client,v interface{})error, vs interface{})error{
	return fn(c,vs)
}

func main(){
	//方式一
	obj := NewClient1("127.0.0.1",options{true,2})
	fmt.Printf("obj = %v",obj)

	//方式二：
	obj2 := new(Client)
	obj2.SetOptions(setHeartbeat,true)	//设置是否心跳
	obj2.SetOptions(setReadTimeout,int32(35))	//设置读超时时间
	fmt.Printf("obj2 = %v",obj2)
}