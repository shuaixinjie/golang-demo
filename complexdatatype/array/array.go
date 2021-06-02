package array

import (
	"fmt"
)

func demo1()  {
	var a [3]int             // array of 3 integers
	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}


	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

}

//指针数组，数组指针
func demo2()  {
	//前者是存了一堆指针
	i1 := 2
	p1 := &i1
	arr := [3]*int{new(int), new(int), p1}
	for i, v := range arr{
		fmt.Println(*arr[i])
		fmt.Println(*v)
	}

	//后者指针指向数组，这个可以做数组的引用传递
	arr2 := [3]int{2:1}
	parr := &arr2
	func(arr *[3]int) {
		arr[0] = 12  //等价于 (*arr)[0] = 12
	}(parr)
	fmt.Println(arr2)
}

