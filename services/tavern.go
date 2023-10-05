package services

import (
	"github.com/google/uuid"
	"log"
)

// TavernConfiguration accepts a pointer to Tavern, in order to modify it
type TavernConfiguration func(t *Tavern) error

// Tavern holds the order service because we want to take orders inside of our tavern
type Tavern struct {
	OrderService *OrderService

	// We need BillingService to accept payments
	// TODO: Create this service
	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

// WithOrderService accepts a fully configured order service
func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os

		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("\nBill the customer: %0.0f\n", price)

	// TODO: Implement a billing service

	return nil
}
