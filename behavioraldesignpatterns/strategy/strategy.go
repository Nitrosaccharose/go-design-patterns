package strategy

import "fmt"

type SellStrategy interface {
	getPrice(price float64) float64
}

type SellStrategyA struct {
	SellStrategy
}

func (s *SellStrategyA) getPrice(price float64) float64 {
	fmt.Println("销售策略A 打66折")
	return price * 0.66
}

type SellStrategyB struct {
	SellStrategy
}

func (s *SellStrategyB) getPrice(price float64) float64 {
	fmt.Println("销售策略B 满100减50")
	if price >= 100 {
		return price - 50
	}
	return price
}

type Salesman struct {
	sellStrategy SellStrategy
}

func (s *Salesman) SetSellStrategy(strategy SellStrategy) {
	s.sellStrategy = strategy
}
