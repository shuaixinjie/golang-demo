package reflect

import (
	"fmt"
	"reflect"
)

//为什么需要反射？


//reflect.Type和reflect.Value


func demo1()  {
	//内部使用 reflect.TypeOf 来输出
	fmt.Printf("%T\n", 3)

	fmt.Println(reflect.TypeOf(3))
	fmt.Println(reflect.TypeOf(3))
	fmt.Println(reflect.TypeOf(3).String())


	fmt.Println(reflect.ValueOf(4))
	fmt.Println(reflect.ValueOf(4).String())
	fmt.Println(reflect.ValueOf(4).Type())
	fmt.Println(reflect.ValueOf(4).Interface())
	fmt.Println(reflect.ValueOf(4).Interface().(int))



}
