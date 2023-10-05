package aggregate

import (
	"errors"
	"github.com/Parsa-Sedigh/go-ddd-percy/entity"
	"github.com/google/uuid"
)

var (
	ErrMissingValues = errors.New("missing important values")
)

type Product struct {
	// root entity
	item     *entity.Item
	price    float64
	quantity int
}

// factory function
func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
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
func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
