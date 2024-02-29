package main

import "fmt"

// Order - структура объекта, который мы строим
type Order struct {
	name   string
	ID     int
	isAble bool
}

// Builder - интерфейс строителя
type Builder interface {
	BuildName()
	BuildID()
	BuildAbility()
	GetOrder() *Order
}

// ConcreteBuilder - конкретная реализация строителя
type ConcreteBuilder struct {
	order Order
}

func (b *ConcreteBuilder) BuildName() {
	b.order.name = "Part 1"
}

func (b *ConcreteBuilder) BuildID() {
	b.order.ID = 42
}

func (b *ConcreteBuilder) BuildAbility() {
	b.order.isAble = true
}

func (b *ConcreteBuilder) GetOrder() *Order {
	return &b.order
}

// Director - директор, который управляет процессом строительства
type Director struct {
	builder Builder
}

func (d *Director) Construct() *Order {
	d.builder.BuildName()
	d.builder.BuildID()
	d.builder.BuildAbility()
	return d.builder.GetOrder()
}

func main() {
	builder := &ConcreteBuilder{}
	director := &Director{builder: builder}
	order := director.Construct()

	fmt.Printf("Part 1: %s\n", order.name)
	fmt.Printf("Part 2: %d\n", order.ID)
	fmt.Printf("Part 3: %t\n", order.isAble)
}
