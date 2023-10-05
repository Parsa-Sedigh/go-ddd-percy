package tavern

import (
	"context"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/product"
	"github.com/Parsa-Sedigh/go-ddd-percy/services/order"
	"github.com/google/uuid"
	"testing"
)

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

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		//WithMemoryCustomerRepository(),
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("Parsa")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}
}
