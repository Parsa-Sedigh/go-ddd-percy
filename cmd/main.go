package main

import (
	"context"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/product"
	"github.com/Parsa-Sedigh/go-ddd-percy/services/order"
	"github.com/Parsa-Sedigh/go-ddd-percy/services/tavern"
	"github.com/google/uuid"
)

func main() {
	products := productInventory()

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	tavernService, err := tavern.NewTavern(tavern.WithOrderService(os))
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Parsa")
	if err != nil {
		panic(err)
	}

	orderUUID := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavernService.Order(uid, orderUUID)
	if err != nil {
		panic(err)
	}
}

// in a real app, you should load the inventory from product repository or sth
func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy beverage", 1.99)
	if err != nil {
		panic(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		panic(err)
	}

	wine, err := product.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		panic(err)
	}

	return []product.Product{beer, peanuts, wine}
}
