package string

import (
	"fmt"
	"unicode/utf8"
)

//一个字符串是一个不可改变的字节序列
//内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目）

func demo1()  {
	s := "hello"
	//字符串的值是不可变的，s的值变了是因为字符串变量分配一个新字符串值 s[1] = 'a'这样则是不被允许的
	s += ",world!"
	fmt.Println(s)

	//我们也可以用下标索引，超出索引即panic
	fmt.Println(s[6:11])
}

func demo2()  {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
}

//字符串和字节切片的转换
func demo3()  {
	s := "abc"
	b := []byte(s)
	fmt.Println(b)
	s2 := string(b)
	fmt.Println(s2)
}

//字符串操作...
