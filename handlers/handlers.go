package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vijayraghavareddy/tax-calculation/models"
	"github.com/vijayraghavareddy/tax-calculation/services"
)

var taxService = services.NewTaxService()

// CalculateTax handles POST requests to calculate tax
func CalculateTax(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.TaxRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Normalize postal_code to zipcode if provided
	if req.Address.ZipCode == "" && req.Address.PostalCode != "" {
		req.Address.ZipCode = req.Address.PostalCode
	}

	response, err := taxService.CalculateTax(&req)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HealthCheck handles GET requests for health check
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "healthy",
		"service": "tax-calculation-api",
		"version": "1.0.0",
	})
}

// sendErrorResponse sends an error response
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Code:    statusCode,
	})
}
