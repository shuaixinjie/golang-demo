package _type

import "fmt"

// IIT 是新定义的类型，并不是别名
// 不能用原类型的方法
// 相互转化必须是显示的，如IIT(int) int(IIT)这样，值得一提的是golang所有的类型都必须显示转换
type IIT int

type A struct {
	ID int
}

func (a *A) f() {
	fmt.Printf("hello,i'm %d\n", a.ID)
}

// AA 这才是取别名的正确姿势，一模一样的功能，不一样的名字
type AA = A

// AB 内部属性相同，可是方法不继承
type AB A

func test1() {
	a := A{ID: 1}
	a.f()
	aa := AA{ID: 2}
	aa.f()
	ab := AB{ID: 3}
	fmt.Println(ab.ID)
	//ab.f() //ab.f undefined (type AB has no field or method f)
}

// typeFunc 更高级一点的用法，函数也是一个确定的类型
type typeFunc func(int, int) int


