package order

import (
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/customer"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/product"
	"github.com/google/uuid"
	"testing"
)

// init_products is used to initialize tests easier
func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy beverage", 1.99)
	if err != nil {
		// we don't want to continue our tests if sth went wrong in the initialization
		t.Fatal(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{beer, peanuts, wine}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Parsa")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
