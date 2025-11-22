package services

import (
	"testing"

	"github.com/vijayraghavareddy/tax-calculation/models"
)

func TestCalculateTax_Success(t *testing.T) {
	service := NewTaxService()

	req := &models.TaxRequest{
		Address: models.Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			Country: "US",
			ZipCode: "10001",
		},
		Items: []models.Item{
			{
				ID:       "item1",
				Name:     "Product A",
				Price:    100.00,
				Quantity: 2,
			},
			{
				ID:       "item2",
				Name:     "Product B",
				Price:    50.00,
				Quantity: 1,
			},
		},
	}

	resp, err := service.CalculateTax(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp == nil {
		t.Fatal("Expected response, got nil")
	}

	// Check subtotal
	expectedSubtotal := 250.00
	if resp.Subtotal != expectedSubtotal {
		t.Errorf("Expected subtotal %f, got %f", expectedSubtotal, resp.Subtotal)
	}

	// Check that tax was calculated
	if resp.TotalTax <= 0 {
		t.Error("Expected tax to be greater than 0")
	}

	// Check grand total
	if resp.GrandTotal != resp.Subtotal+resp.TotalTax {
		t.Error("Grand total should equal subtotal + tax")
	}

	// Check items count
	if len(resp.Items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(resp.Items))
	}

	// Check jurisdiction
	if resp.TaxJurisdiction == "" {
		t.Error("Expected tax jurisdiction to be set")
	}
}

func TestCalculateTax_MissingState(t *testing.T) {
	service := NewTaxService()

	req := &models.TaxRequest{
		Address: models.Address{
			Country: "US",
			ZipCode: "10001",
		},
		Items: []models.Item{
			{
				ID:       "item1",
				Name:     "Product A",
				Price:    100.00,
				Quantity: 1,
			},
		},
	}

	_, err := service.CalculateTax(req)

	if err == nil {
		t.Fatal("Expected error for missing state, got nil")
	}
}

func TestCalculateTax_MissingZipCode(t *testing.T) {
	service := NewTaxService()

	req := &models.TaxRequest{
		Address: models.Address{
			State:   "NY",
			Country: "US",
		},
		Items: []models.Item{
			{
				ID:       "item1",
				Name:     "Product A",
				Price:    100.00,
				Quantity: 1,
			},
		},
	}

	_, err := service.CalculateTax(req)

	if err == nil {
		t.Fatal("Expected error for missing zipcode, got nil")
	}
}

func TestCalculateTax_NoItems(t *testing.T) {
	service := NewTaxService()

	req := &models.TaxRequest{
		Address: models.Address{
			Country: "US",
			ZipCode: "10001",
		},
		Items: []models.Item{},
	}

	_, err := service.CalculateTax(req)

	if err == nil {
		t.Fatal("Expected error for no items, got nil")
	}
}

func TestCalculateTax_NegativePrice(t *testing.T) {
	service := NewTaxService()

	req := &models.TaxRequest{
		Address: models.Address{
			Country: "US",
			ZipCode: "10001",
		},
		Items: []models.Item{
			{
				ID:       "item1",
				Name:     "Product A",
				Price:    -100.00,
				Quantity: 1,
			},
		},
	}

	_, err := service.CalculateTax(req)

	if err == nil {
		t.Fatal("Expected error for negative price, got nil")
	}
}

func TestCalculateTax_InvalidQuantity(t *testing.T) {
	service := NewTaxService()

	req := &models.TaxRequest{
		Address: models.Address{
			Country: "US",
			ZipCode: "10001",
		},
		Items: []models.Item{
			{
				ID:       "item1",
				Name:     "Product A",
				Price:    100.00,
				Quantity: 0,
			},
		},
	}

	_, err := service.CalculateTax(req)

	if err == nil {
		t.Fatal("Expected error for invalid quantity, got nil")
	}
}

func TestCalculateTax_DifferentStates(t *testing.T) {
	service := NewTaxService()

	states := []string{"NY", "CA", "TX", "FL", "IL", "PA", "OH", "MI"}

	for _, state := range states {
		req := &models.TaxRequest{
			Address: models.Address{
				State:   state,
				Country: "US",
				ZipCode: "12345",
			},
			Items: []models.Item{
				{
					ID:       "item1",
					Name:     "Product A",
					Price:    100.00,
					Quantity: 1,
				},
			},
		}

		resp, err := service.CalculateTax(req)

		if err != nil {
			t.Errorf("Expected no error for state %s, got %v", state, err)
			continue
		}

		if resp.TotalTax < 0 {
			t.Errorf("Expected tax to be 0 or greater for state %s", state)
		}
	}
}

func TestCalculateTax_MultipleItems(t *testing.T) {
	service := NewTaxService()

	req := &models.TaxRequest{
		Address: models.Address{
			Street:  "456 Oak Ave",
			City:    "Los Angeles",
			State:   "CA",
			Country: "US",
			ZipCode: "90001",
		},
		Items: []models.Item{
			{
				ID:       "item1",
				Name:     "Product A",
				Price:    100.00,
				Quantity: 1,
			},
			{
				ID:       "item2",
				Name:     "Product B",
				Price:    50.00,
				Quantity: 2,
			},
			{
				ID:       "item3",
				Name:     "Product C",
				Price:    75.00,
				Quantity: 3,
			},
		},
	}

	resp, err := service.CalculateTax(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedSubtotal := 100.00 + 100.00 + 225.00
	if resp.Subtotal != expectedSubtotal {
		t.Errorf("Expected subtotal %f, got %f", expectedSubtotal, resp.Subtotal)
	}

	if len(resp.Items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(resp.Items))
	}
}

func TestGetTaxRateForLocation(t *testing.T) {
	service := NewTaxService()

	tests := []struct {
		state   string
		minRate float64
		maxRate float64
	}{
		{"NY", 0.08, 0.09},
		{"CA", 0.08, 0.09},
		{"TX", 0.08, 0.09},
		{"FL", 0.06, 0.08},
		{"DE", 0.00, 0.00}, // No sales tax
		{"MT", 0.00, 0.00}, // No sales tax
		{"OR", 0.00, 0.00}, // No sales tax
	}

	for _, tt := range tests {
		address := &models.Address{
			State:   tt.state,
			Country: "US",
			ZipCode: "12345",
		}

		rate := service.getTaxRateForLocation(address)

		if rate < tt.minRate || rate > tt.maxRate {
			t.Errorf("Tax rate for %s (%f) is outside expected range [%f, %f]",
				tt.state, rate, tt.minRate, tt.maxRate)
		}
	}
}

func TestRoundToTwoDecimals(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{10.123, 10.12},
		{10.126, 10.13},
		{10.999, 11.00},
		{10.001, 10.00},
		{10.50, 10.50},
	}

	for _, tt := range tests {
		result := roundToTwoDecimals(tt.input)
		if result != tt.expected {
			t.Errorf("roundToTwoDecimals(%f) = %f, expected %f",
				tt.input, result, tt.expected)
		}
	}
}
