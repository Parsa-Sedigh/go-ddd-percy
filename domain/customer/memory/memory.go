// Package memory is a in-memory implementation of CustomerRepository.
package memory

import (
	"fmt"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/customer"
	"github.com/google/uuid"
	"sync"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryCustomerRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	// the customer domain has an error for when it's not found and the repository is using that error.
	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryCustomerRepository) Add(c customer.Customer) error {
	/* Note: The factory function(New function of MemoryRepository) should have protected us against an empty customers map. So this
	if block is kinda unnecessary. But for extra safety, we make sure that the map is not nil and if it is, create the map. */
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}

	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryCustomerRepository) Update(c customer.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}
