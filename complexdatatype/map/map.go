package _map

import "fmt"

func demo()  {
	//内置的make函数可以创建一个map
	ages := make(map[string]int) // mapping from strings to ints

	//我们也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	//使用内置的delete函数可以删除元素
	delete(ages2, "alice") // remove element ages["alice"]

	//map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作
	//_ = &ages["bob"] // compile error: cannot take address of map element

	//迭代顺序是不确定的
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}


	age, ok := ages["bob"]
	if !ok { /* "bob" is not a key in this map; age == 0. */ }


}
