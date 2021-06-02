package int

import "fmt"

//int8、int16、int32和int64
//8、16、32、64bit

//uint8、uint16、uint32和uint64
//int8值域是从-128到127  uint8值域是从0到255

//Unicode字符rune类型是和int32等价的类型
//byte也是uint8类型的等价类型

//还有一种无符号的整数类型lintptr，没有指定具体的bit大小但是足以容纳指针

//除法运算符/的行为则依赖于操作数是否为全为整数，比如5.0/4.0的结果是1.25，但是5/4的结果是1，因为整数除法会向着0方向截断余数
func demo1() {
	fmt.Println(4/3)
	fmt.Println(4.0/3)
}

//位运算
func demo2()  {
	fmt.Println(1 << 3)
	fmt.Println(13 >> 2)
	//异或，不相同的时候为1   13-->1101  15 15 1
	fmt.Println(13 ^ 2)
	//或，有1的时候就为1
	fmt.Println(13 | 3)
	//与，两个1才是1
	fmt.Println(13 & 3)
	//正数取反  0000 0010 --> 1111 1101（取反） --> 1111 1100（减1） --> 1000 0011（除符号位其余取反）-->-3
	fmt.Println(^ 2)
	//负数取反  1000 0101 --> 1111 1010（取补码）--> 1111 1011（加1） --> 0000 0100（取反）-->4
	fmt.Println(^ -5)
}


