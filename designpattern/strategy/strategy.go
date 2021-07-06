package strategy

/**
	策略模式
	在策略模式中，我们创建表示各种策略的对象和一个行为随着策略对象改变而改变的 context 对象。策略对象改变 context 对象的执行算法
	1、算法可以自由切换。 2、避免使用多重条件判断。 3、扩展性良好。
 */

func test(t string, a, b int) int {
	m := new(context)
	switch t {
	case "add":
		m = newContext(newAdd())
	case "sub":
		m = newContext(newSubtract())
	case "mul":
		m = newContext(newMultiply())
	}
	return m.str.doOperation(a, b)
}

type strategy interface {
	doOperation(a, b int) int
}

type add struct {
}

func newAdd() *add {
	return &add{}
}

func (s *add) doOperation(a, b int) int {
	return a + b
}

type subtract struct {
}

func newSubtract() *subtract {
	return &subtract{}
}

func (s *subtract) doOperation(a, b int) int {
	return a - b
}

type multiply struct {
}

func newMultiply() *multiply {
	return &multiply{}
}

func (s *multiply) doOperation(a, b int) int {
	return a * b
}

type context struct {
	str strategy
}

func newContext(str strategy) *context {
	return &context{str: str}
}