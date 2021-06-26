package _struct

// 结构体是值类型，所以我们用new来初始化它，并返回一个指针
// 我们一般的惯用方法是：t := new(T)，变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值

// 表达式 new(Type) 和 &Type{} 是等价的。&struct1{a, b, c} 是一种简写，底层仍然会调用 new ()，这
// 里值的顺序必须按照字段顺序来写。也可以通过在值的前面放上字段名来初始化字段的方式，这种方式就不必按照顺序来写了

type Human struct {
	name string
}

type Student struct { // 含内嵌结构体Human
	Human // 匿名（内嵌）字段，这个实现了其他语言所谓的继承，在go中叫组合
	int   // 匿名（内嵌）字段
}



