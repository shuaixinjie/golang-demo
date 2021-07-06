package main

import (
	"errors"
	"fmt"
)

//https://books.studygolang.com/gopl-zh/ch5/ch5-10.html

func main() {
	err := test1()

	fmt.Println("aaa")
}

type T struct {
	Name string
}

func test1() error {
	return errors.New("err1")
}