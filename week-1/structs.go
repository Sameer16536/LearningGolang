package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name       string
	age        int
	membership string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	Customer
}

func testStructs() {
	order := Order{
		id:        "1234",
		amount:    250.50,
		status:    "pending",
		createdAt: time.Now(),
		Customer: Customer{
			name:       "user",
			age:        5,
			membership: "Premium",
		},
	}

	order.Customer.name = "newUser"

	// var newOrder Order
	// newOrder.id = "5678"
	// newOrder.amount = 100.00
	// newOrder.status = "completed"
	// newOrder.createdAt = time.Now()

	// newOrder.changeStatus("Paid")

	fmt.Println(order)
	// fmt.Println(newOrder)

	newOrderPtr := newOrder("91011", 300.75, "processing")
	fmt.Println(newOrderPtr)
}

// Receiver type
func (o *Order) changeStatus(status string) {
	o.status = status
}

func newOrder(id string, amount float32, status string) *Order {

	order := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}

	return &order
}
