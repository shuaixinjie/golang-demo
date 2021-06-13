package standradlib

import (
	"fmt"
	"strings"
)

//strings包实现了用于操作字符的简单函数

func demo1()  {
	s := "hello"
	fmt.Println(strings.HasPrefix(s, "he")) //true
	fmt.Println(strings.HasSuffix(s, "lo")) //true
	fmt.Println(strings.Contains(s, "ell")) //true
	fmt.Println(strings.ContainsRune(s, 'e'))   //true
	fmt.Println(strings.ContainsAny(s, "aaaae"))  //true，字符串中有任意的匹配就为true
	fmt.Println(strings.Count(s, "ll")) //1, 返回字符串s中有几个不重复的子串
	//基本上也都有LastIndex...
	fmt.Println(strings.Index(s, "ll")) //2, 子串在字符串s中第一次出现的位置，不存在则返回-1
	fmt.Println(strings.IndexByte(s, 101))  //1，字符在s中第一次出现的位置，不存在则返回-1
	fmt.Println(strings.IndexRune(s, 'e'))  //1，unicode码值r在s中第一次出现的位置，不存在则返回-1
	fmt.Println(strings.IndexAny(s, "ehoasd")) //0，字符串chars中的任一utf-8码值在s中第一次出现的位置，如果不存在或者chars为空字符串则返回-1
	fmt.Println(strings.IndexFunc(s, func(r rune) bool {
		return r == 'l'
	}))  //2, s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1
	fmt.Println(strings.Title("hello")) //Hello 首字母大写
	fmt.Println(strings.ToLower("HeLlO")) //所有字母小写
	fmt.Println(strings.ToTitle("hello")) //全大写
	fmt.Println(strings.Repeat(s, 3)) //hellohellohello，复制字符串
	//当然也直接提供了ReplaceAll
	fmt.Println(strings.Replace(s, "l", "2", -1)) //he22o，返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'l' {
			return 'A'
		}
		return r
	},s))  //heAAo
	//TrimLeft、TrimRight
	fmt.Println(strings.Trim(s, "hlo"))  //e，返回将s前后端所有包含的utf-8码值都去掉的字符串
	fmt.Println(strings.TrimSpace("  he llo  ")) //he llo，返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串
	//TrimLeftFunc、TrimRightFunc
	fmt.Println(strings.TrimFunc(s, func(r rune) bool {
		return r == 'o'
	})) //hell，trim只能操作前后端，中间不行
	//TrimSuffix
	fmt.Println(strings.TrimPrefix(s, "he")) //llo
	fmt.Println(strings.Fields("  ab c de "))  //[ab c de]   []string
	fmt.Println(strings.FieldsFunc("abbcbdb beb", func(r rune) bool {
		return r == 'b'
	}))  //[a c d   e] 自定义分割条件

	//上面的方式分割，下面的切割有所不同，即使两个sep相邻，也会进行两次切割，如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串
	//SplitN可指定切割数，SplitAfter反向切割，SplitAfterN
	split := strings.Split("abbcbdb beb", "b")
	for i, s2 := range split {
		fmt.Println(i, s2)
	}

	fmt.Println(strings.Join([]string{"aa", "bb", "cc"}, "~~")) //将一系列字符串连接为一个字符串，之间用sep来分隔

	//还有两个结构体Reader、Replacer暂时略
}