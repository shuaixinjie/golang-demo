package factory

// 简单工厂

/**
	现在项目里面也会涉及到这样的操作，只是没有这么多的实例，单个的我们也这么写了
 */

type ICar interface {
	ProduceCar() string
}

type BMW struct {

}

func (b *BMW) ProduceCar() string {
	return "BMW X5"
}

type Benz struct {

}
func (b *Benz) ProduceCar() string {
	return "Benz G"
}

func NewCar(t string) ICar {
	switch t {
	case "benz":
		return &Benz{}
	case "bmw":
		return &BMW{}
	}
	return nil
}