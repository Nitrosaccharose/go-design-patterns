package decorator

import (
	"fmt"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	hamburger := &Hamburger{}
	hamburgerWithMeat := &Meat{JardiniereDecorator{food: hamburger}}
	hamburgerWithCheese := &Cheese{JardiniereDecorator{food: hamburger}}
	hamburgerWithMeatAndCheese := &Cheese{JardiniereDecorator{food: hamburgerWithMeat}}
	fmt.Println("hamburger price is " + strconv.Itoa(hamburger.getPrice()))
	fmt.Println("hamburgerWithMeat price is " + strconv.Itoa(hamburgerWithMeat.getPrice()))
	fmt.Println("hamburgerWithCheese price is " + strconv.Itoa(hamburgerWithCheese.getPrice()))
	fmt.Println("hamburgerWithMeatAndCheese price is " + strconv.Itoa(hamburgerWithMeatAndCheese.getPrice()))
}

func Test2(t *testing.T) {
	sandwich := &Sandwich{}
	sandwichWithMeat := &Meat{JardiniereDecorator{food: sandwich}}
	sandwichWithCheese := &Cheese{JardiniereDecorator{food: sandwich}}
	sandwichWithMeatAndCheese := &Cheese{JardiniereDecorator{food: sandwichWithMeat}}
	fmt.Println("sandwich price is " + strconv.Itoa(sandwich.getPrice()))
	fmt.Println("sandwichWithMeat price is " + strconv.Itoa(sandwichWithMeat.getPrice()))
	fmt.Println("sandwichWithCheese price is " + strconv.Itoa(sandwichWithCheese.getPrice()))
	fmt.Println("sandwichWithMeatAndCheese price is " + strconv.Itoa(sandwichWithMeatAndCheese.getPrice()))
}
