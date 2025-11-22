package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vijayraghavareddy/tax-calculation/models"
)

func TestCalculateTax_ValidRequest(t *testing.T) {
	reqBody := models.TaxRequest{
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
		},
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate-tax", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CalculateTax(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var resp models.TaxResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if resp.Subtotal != 200.00 {
		t.Errorf("Expected subtotal 200.00, got %f", resp.Subtotal)
	}

	if resp.TotalTax <= 0 {
		t.Error("Expected tax to be greater than 0")
	}
}

func TestCalculateTax_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate-tax", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CalculateTax(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCalculateTax_MissingState(t *testing.T) {
	reqBody := models.TaxRequest{
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

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate-tax", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CalculateTax(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCalculateTax_WithPostalCode(t *testing.T) {
	reqBody := models.TaxRequest{
		Address: models.Address{
			Street:     "123 Main St",
			City:       "Boston",
			State:      "MA",
			Country:    "US",
			PostalCode: "02101",
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

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate-tax", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CalculateTax(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	w := httptest.NewRecorder()

	HealthCheck(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if resp["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got '%s'", resp["status"])
	}

	if resp["service"] != "tax-calculation-api" {
		t.Errorf("Expected service 'tax-calculation-api', got '%s'", resp["service"])
	}
}

func TestCalculateTax_EmptyItems(t *testing.T) {
	reqBody := models.TaxRequest{
		Address: models.Address{
			Country: "US",
			ZipCode: "10001",
		},
		Items: []models.Item{},
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate-tax", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CalculateTax(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	var resp models.ErrorResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected error code %d, got %d", http.StatusBadRequest, resp.Code)
	}
}
