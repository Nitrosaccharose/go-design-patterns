package simplefactorypattern

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	factory := new(Factory)
	fmt.Println("Created FruitFactory")
	fruit := factory.CreateProduct("apple")
	fmt.Println("Incoming parameter \"apple\"")
	fruit.show()
}

func Test2(t *testing.T) {
	factory := new(Factory)
	fmt.Println("Created FruitFactory")
	fruit := factory.CreateProduct("banana")
	fmt.Println("Incoming parameter \"banana\"")
	fruit.show()
}

func Test3(t *testing.T) {
	factory := new(Factory)
	fmt.Println("Created FruitFactory")
	fruit := factory.CreateProduct("pear")
	fmt.Println("Incoming parameter \"pear\"")
	fruit.show()
}
