// Package aggregate holds our aggregates that combines many entities into a full object
package aggregate

import (
	"errors"
	"github.com/Parsa-Sedigh/go-ddd-percy/entity"
	"github.com/Parsa-Sedigh/go-ddd-percy/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid name")
)

type Customer struct {
	/* person is the root entity of the customer which means that person.ID is the main identifier for the customer */
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new customer aggregate. It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
