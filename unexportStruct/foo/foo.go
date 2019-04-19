package foo

type St struct{
	Name string
	age int32
}

//Init a St
func NewSt(name string, age int32)*St{
	st := new(St)
	st.Name = name
	st.age = age

	return st
}

func (st *St)Age()int32{
	return st.age
}
