package strategy

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	salesman := Salesman{}
	salesman.SetSellStrategy(&SellStrategyA{})
	price := 102.9
	fmt.Println("原价：", price)
	fmt.Println("A策略报价：", salesman.sellStrategy.getPrice(price))
	salesman.SetSellStrategy(&SellStrategyB{})
	fmt.Println("B策略报价：", salesman.sellStrategy.getPrice(price))
}

func Test2(t *testing.T) {
	salesman := Salesman{}
	salesman.SetSellStrategy(&SellStrategyA{})
	price := 531.2
	fmt.Println("原价:", price)
	fmt.Println("A策略报价：", salesman.sellStrategy.getPrice(price))
	salesman.SetSellStrategy(&SellStrategyB{})
	fmt.Println("B策略报价：", salesman.sellStrategy.getPrice(price))
}
