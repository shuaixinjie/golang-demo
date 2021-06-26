package error

import "fmt"

func div(a, b int)  {
	defer func() {
		if r := recover(); r !=nil {
			fmt.Printf("捕获到异常：%s\n", r)
		}
	}()
	fmt.Println("余数为：", a/b)
}