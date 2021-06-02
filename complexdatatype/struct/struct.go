package _struct

import "fmt"

//结构体是一种聚合的数据类型

//一个结构体不能包含该结构体成员，但可以包含该结构体的指针（这个可以实现链表，二叉树等）
type tree struct {
	value       int
	left, right *tree
}

type Point struct {
	x, y int
}

func demo1()  {
	//常用的定义方式
	po := Point{
		x: 1,
		y: 2,
	}

	//较大的结构体，操作的时候一般用指针，下面两种方式等价
	pp1 := &Point{1, 2}
	pp2 := new(Point)
	*pp2 = Point{1, 2}

	fmt.Println(po, pp1, pp2)
}

//如果结构体的全部成员都是可以比较的，那么该结构体就可以比较
//map[key]value  可比较的类型是可以放入这个key的，当然value是什么都可以放

