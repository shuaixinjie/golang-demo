package reflect

import (
	"fmt"
	"reflect"
)

//为什么需要反射？

//程序运行的时候去获取数据类型


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

//获取传进来的参数类型
//1.类型断言，可是类型太多了
//2.反射
func reflectType(x interface{})  {
	obj := reflect.TypeOf(x)
	//kind是所属的种类，Cat归属于struct
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n",obj)
}

type Cat struct {}

//获取类型
func demo2()  {
	var a float32 = 1.234
	reflectType(a)
	var b int8 =2
	reflectType(b)
	var c Cat
	reflectType(c)
	var d string
	reflectType(d)
	var e []int
	//map、array、slice、channel获取不到Name()
	reflectType(e)
}

func reflectValue(x interface{})  {
	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n", v, v)
	kind := v.Kind()
	switch kind {
	case reflect.Float32:
		//转化为int32
		f := float32(v.Float())
		fmt.Printf("%v, %T\n", f, f)
	case reflect.Int32:
		i := int32(v.Int())
		fmt.Printf("%v, %T\n", i, i)
	}
}
//获取值
func demo3()  {
	var a float32 = 1.234
	reflectValue(a)
	var b int32 =2
	reflectValue(b)
	//var c Cat
	//reflectValue(c)
	//var d string
	//reflectValue(d)
	//var e []int
	//reflectValue(e)
}

func reflectSetValue(x interface{})  {
	v := reflect.ValueOf(x)
	//Elem()用来根据指针取对应的值
	kind := v.Elem().Kind()
	switch kind {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)
	}
}
//设置值
func demo4()  {
	var a int32 = 10
	reflectSetValue(&a)
	fmt.Println(a)
}

//https://www.bilibili.com/video/BV13t411N7jZ?from=search&seid=2683123561545666212



