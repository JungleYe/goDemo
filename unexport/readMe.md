#### 参考文章：
    https://colobu.com/2017/05/12/call-private-functions-in-other-packages/
    
#### demo说明：
> golang的首字母大写的方法才对外可访问，但看源码的时候，发现很多标准库方法的实现都是在runtime中，而runtime中的那些方法是小写的，如何组织的呢？
> 看了参考文章之后，写下此demo进行验证。
