package function

import "fmt"

//recover捕获异常
func demo1()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [10]int
	index := 11
	a[index] = 1
}
