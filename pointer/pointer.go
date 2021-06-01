package pointer

import "fmt"

func demo1()  {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"


	

}


//在Go语言中，返回函数中局部变量的地址也是安全的
var p = f()
func f() *int {
	v := 1
	return &v
}
func demo2()  {
	fmt.Println(f() == f()) // "false"
}


func demo3()  {
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"
}
//下面的两个newInt函数有着相同的行为,new函数类似是一种语法糖，而不是一个新的基础概念
func newInt() *int {
	return new(int)
}
func newInt() *int {
	var dummy int
	return &dummy
}