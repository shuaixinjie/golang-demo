package standradlib

import (
	"fmt"
	"strconv"
)

func strconvDemo1()  {
	/**
		case "1", "t", "T", "true", "TRUE", "True":
			return true, nil
		case "0", "f", "F", "false", "FALSE", "False":
			return false, nil
	 */
	fmt.Println(strconv.ParseBool("1")) //true <nil>

	/**
		参数1 数字的字符串形式
		参数2 数字字符串的进制 比如10 二进制 八进制 十进制 十六进制
		参数3 返回结果的bit大小 也就是int int8 int16 int32 int64
	 */
	fmt.Println(strconv.ParseInt("123", 0, 0))
	//Atoi是ParseInt(s, 10, 0)的简写
	fmt.Println(strconv.Atoi("123"))
	fmt.Println(strconv.Itoa(123))

}
