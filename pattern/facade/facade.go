package main

import "fmt"

type OrderService struct{}

func (s *OrderService) Cart() string {
	return "OrderService: формирование корзины"
}

type PaymentService struct{}

func (s *PaymentService) ProcessPayment() string {
	return "PaymentService: выполнение платежа"
}

// Фасад - упрощенный интерфейс для работы с подсистемами
type OnlineStore struct {
	subsystemA *OrderService
	subsystemB *PaymentService
}

func NewOnlineStoreFacade() *OnlineStore {
	return &OnlineStore{
		subsystemA: &OrderService{},
		subsystemB: &PaymentService{},
	}
}

func (f *OnlineStore) Purchase() string {
	result := "Facade: Операции подсистем\n"
	result += f.subsystemA.Cart()
	result += f.subsystemB.ProcessPayment()
	return result
}

func main() {
	facade := NewOnlineStoreFacade()
	result := facade.Purchase()
	fmt.Println(result)
}
