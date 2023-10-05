package customer_test

import (
	"errors"
	"github.com/Parsa-Sedigh/go-ddd-percy/domain/customer"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{

		// one test case is OFC when we have an empty name. We want to know that the expected error is returned
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: customer.ErrInvalidPerson,
		},
		// always test the happy path
		{
			test:        "Valid name",
			name:        "Parsa Sedigh",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			// execute a unit test for each test case

			_, err := customer.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
