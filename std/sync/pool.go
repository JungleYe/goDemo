package main

import (
	"sync"
	"log"
)

var p = sync.Pool{
	New : func()interface{}{
		return new(ST)
	},
}

type ST struct{
	Name string
	Age int32
}

func main(){
	sta := new(ST)
	stb := new(ST)
	p.Put(sta)
	p.Put(stb)

	s := p.Get().(*ST)
	s.Name = "111"
	s.Age = 111

	s2 := p.Get().(*ST)
	s2.Name = "222"
	s2.Age = 222

	p.Put(s)
	p.Put(s2)

	s3:= p.Get().(*ST)
	log.Println(s3)

}

