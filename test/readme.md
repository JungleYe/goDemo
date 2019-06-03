//golang V1.10之后，在功能代码和测试代码没有变更的情况下执行go test会直接读取缓存中的测试结果；
//而且 go test -v . 和 go test .是分开缓存的

//所以测试代码如果是要跑一些网络服务的情况下，网络服务上的数据更新是不会立刻刷新的
//golang测试的坑：go test -count=1  -v .

//在执行go test -v的时候是有相关的线索的：ok  	mygo/lib	(cached)
