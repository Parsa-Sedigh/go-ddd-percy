package product

import (
	"errors"
	"github.com/Parsa-Sedigh/go-ddd-percy/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("no such product")
	ErrProductAlreadyExists = errors.New("there is already such a product")
)

// ProductRepository manages and handles the product aggregate
type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
