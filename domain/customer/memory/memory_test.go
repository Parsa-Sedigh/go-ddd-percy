package memory

import (
	"errors"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/customer"
	"github.com/google/uuid"
	"testing"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// create a new customer
	cust, err := customer.NewCustomer("parsa")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	// we're not gonna use the factory function here
	repo := MemoryCustomerRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse(""),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
