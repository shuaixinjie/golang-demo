package factory

// 工厂方法

/**
	是简单工厂的进一步封装
	什么时候用呢？
	简单工厂可以在对象很简单的时候直接return一个实例
	而对象稍微复杂一些的时候，可以用工厂方法，将复杂的逻辑拆分到各个工厂类中，让每个工厂类都不至于过于复杂
 */

type ICarFac interface {
	CreateCar() ICar
}

// bmwFac 工厂类
type bmwFac struct {

}

func (b *bmwFac) CreateCar() ICar {
	return &BMW{}
}

type benzFac struct {

}

func (b *benzFac) CreateCar() ICar {
	return &Benz{}
}

func NewCarFac(t string) ICarFac {
	switch t {
	case "bmw":
		return &bmwFac{}
	case "benz":
		return &benzFac{}
	}
}