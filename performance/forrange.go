package performance

import (
	"fmt"
	"math/rand"
	"time"
)

func demo1()  {
	words := []string{"Go", "语言", "高性能", "编程"}

	//以下两种遍历方式结果一致，第二种基本就和fori没什么差异了
	for i, s := range words {
		words = append(words, "test") //追加或删除数据不会影响遍历次数
		fmt.Println(i, s)
	}
	for i := range words {
		fmt.Println(i, words[i])
	}
}

//for和range的性能比较，看基准测试

type Item struct {
	id  int
	val [4096]byte
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}