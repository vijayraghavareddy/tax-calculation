package models

// Address represents the customer's address for tax calculation
type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	ZipCode    string `json:"zipcode"`
	PostalCode string `json:"postal_code,omitempty"` // Alternative field name
}

// Item represents a product or service to be taxed
type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

// TaxRequest represents the incoming request for tax calculation
type TaxRequest struct {
	Address Address `json:"address"`
	Items   []Item  `json:"items"`
}

// ItemTaxDetail represents tax details for a single item
type ItemTaxDetail struct {
	ItemID      string  `json:"item_id"`
	ItemName    string  `json:"item_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
	TaxRate     float64 `json:"tax_rate"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount"`
}

// TaxResponse represents the response with calculated taxes
type TaxResponse struct {
	Address         Address         `json:"address"`
	Items           []ItemTaxDetail `json:"items"`
	Subtotal        float64         `json:"subtotal"`
	TotalTax        float64         `json:"total_tax"`
	GrandTotal      float64         `json:"grand_total"`
	TaxJurisdiction string          `json:"tax_jurisdiction"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
