package templatemethod

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("Americano:")
	Americano := &Americano{}
	Americano.addWater()
	Americano.addMilk()
	Americano.addCoffeeLiquid()
}

func Test2(t *testing.T) {
	fmt.Println("Cappuccino:")
	Cappuccino := &Cappuccino{}
	Cappuccino.addWater()
	Cappuccino.addMilk()
	Cappuccino.addCoffeeLiquid()
}
