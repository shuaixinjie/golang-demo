package constant

import "fmt"

type Weekday int

const (
	a = 1
	b
	c = 2
	d
	Thursday Weekday = iota   //常量生成器，这里是4了，下面是5和6
	Friday
	Saturday

	cc = 1 << iota  //2**7  可以给iota指定位运算规则
	dd              //2**8
)

func demo1()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(cc)
	fmt.Println(dd)
}
