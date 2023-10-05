package order

import (
	"context"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/customer"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/customer/memory"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/customer/mongo"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/product"
	prodmem "github.com/Parsa-Sedigh/go-ddd-percy/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

/* The reason we take a pointer to the OrderService is because we want to modify the service */
type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository

	//billing billing.Service
}

// factory function
// an example of calling this function: NewOrderService(WithCustomerRepository, WithMemoryProductRepository). With this pattern,
// we pass functions to modify the service. Using this,it makes it easy to change the behavior of the service.
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// loop through all the cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// WithCustomerRepository applies a customer repository to the order service
func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	// return a function that matches OrderConfiguration alias(we need to return a function so that we can chain these)
	return func(os *OrderService) error {
		os.customers = cr

		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()

	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connStr string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connStr)
		if err != nil {
			return err
		}

		os.customers = cr

		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr

		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	/* Get each product, so we need a product repository and we would also need a WithMemoryProductRepository which sets the product repo
	for the order service. In order to have a product repo, we need a product aggregate. */

	var products []product.Product
	var total float64

	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return total, nil
}

// AddCustomer exposes the adding a customer functionality instead of exporting the customers field of OrderService.
// In other words, this method wraps the functionality of customers repo.
func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	if err = o.customers.Add(c); err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
