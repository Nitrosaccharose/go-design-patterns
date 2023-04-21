package command

import "fmt"

// Customer 顾客
type Customer struct {
}

func (c *Customer) createOrder(dishName string) *Order {
	return &Order{dishName: dishName}
}

// Waiter  服务员
type Waiter struct {
	order *Order
}

func (w *Waiter) setOrder(order *Order) {
	w.order = order
}

// Cook 厨师
type Cook struct {
}

func (c *Cook) cook(dishName string) {
	fmt.Println("正在制作 " + dishName + "...")
}

// Order 订单
type Order struct {
	dishName string
}

func (o *Order) execute() {
	cook := &Cook{}
	cook.cook(o.dishName)
}
