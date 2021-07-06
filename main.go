package main

import "fmt"

//https://books.studygolang.com/gopl-zh/ch5/ch5-10.html

func main() {
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

}
