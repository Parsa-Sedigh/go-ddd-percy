// Package memory is an in-memory implementation of the ProductRepository.
package memory

import (
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/product"
	"github.com/google/uuid"
	"sync"
)

type MemoryProductRepository struct {
	products   map[uuid.UUID]product.Product
	sync.Mutex // protect from concurrent writes
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product

	for _, product := range mpr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}

	return product.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(newProduct product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	mpr.products[newProduct.GetID()] = newProduct

	return nil
}

func (mpr *MemoryProductRepository) Update(update product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[update.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[update.GetID()] = update

	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mpr.products, id)

	return nil
}
