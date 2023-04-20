package simplefactorypattern

import "fmt"

// Fruit 水果
type Fruit interface {
	show()
}

// Apple 苹果
type Apple struct {
	Fruit
}

func (*Apple) show() {
	fmt.Println("我是apple")
}

// Banana 香蕉
type Banana struct {
}

func (*Banana) show() {
	fmt.Println("我是banana")
}

// Pear 梨
type Pear struct {
}

func (*Pear) show() {
	fmt.Println("我是pear")
}

type Factory struct {
}

func (fac *Factory) CreateProduct(kind string) Fruit {
	var fruit Fruit
	switch kind {
	case "apple":
		fruit = new(Apple)
	case "banana":
		fruit = new(Banana)
	case "pear":
		fruit = new(Pear)
	}
	return fruit
}
