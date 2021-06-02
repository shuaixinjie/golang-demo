package slice

import "fmt"

//关于切片的底层完全是数组，就不过多赘述了

func demo1()  {

	//数组可以用 == 判断值是否相等，切片不可以，唯一合法的是和nil比较

	//为了创建的切片不为nil，一般直接初始化，有如下两种方式
	s1 := []int{1, 2, 3}
	s2 := make([]string, 0)  //make([]T, len, cap)

	fmt.Println(s1, s2)
}

func demo2()  {
	s1 := make([]int, 0)
	//扩容的规则是2倍
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
		fmt.Println("len = ", len(s1))
		fmt.Println("cap = ", cap(s1))
		fmt.Println("slice = ", s1)
	}

	//append(slice []Type, elems ...Type) []Type
	s2 := []int{1,2}
	s2 = append(s2, 3, 4)
	s2 = append(s2, s1...)
	fmt.Println(s2)
}

//append操作的底层实现依赖copy
func demo3()  {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中  slice2:[1 2 3]
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

}

//delete elem
func demo4()  {
	ageList := []int{1, 3, 7, 7, 8, 2, 5}
	//遍历删除6以下的
	func(a []int) {
		for i := 0; i < len(a); {
			if a[i] < 6 {
				a = append(a[:i], a[i+1:]...)
			} else {
				i++
			}
		}
		fmt.Println(a)    //[7 7 8]
	}(ageList)
	fmt.Println(ageList)  //[7 7 8 5 5 5 5]

	sli := []int{1 ,2, 3, 4, 5}
	//删除下标为1的元素
	fmt.Println(append(sli[:1], sli[2:]...))
}














