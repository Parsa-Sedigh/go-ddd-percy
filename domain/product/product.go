package product

import (
	"errors"
	tavern "github.com/Parsa-Sedigh/go-ddd-percy"
	"github.com/google/uuid"
)

var (
	ErrMissingValues = errors.New("missing important values")
)

type Product struct {
	// root entity
	item     *tavern.Item
	price    float64
	quantity int
}

// factory function
func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &tavern.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

// some functions to expose the information
func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

// GetItem extracts the entity(item) from the product aggregate
func (p Product) GetItem() *tavern.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
