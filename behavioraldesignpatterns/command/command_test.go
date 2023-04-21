package command

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	customer := &Customer{}
	waiter := &Waiter{}
	fmt.Println("顾客点了一份螺吸魂")
	waiter.setOrder(customer.createOrder("螺吸魂"))
	waiter.order.execute()
}
