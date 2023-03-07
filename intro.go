package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func (o *Order) SetPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do setPrice: ", o.Price)
}

func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func NewOrder() *Order {
	return &Order{}
}

func main() {
	fmt.Println("Order")
	order := NewOrder()
	order.ID = "123"
	order.Quantity = 5
	order.SetPrice(20.0)
	fmt.Println("Price total: ", order.getTotal())
}
