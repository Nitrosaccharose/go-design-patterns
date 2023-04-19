package factorypattern

import "fmt"

// Fruit 水果
type Fruit interface {
	show()
}
type AbstractFactory interface {
	CreateProduct() Fruit
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

type AppleFactory struct {
	AbstractFactory
}

func (*AppleFactory) CreateProduct() Fruit {
	return new(Apple)
}

type BananaFactory struct {
	AbstractFactory
}

func (*BananaFactory) CreateProduct() Fruit {
	return new(Banana)
}

type PearFactory struct {
	AbstractFactory
}

func (*PearFactory) CreateProduct() Fruit {
	return new(Pear)
}

type Factory struct {
	AbstractFactory
}
