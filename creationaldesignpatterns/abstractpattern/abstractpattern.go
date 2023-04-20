package abstractpattern

import "fmt"

// Fruit 水果
type Fruit interface {
	show()
}

// Apple 苹果
type Apple interface {
	Fruit
}

// Banana 香蕉
type Banana interface {
	Fruit
}

// Pear 梨
type Pear interface {
	Fruit
}

// ChinaApple 中国苹果
type ChinaApple struct {
	Apple
}

func (*ChinaApple) show() {
	fmt.Println("我是china apple")
}

// ChinaBanana 中国香蕉
type ChinaBanana struct {
	Banana
}

func (*ChinaBanana) show() {
	fmt.Println("我是china banana")
}

// ChinaPear 中国梨
type ChinaPear struct {
	Pear
}

func (*ChinaPear) show() {
	fmt.Println("我是china pear")
}

// USAApple 美国苹果
type USAApple struct {
	Apple
}

func (*USAApple) show() {
	fmt.Println("我是usa apple")
}

// USABanana 美国香蕉
type USABanana struct {
	Banana
}

func (*USABanana) show() {
	fmt.Println("我是usa banana")
}

// USAPear 美国梨
type USAPear struct {
	Pear
}

func (*USAPear) show() {
	fmt.Println("我是usa pear")
}

// JapanApple 日本苹果
type JapanApple struct {
	Apple
}

func (*JapanApple) show() {
	fmt.Println("我是japan apple")
}

// JapanBanana 日本香蕉
type JapanBanana struct {
	Banana
}

func (*JapanBanana) show() {
	fmt.Println("我是japan banana")
}

// JapanPear 日本梨
type JapanPear struct {
	Pear
}

func (*JapanPear) show() {
	fmt.Println("我是japan pear")
}

type AbstractFactory interface {
	CreateApple() Fruit
	CreateBanana() Fruit
	CreatePear() Fruit
}

// ChinaFactory 中国工厂
type ChinaFactory struct {
	AbstractFactory
}

func (*ChinaFactory) CreateApple() Fruit {
	return new(ChinaApple)
}

func (*ChinaFactory) CreateBanana() Fruit {
	return new(ChinaBanana)
}

func (*ChinaFactory) CreatePear() Fruit {
	return new(ChinaPear)
}

// USAFactory 美国工厂
type USAFactory struct {
	AbstractFactory
}

func (*USAFactory) CreateApple() Fruit {
	return new(USAApple)
}

func (*USAFactory) CreateBanana() Fruit {
	return new(USABanana)
}

func (*USAFactory) CreatePear() Fruit {
	return new(USAPear)
}

// JapanFactory 日本工厂
type JapanFactory struct {
	AbstractFactory
}

func (*JapanFactory) CreateApple() Fruit {
	return new(JapanApple)
}

func (*JapanFactory) CreateBanana() Fruit {
	return new(JapanBanana)
}

func (*JapanFactory) CreatePear() Fruit {
	return new(JapanPear)
}
