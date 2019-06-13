package bigmap

type st struct{
	v1 string
	v2 int32
	v3 string
	v4 string
	v5 string
	v6 string
	v7 string
	v8 string
}

var mp = make(map[string]*st,1000000)

func AddOneItem(k string){
	mp[k] = &st{v1:"hhhhhv1",v2:2,v8:"hhhhhv8"}
}

func DelOneItem(k string){
	delete(mp,k)
}