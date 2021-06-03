package performance

import "testing"

func TestDemo1(t *testing.T) {
	demo1()
}

func BenchmarkForIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)
	for i := 0; i < b.N; i++ {
		len := len(nums)
		var tmp int
		for k := 0; k < len; k++ {
			tmp = nums[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, num := range nums {
			tmp = num
		}
		_ = tmp
	}
}

func BenchmarkForStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeIndexStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		//特别注意，仅仅遍历下标的时候，和fori性能无差
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		//当遍历值的时候，会拷贝这个值，普通int类型，效率无差，这里的struct类型性能降低了很多
		//当然切片中如果放入的是指针，性能就没差了
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}