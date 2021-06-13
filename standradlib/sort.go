package standradlib

import (
	"fmt"
	"sort"
)

func sortDemo1()  {
	//a := []int{1 ,22,2,33,1}
	//sort.Ints(a)
	//fmt.Println(a)
	s1 := Student{
		name:  "harry",
		score: 45,
	}
	s2 := Student{
		name:  "lucy",
		score: 42,
	}
	s3 := Student{
		name:  "jerry",
		score: 55,
	}
	sss := []Student{s1, s2, s3}
	sort.Sort(stus(sss))
	fmt.Println(sss)

	a := []int{4,3,2,1,5,9,8,7,6}
	sort.Sort(IntSlice(a))
	fmt.Println("After sorted: ", a)


	b := []string{"aa", "dd", "cc", "bb"}
	sort.Sort(MyString(b))
	fmt.Println("After sorted: ", b)
}

//实现三个接口则可以实现排序功能

type IntSlice []int
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

type MyString []string
func (s MyString) Len() int { return len(s) }
func (s MyString) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s MyString) Less(i, j int) bool { return s[i] < s[j] }



type Student struct {
	name string
	score int
}

type stus []Student

func (s stus) Len() int {  //排序长度
	return len(s)
}
func (s stus) Less(i, j int) bool { //排序顺序
	return s[i].score < s[j].score
}
func (s stus) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

