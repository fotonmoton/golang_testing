package warehouse

import (
	"log"
	"strings"
)

type InMemoryState struct {
	products              []Product
	customerNotifications []CustomerNotification
	subscriptions         []CustomerSubscription
	customers             []Customer
	productIdGenerator    int
}

func NewInMemoryState() *InMemoryState {
	return &InMemoryState{productIdGenerator: 1}
}

func (s *InMemoryState) SaveProduct(p Product) Product {
	p.ID = s.productIdGenerator
	s.productIdGenerator++
	s.products = append(s.products, p)

	return p
}

func (s *InMemoryState) ListProducts() []Product {
	return s.products
}

func (s *InMemoryState) NotifyCustomers(notifications []CustomerNotification) {
	s.customerNotifications = append(s.customerNotifications, notifications...)
	log.Println(s.customerNotifications)
}

func (s *InMemoryState) FindSubscriptions(productTitle string) []CustomerSubscription {
	found := []CustomerSubscription{}

	for idx := range s.subscriptions {
		if strings.Contains(s.subscriptions[idx].ProductTitle, productTitle) {
			found = append(found, s.subscriptions[idx])
		}
	}

	return found
}

func (s *InMemoryState) AddSubscriptions(subs []CustomerSubscription) {
	s.subscriptions = append(s.subscriptions, subs...)
}

func (s *InMemoryState) AddCustomers(customers []Customer) {
	s.customers = append(s.customers, customers...)
}
