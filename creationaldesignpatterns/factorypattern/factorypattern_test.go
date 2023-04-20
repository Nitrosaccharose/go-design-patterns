package factorypattern

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	factory := new(AppleFactory)
	fmt.Println("Created AppleFactory")
	fruit := factory.CreateProduct()
	fruit.show()
}

func Test2(t *testing.T) {
	factory := new(BananaFactory)
	fmt.Println("Created BananaFactory")
	fruit := factory.CreateProduct()
	fruit.show()
}

func Test3(t *testing.T) {
	factory := new(PearFactory)
	fmt.Println("Created PearFactory")
	fruit := factory.CreateProduct()
	fruit.show()
}
