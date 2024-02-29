package main

import (
	"fmt"
)

// State определяет интерфейс для всех состояний заказа.
type State interface {
	Process(order *Order)
	Undo(order *Order)
}

// Order представляет заказ.
type Order struct {
	state State
}

// NewOrder создает новый заказ и устанавливает начальное состояние.
func NewOrder() *Order {
	return &Order{state: &ProcessingState{}}
}

// Process начинает обработку заказа.
func (o *Order) Process() {
	o.state.Process(o)
}

// SetState устанавливает новое состояние заказа.
func (o *Order) SetState(state State) {
	o.state = state
}

// ProcessingState представляет состояние обработки заказа.
type ProcessingState struct{}

func (s *ProcessingState) Process(order *Order) {
	fmt.Println("Order is being processed...")
	// Simulate processing
	order.SetState(&ShippedState{})
}

func (s *ProcessingState) Undo(order *Order) {
	fmt.Println("Order is being undo...")
	// Simulate processing
	order.SetState(&ProcessingState{})
}

// ShippedState представляет состояние, когда заказ отправлен.
type ShippedState struct{}

func (s *ShippedState) Process(order *Order) {
	fmt.Println("Order has been shipped!")
	// Simulate delivery
	order.SetState(&DeliveredState{})
}

func (s *ShippedState) Undo(order *Order) {
	fmt.Println("Order has been cancelled")
	// Simulate delivery
	order.SetState(&ProcessingState{})
}

// DeliveredState представляет состояние, когда заказ доставлен.
type DeliveredState struct{}

func (s *DeliveredState) Process(order *Order) {
	fmt.Println("Order has been delivered!")
	// End of process
}

func (s *DeliveredState) Undo(order *Order) {
	fmt.Println("Order has been cancelled")
	order.SetState(&ShippedState{})
}

func main() {
	order := NewOrder()
	order.Process()
	order.Process()
	order.Process()
}
