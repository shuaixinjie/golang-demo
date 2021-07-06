# go编码规范

仅仅是为了写出更优雅的代码，不是代码层面的硬性要求  
想一点写一点吧

1. 给难以区分的包取个好听的名字

```go
jwt "github.com/dgrijalva/jwt-go/v4"
```

2. 简短声明的坑

```go
var s *sql.DB
s, err := ... // 这里s就会被重新定义声明
```

2. 有必要的预分配

```go
v1 := make(map[int]string, 4)
v2 := make([]string, 0, 4)
```

3. 对于未导出的顶层常量和变量，使用 _ 作为前缀

```go
// bad
const (
defaultHost = "127.0.0.1"
defaultPort = 8080)

// good
const (
_defaultHost = "127.0.0.1"
_defaultPort = 8080
)
```

4. 错误处理1 对于有返回error的函数，不可以不返回，两种策略   
   用参数err接收处理，或者用_直接忽略这个错误   
   对于defer xx.Close() 可以不处理

```go
_ = load() // load()返回一个err直接不处理
```

5. 错误处理2 error返回要放在最后

```go
// bad
func load() (error, int)

// good
func load() (int, error) 
```

6. 错误处理3 err要单独判断，不与其他逻辑组合

```go
// bad
v, err := foo()
if err != nil || v  == nil
// good
v, err := foo()
if err != nil
```

7. 错误处理4 err要单独判断，不与其他逻辑组合

```go
// bad
v, err := foo()
if err != nil || v  == nil
// good
v, err := foo()
if err != nil
```

8. 错误处理5  
   a. 告诉用户他们可以做什么，而不是告诉他们不能做什么。  
   b. 当声明一个需求时，用 must 而不是 should。例如，must be greater than 0、must match regex '[a-z]+'。  
   c. 当声明一个格式不对时，用 must not。例如，must not contain。  
   d. 当声明一个动作时用 may not。例如，may not be specified when otherField is empty、only name may be specified。  
   e. 引用文字字符串值时，请在单引号中指示文字。例如，ust not contain '..'。  
   f. 当引用另一个字段名称时，请在反引号中指定该名称。例如，must be greater than request。  
   g. 指定不等时，请使用单词而不是符号。例如，must be less than 256、must be greater than or equal to 0 (不要用 larger than、bigger than、more
   than、higher than)。  
   h. 指定数字范围时，请尽可能使用包含范围。  
   i. 建议 Go 1.13 以上，error 生成方式为 fmt.Errorf("module xxx: %w", err)。  
   j. 错误描述用小写字母开头，结尾不要加标点符号

9. panic  
   **一般不用**

10. 类型断言   
    一般必须加上第二个参数判断是否断言成功

```go
// x一定是接口类型
value, ok := x.(int)
```

11. 接口命名   
    a. 单个函数的接口名以 “er"”作为后缀（例如 Reader，Writer），有时候可能导致蹩脚的英文，但是没关系。   
    b. 两个函数的接口名以两个函数名命名，例如 ReadWriter。   
    c. 三个以上函数的接口名，类似于结构体名。  

12. 一些经验  
```go
// 空字符串判断
// bad
if s == "" 
// good
if len(s) == 0 

// []byte/string 相等比较
// bad
var s1 []byte
var s2 []byte
...
bytes.Equal(s1, s2) == 0
bytes.Equal(s1, s2) !=0

// good
var s1 []byte
var s2 []byte
...
bytes.Compare(s1, s2) == 0
bytes.Compare(s1, s2) != 
	
// 是否包含字串
// bad
strings.Contains(s, subStr)
strings.ContainsAny(s, char)
strings.ContainRune(s, )
// good
strings.Index(s, subStr) > -1
strings.IndexAny(s, char) > -1
strings.IndexRune(s, r) > -

// 除去前后字符串  
// bad
var s1 = "a string value"
var s2 = "a "
var s3 = strings.TrimPrefix(s1, s)
// good
var s1 = "a string value"
var s2 = "a "
var s3 string
if strings.HasPrefix(s1, s2) {
s3 = s1[len(s2):]
}

// 空slice判断，map、channel同样如此
// bad
if len(slice) > 0
// good
if slice != nil && len(slice) > 0 


// slice复制
// bad
var b1, b2 []byte
for i, v := range b1 {
b2[i] = v
}
for i := range b1 {
b2[i] = b1[i]
}
// good
copy(b2, b1)

// slice新增
// bad
var a, b []int
for _, v := range a {
b = append(b, v)
}
// good
var a, b []int
b = append(b, a...)

```

13. 函数  
**尽量值传递，而非指针传递**  
**传入map、slice、chan、interface，本身就是引用传递，不要传指针**    
**返回值超过三个的时候，请使用struct**      

