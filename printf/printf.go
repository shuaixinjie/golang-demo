package printf

import "fmt"

/**
格式化输出
*/

// Test1 当不知道是什 么类型的数据可以直接使用%v进行输出
func Test1() {
	var a int = 100
	var b bool
	var c byte = 'a'

	fmt.Printf("a:%v\n", a)
	fmt.Printf("b:%v\n", b)
	fmt.Printf("c:%v\n", c)
	/**
		a:100
		b:false
		c:97
	*/
}

func Test2()  {
	type student struct {
		Name string
		Age  int
	}
	stu := student{
		Name: "lily",
		Age:  25,
	}
	fmt.Printf("stu:%v\n", stu)
	fmt.Printf("stu:%+v\n", stu)
	fmt.Printf("stu:%#v\n", stu)
	/**
		stu:{lily 25}
		stu:{Name:lily Age:25}
		stu:main.student{Name:"lily", Age:25}
	*/
}








