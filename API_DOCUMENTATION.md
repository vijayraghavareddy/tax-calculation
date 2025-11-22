# Tax Calculation API - Detailed Documentation

## Overview

The Tax Calculation API is a RESTful web service that calculates sales tax, VAT, or GST based on the customer's address and the items being purchased. The API uses location-based tax rates that simulate real-world tax calculations for various countries.

## Architecture

### Components

1. **Handlers** (`handlers/handlers.go`)
   - HTTP request/response handling
   - Request validation
   - Response formatting

2. **Services** (`services/tax_service.go`)
   - Business logic for tax calculation
   - Tax rate determination
   - Mathematical operations

3. **Models** (`models/models.go`)
   - Data structures for requests and responses
   - Type definitions

### Flow Diagram

```
Client Request
    ↓
HTTP Handler (handlers/handlers.go)
    ↓
Request Validation
    ↓
Tax Service (services/tax_service.go)
    ↓
Tax Calculation Logic
    ↓
Response Generation
    ↓
JSON Response to Client
```

## API Reference

### Base URL

```
http://localhost:8080/api/v1
```

### Authentication

Currently, the API does not require authentication. Future versions may include API key authentication.

---

## Endpoints

### 1. POST /calculate-tax

Calculate tax for a list of items based on the provided address.

#### Request

**Headers:**
```
Content-Type: application/json
```

**Body Schema:**

```json
{
  "address": {
    "street": "string (optional)",
    "city": "string (optional)",
    "state": "string (optional)",
    "country": "string (required)",
    "zipcode": "string (required*)",
    "postal_code": "string (required*)"
  },
  "items": [
    {
      "id": "string (required)",
      "name": "string (required)",
      "description": "string (optional)",
      "price": "number (required, >= 0)",
      "quantity": "integer (required, > 0)"
    }
  ]
}
```

*Note: Either `zipcode` or `postal_code` must be provided.

#### Response

**Success (200 OK):**

```json
{
  "address": {
    "street": "string",
    "city": "string",
    "state": "string",
    "country": "string",
    "zipcode": "string"
  },
  "items": [
    {
      "item_id": "string",
      "item_name": "string",
      "price": "number",
      "quantity": "integer",
      "subtotal": "number",
      "tax_rate": "number (percentage)",
      "tax_amount": "number",
      "total_amount": "number"
    }
  ],
  "subtotal": "number",
  "total_tax": "number",
  "grand_total": "number",
  "tax_jurisdiction": "string"
}
```

**Error (400 Bad Request):**

```json
{
  "error": "Bad Request",
  "message": "Detailed error message",
  "code": 400
}
```

#### Example Requests

**Example 1: US Purchase**

```bash
curl -X POST http://localhost:8080/api/v1/calculate-tax \
  -H "Content-Type: application/json" \
  -d '{
    "address": {
      "street": "1600 Amphitheatre Parkway",
      "city": "Mountain View",
      "state": "CA",
      "country": "US",
      "zipcode": "94043"
    },
    "items": [
      {
        "id": "LAPTOP-001",
        "name": "MacBook Pro",
        "description": "16-inch, M3 Pro",
        "price": 2499.00,
        "quantity": 1
      },
      {
        "id": "MOUSE-001",
        "name": "Magic Mouse",
        "price": 79.00,
        "quantity": 2
      }
    ]
  }'
```

**Example 2: UK Purchase with Postal Code**

```bash
curl -X POST http://localhost:8080/api/v1/calculate-tax \
  -H "Content-Type: application/json" \
  -d '{
    "address": {
      "street": "221B Baker Street",
      "city": "London",
      "country": "UK",
      "postal_code": "NW1 6XE"
    },
    "items": [
      {
        "id": "BOOK-001",
        "name": "Sherlock Holmes Complete Works",
        "price": 29.99,
        "quantity": 1
      }
    ]
  }'
```

**Example 3: Canadian Purchase**

```bash
curl -X POST http://localhost:8080/api/v1/calculate-tax \
  -H "Content-Type: application/json" \
  -d '{
    "address": {
      "street": "150 King St W",
      "city": "Toronto",
      "state": "ON",
      "country": "CA",
      "zipcode": "M5H1J9"
    },
    "items": [
      {
        "id": "COFFEE-001",
        "name": "Premium Coffee Beans",
        "price": 15.50,
        "quantity": 3
      },
      {
        "id": "MUG-001",
        "name": "Ceramic Coffee Mug",
        "price": 12.00,
        "quantity": 2
      }
    ]
  }'
```

---

### 2. GET /health

Health check endpoint to verify the API is running.

#### Request

No parameters required.

#### Response

**Success (200 OK):**

```json
{
  "status": "healthy",
  "service": "tax-calculation-api",
  "version": "1.0.0"
}
```

#### Example Request

```bash
curl http://localhost:8080/api/v1/health
```

---

## Tax Calculation Logic

### Tax Rate Determination

The API determines tax rates based on the country provided in the address. The rates are randomized within realistic ranges for each country:

| Country | Country Code | Tax Type | Rate Range |
|---------|--------------|----------|------------|
| United States | US, USA | Sales Tax | 5.0% - 12.0% |
| Canada | CA | GST/HST | 5.0% - 15.0% |
| United Kingdom | UK, GB | VAT | ~20.0% |
| Germany | DE | VAT | ~19.0% |
| France | FR | VAT | ~20.0% |
| India | IN | GST | 5.0% - 28.0% |
| Australia | AU | GST | ~10.0% |
| Japan | JP | Consumption Tax | ~10.0% |
| Other Countries | * | Standard Tax | 10.0% - 20.0% |

### Calculation Formula

For each item:
```
Item Subtotal = Price × Quantity
Item Tax = Item Subtotal × Tax Rate
Item Total = Item Subtotal + Item Tax
```

For the entire order:
```
Order Subtotal = Sum of all Item Subtotals
Total Tax = Sum of all Item Taxes
Grand Total = Order Subtotal + Total Tax
```

### Rounding

All monetary values are rounded to 2 decimal places using standard rounding rules (0.5 rounds up).

---

## Validation Rules

### Address Validation

1. **country** - Required, must not be empty
2. **zipcode or postal_code** - At least one is required
3. Other address fields (street, city, state) are optional

### Item Validation

1. **id** - Required, must not be empty
2. **name** - Required, must not be empty
3. **price** - Required, must be >= 0
4. **quantity** - Required, must be > 0
5. **description** - Optional
6. At least one item must be present in the request

---

## Error Handling

### Error Response Format

All errors follow this format:

```json
{
  "error": "HTTP Status Text",
  "message": "Detailed error description",
  "code": 400
}
```

### Common Errors

| Error Message | Status Code | Cause |
|---------------|-------------|-------|
| "Invalid request body" | 400 | Malformed JSON or invalid structure |
| "country is required" | 400 | Missing country field |
| "zipcode or postal_code is required" | 400 | Both zipcode fields are empty |
| "at least one item is required" | 400 | Empty items array |
| "item X has invalid price" | 400 | Negative price value |
| "item X has invalid quantity" | 400 | Zero or negative quantity |

---

## Rate Limiting

Currently, the API does not implement rate limiting. Consider implementing rate limiting for production deployments.

Recommended implementation:
- Use middleware to track requests per IP
- Implement token bucket or sliding window algorithm
- Return 429 Too Many Requests when limit exceeded

---

## Performance Considerations

### Response Times

Typical response times:
- Health check: < 5ms
- Tax calculation (single item): < 10ms
- Tax calculation (10 items): < 15ms

### Scalability

The API is stateless and can be horizontally scaled:
- Deploy multiple instances behind a load balancer
- No database dependencies
- No session state

### Caching

Consider implementing caching for:
- Tax rates by country (if using fixed rates)
- Frequently requested calculations

---

## Security Recommendations

For production deployments, consider:

1. **HTTPS/TLS** - Encrypt all traffic
2. **CORS** - Configure allowed origins
3. **Input Sanitization** - Additional validation
4. **API Authentication** - Implement API keys or OAuth
5. **Rate Limiting** - Prevent abuse
6. **Request Size Limits** - Prevent large payload attacks
7. **Logging** - Monitor for suspicious activity

---

## Monitoring and Observability

### Health Checks

Use the `/api/v1/health` endpoint for:
- Kubernetes liveness/readiness probes
- Load balancer health checks
- Monitoring system checks

### Metrics to Monitor

- Request rate
- Response times (p50, p95, p99)
- Error rates
- Active connections

### Logging

Implement structured logging:
```go
log.Printf("[INFO] Calculated tax for country=%s, items=%d, total=%.2f", 
    country, itemCount, grandTotal)
```

---

## Future Enhancements

Potential improvements:
1. Integration with real tax APIs (Avalara, TaxJar)
2. Database persistence for audit trails
3. Support for promotional codes and discounts
4. Multi-currency support
5. Tax exemption handling
6. Historical tax rate tracking
7. Webhooks for tax calculation events
8. GraphQL API support

---

## Support and Contact

For technical support or questions:
- GitHub Issues: https://github.com/vijayraghavareddy/tax-calculation/issues
- Email: support@example.com

---

## Changelog

### Version 1.0.0 (Current)
- Initial release
- Support for multiple countries
- RESTful API endpoints
- Comprehensive unit tests
- Documentation and Postman collection
