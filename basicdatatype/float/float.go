package float

import (
	"fmt"
)

//float32  float64
func demo1()  {
	//超出表达范围会出现误差
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
}
