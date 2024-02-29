package main

import "fmt"

type DeliveryVisitor interface {
	VisitHomeDelivery(*OrderHome)
	VisitPickup(*OrderPickup)
}

type DeliveryService struct{}

func (v *DeliveryService) VisitHomeDelivery(e *OrderHome) {
	fmt.Println("Visitor visits OrderHOme")
}

func (v *DeliveryService) VisitPickup(e *OrderPickup) {
	fmt.Println("Visitor visits Pickup")
}

// Element - интерфейс элемента
type Element interface {
	Accept(DeliveryVisitor)
}

type OrderHome struct{}

func (e *OrderHome) Accept(visitor DeliveryVisitor) {
	visitor.VisitHomeDelivery(e)
}

type OrderPickup struct{}

func (e *OrderPickup) Accept(visitor DeliveryVisitor) {
	visitor.VisitPickup(e)
}

// ObjectStructure - структура объектов
type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) Attach(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Accept(visitor DeliveryVisitor) {
	for _, e := range os.elements {
		e.Accept(visitor)
	}
}

func main() {
	visitor := &DeliveryService{}

	objectStructure := &ObjectStructure{}
	objectStructure.Attach(&OrderHome{})
	objectStructure.Attach(&OrderPickup{})

	objectStructure.Accept(visitor)
}
