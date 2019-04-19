package main

import (
	"goDemo/unexportStruct/foo"
	"reflect"
	"log"
)

func main(){
	st := foo.NewSt("Lucifer",25)
	log.Printf("the st name = %s",st.Name)

	//How to get st.age,solution one:reflect
	refV := reflect.ValueOf(st)
	//the statement below will panic:panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
	//age := refV.Elem().FieldByName("age").Interface()
	age := int32(refV.Elem().FieldByName("age").Int())
	log.Printf("the solution one,st age = %d",age)

	//How to get st.age,solution two:
	log.Printf("the solution two,st age = %d",st.Age())
}
