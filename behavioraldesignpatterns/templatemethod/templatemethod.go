package templatemethod

import "fmt"

// BeverageTemplate 饮品模板
type BeverageTemplate interface {
	addWater()
	addMilk()
	addCoffeeLiquid()
}

// Americano 美式咖啡
type Americano struct {
	BeverageTemplate
}

func (a *Americano) addWater() {
	fmt.Println("加足量水...")
}

func (a *Americano) addMilk() {
	fmt.Println("不加牛奶...")
}

func (a *Americano) addCoffeeLiquid() {
	fmt.Println("加咖啡液...")
}

// Cappuccino 卡布奇诺
type Cappuccino struct {
	BeverageTemplate
}

func (c *Cappuccino) addWater() {
	fmt.Println("加少量水...")
}

func (c *Cappuccino) addMilk() {
	fmt.Println("加牛奶...")
}

func (c *Cappuccino) addCoffeeLiquid() {
	fmt.Println("加咖啡液...")
}
