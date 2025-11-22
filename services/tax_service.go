package services

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/vijayraghavareddy/tax-calculation/models"
)

// TaxService handles tax calculation logic
type TaxService struct {
	rand *rand.Rand
}

// NewTaxService creates a new instance of TaxService
func NewTaxService() *TaxService {
	source := rand.NewSource(time.Now().UnixNano())
	return &TaxService{
		rand: rand.New(source),
	}
}

// CalculateTax calculates tax for the given request
func (s *TaxService) CalculateTax(req *models.TaxRequest) (*models.TaxResponse, error) {
	if err := s.validateRequest(req); err != nil {
		return nil, err
	}

	// Get tax rate based on location
	taxRate := s.getTaxRateForLocation(&req.Address)
	jurisdiction := s.getTaxJurisdiction(&req.Address)

	var itemDetails []models.ItemTaxDetail
	var subtotal, totalTax float64

	// Calculate tax for each item
	for _, item := range req.Items {
		itemSubtotal := item.Price * float64(item.Quantity)
		itemTax := itemSubtotal * taxRate
		itemTotal := itemSubtotal + itemTax

		detail := models.ItemTaxDetail{
			ItemID:      item.ID,
			ItemName:    item.Name,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Subtotal:    roundToTwoDecimals(itemSubtotal),
			TaxRate:     roundToTwoDecimals(taxRate * 100), // Convert to percentage
			TaxAmount:   roundToTwoDecimals(itemTax),
			TotalAmount: roundToTwoDecimals(itemTotal),
		}

		itemDetails = append(itemDetails, detail)
		subtotal += itemSubtotal
		totalTax += itemTax
	}

	response := &models.TaxResponse{
		Address:         req.Address,
		Items:           itemDetails,
		Subtotal:        roundToTwoDecimals(subtotal),
		TotalTax:        roundToTwoDecimals(totalTax),
		GrandTotal:      roundToTwoDecimals(subtotal + totalTax),
		TaxJurisdiction: jurisdiction,
	}

	return response, nil
}

// validateRequest validates the tax calculation request
func (s *TaxService) validateRequest(req *models.TaxRequest) error {
	if req.Address.State == "" {
		return fmt.Errorf("state is required")
	}
	if req.Address.ZipCode == "" && req.Address.PostalCode == "" {
		return fmt.Errorf("zipcode is required")
	}
	if len(req.Items) == 0 {
		return fmt.Errorf("at least one item is required")
	}

	for i, item := range req.Items {
		if item.Price < 0 {
			return fmt.Errorf("item %d has invalid price", i)
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("item %d has invalid quantity", i)
		}
	}

	return nil
}

// getTaxRateForLocation returns a tax rate based on location
// Routes to appropriate country-specific tax calculation
func (s *TaxService) getTaxRateForLocation(address *models.Address) float64 {
	country := strings.ToUpper(address.Country)
	
	// Route to country-specific tax calculation
	if country == "IN" || country == "INDIA" {
		return s.getIndianStateTaxRate(address)
	}
	
	// Default to US tax calculation
	return s.getUSTaxRate(address)
}

// getUSTaxRate returns a tax rate based on the US state
// Rates are approximate and based on combined state and average local rates
func (s *TaxService) getUSTaxRate(address *models.Address) float64 {
	state := strings.ToUpper(address.State)

	// US state sales tax rates (approximate combined rates)
	var baseRate float64
	switch state {
	case "AL": // Alabama
		baseRate = 0.0913
	case "AK": // Alaska
		baseRate = 0.0176
	case "AZ": // Arizona
		baseRate = 0.0831
	case "AR": // Arkansas
		baseRate = 0.0947
	case "CA": // California
		baseRate = 0.0850
	case "CO": // Colorado
		baseRate = 0.0763
	case "CT": // Connecticut
		baseRate = 0.0635
	case "DE": // Delaware
		baseRate = 0.0000 // No sales tax
	case "FL": // Florida
		baseRate = 0.0705
	case "GA": // Georgia
		baseRate = 0.0733
	case "HI": // Hawaii
		baseRate = 0.0444
	case "ID": // Idaho
		baseRate = 0.0602
	case "IL": // Illinois
		baseRate = 0.0868
	case "IN": // Indiana
		baseRate = 0.0700
	case "IA": // Iowa
		baseRate = 0.0694
	case "KS": // Kansas
		baseRate = 0.0865
	case "KY": // Kentucky
		baseRate = 0.0600
	case "LA": // Louisiana
		baseRate = 0.0952
	case "ME": // Maine
		baseRate = 0.0550
	case "MD": // Maryland
		baseRate = 0.0600
	case "MA": // Massachusetts
		baseRate = 0.0625
	case "MI": // Michigan
		baseRate = 0.0600
	case "MN": // Minnesota
		baseRate = 0.0744
	case "MS": // Mississippi
		baseRate = 0.0707
	case "MO": // Missouri
		baseRate = 0.0824
	case "MT": // Montana
		baseRate = 0.0000 // No sales tax
	case "NE": // Nebraska
		baseRate = 0.0694
	case "NV": // Nevada
		baseRate = 0.0823
	case "NH": // New Hampshire
		baseRate = 0.0000 // No sales tax
	case "NJ": // New Jersey
		baseRate = 0.0663
	case "NM": // New Mexico
		baseRate = 0.0779
	case "NY": // New York
		baseRate = 0.0852
	case "NC": // North Carolina
		baseRate = 0.0698
	case "ND": // North Dakota
		baseRate = 0.0696
	case "OH": // Ohio
		baseRate = 0.0723
	case "OK": // Oklahoma
		baseRate = 0.0897
	case "OR": // Oregon
		baseRate = 0.0000 // No sales tax
	case "PA": // Pennsylvania
		baseRate = 0.0634
	case "RI": // Rhode Island
		baseRate = 0.0700
	case "SC": // South Carolina
		baseRate = 0.0744
	case "SD": // South Dakota
		baseRate = 0.0645
	case "TN": // Tennessee
		baseRate = 0.0955
	case "TX": // Texas
		baseRate = 0.0820
	case "UT": // Utah
		baseRate = 0.0719
	case "VT": // Vermont
		baseRate = 0.0624
	case "VA": // Virginia
		baseRate = 0.0575
	case "WA": // Washington
		baseRate = 0.0920
	case "WV": // West Virginia
		baseRate = 0.0650
	case "WI": // Wisconsin
		baseRate = 0.0543
	case "WY": // Wyoming
		baseRate = 0.0536
	default:
		// Default rate if state not recognized
		baseRate = 0.0700
	}

	return baseRate
}

// getIndianStateTaxRate returns GST rate for Indian states
// India uses a centralized GST system with standard rates: 5%, 12%, 18%, 28%
// Using 18% as default (most common rate for goods and services)
func (s *TaxService) getIndianStateTaxRate(address *models.Address) float64 {
	state := strings.ToUpper(address.State)

	// Indian state GST rates (CGST + SGST = Total GST)
	// Most goods/services fall under 18% GST (9% CGST + 9% SGST)
	var gstRate float64
	switch state {
	case "MH", "MAHARASHTRA":
		gstRate = 0.18 // 18% GST
	case "KA", "KARNATAKA":
		gstRate = 0.18 // 18% GST
	case "TN", "TAMIL NADU", "TAMILNADU":
		gstRate = 0.18 // 18% GST
	case "DL", "DELHI":
		gstRate = 0.18 // 18% GST
	case "GJ", "GUJARAT":
		gstRate = 0.18 // 18% GST
	case "WB", "WEST BENGAL", "WESTBENGAL":
		gstRate = 0.18 // 18% GST
	case "RJ", "RAJASTHAN":
		gstRate = 0.18 // 18% GST
	case "UP", "UTTAR PRADESH", "UTTARPRADESH":
		gstRate = 0.18 // 18% GST
	case "MP", "MADHYA PRADESH", "MADHYAPRADESH":
		gstRate = 0.18 // 18% GST
	case "AP", "ANDHRA PRADESH", "ANDHRAPRADESH":
		gstRate = 0.18 // 18% GST
	case "TS", "TELANGANA":
		gstRate = 0.18 // 18% GST
	case "BR", "BIHAR":
		gstRate = 0.18 // 18% GST
	case "HR", "HARYANA":
		gstRate = 0.18 // 18% GST
	case "PB", "PUNJAB":
		gstRate = 0.18 // 18% GST
	case "KL", "KERALA":
		gstRate = 0.18 // 18% GST
	case "OR", "ODISHA", "ORISSA":
		gstRate = 0.18 // 18% GST
	case "JH", "JHARKHAND":
		gstRate = 0.18 // 18% GST
	case "AS", "ASSAM":
		gstRate = 0.18 // 18% GST
	case "CT", "CHHATTISGARH":
		gstRate = 0.18 // 18% GST
	case "UK", "UTTARAKHAND":
		gstRate = 0.18 // 18% GST
	case "HP", "HIMACHAL PRADESH", "HIMACHALPRADESH":
		gstRate = 0.18 // 18% GST
	case "JK", "JAMMU AND KASHMIR", "JAMMUANDKASHMIR":
		gstRate = 0.18 // 18% GST
	case "GA", "GOA":
		gstRate = 0.18 // 18% GST
	default:
		// Default GST rate for unrecognized states
		gstRate = 0.18
	}

	return gstRate
}

// getTaxJurisdiction returns the tax jurisdiction string
func (s *TaxService) getTaxJurisdiction(address *models.Address) string {
	country := strings.ToUpper(address.Country)
	
	if country == "IN" || country == "INDIA" {
		return fmt.Sprintf("%s, India (GST)", address.State)
	}
	
	return fmt.Sprintf("%s, USA", address.State)
}

// roundToTwoDecimals rounds a float64 to 2 decimal places
func roundToTwoDecimals(value float64) float64 {
	return float64(int(value*100+0.5)) / 100
}
