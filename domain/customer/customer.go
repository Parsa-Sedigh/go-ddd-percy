// Package customer holds our aggregates that combines many entities into a full object
package customer

import (
	"errors"
	tavern "github.com/Parsa-Sedigh/go-ddd-percy"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid name")
)

type Customer struct {
	/* person is the root entity of the customer which means that person.ID is the main identifier for the customer */
	person       *tavern.Person
	products     []*tavern.Item
	transactions []tavern.Transaction
}

// NewCustomer is a factory to create a new customer aggregate. It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &tavern.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]tavern.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
