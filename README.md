# Tax Calculation API

A RESTful API service built with Go that calculates taxes based on address information and item details. The service provides location-based tax rate calculation with support for multiple countries.

## Features

- 🎨 **Modern Web UI** - Beautiful, responsive interface for easy tax calculation
- 🌍 Multi-country tax calculation support
- 📦 Item-based tax computation
- 🔍 Address validation (country, zipcode/postal code)
- ✅ Comprehensive unit tests
- 📚 Well-documented API endpoints
- 🚀 Simple deployment with no external dependencies
- 📱 Mobile-responsive design

## Supported Countries

The API provides realistic tax rates for the following countries:
- 🇺🇸 United States (0-12% sales tax)
- 🇨🇦 Canada (5-15% GST/HST)
- 🇬🇧 United Kingdom (20% VAT)
- 🇩🇪 Germany (19% VAT)
- 🇫🇷 France (20% VAT)
- 🇮🇳 India (5-28% GST)
- 🇦🇺 Australia (10% GST)
- 🇯🇵 Japan (10% consumption tax)
- And more...

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/vijayraghavareddy/tax-calculation.git
cd tax-calculation
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o tax-api
```

## Running the API

### Development Mode

```bash
go run main.go
```

### Production Mode

```bash
./tax-api
```

The server will start on `http://localhost:8080` by default.

**Access Points:**
- **Web UI:** http://localhost:8080
- **API Endpoint:** http://localhost:8080/api/v1/calculate-tax
- **Health Check:** http://localhost:8080/api/v1/health

To use a different port, set the `PORT` environment variable:
```bash
PORT=3000 go run main.go
```

## Using the Web UI

1. Open your browser and navigate to `http://localhost:8080`
2. Fill in the address information (country and zip code are required)
3. Add one or more items with name, price, and quantity
4. Click "Calculate Tax" to see the results
5. View detailed breakdown of tax calculations for each item

The UI provides:
- ✨ Real-time validation
- 🎯 Easy item management (add/remove items)
- 📊 Detailed tax breakdown
- 🌐 Support for 8+ countries
- 📱 Mobile-friendly responsive design

## API Endpoints

### 1. Calculate Tax

Calculate tax for items based on address.

**Endpoint:** `POST /api/v1/calculate-tax`

**Request Body:**
```json
{
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "country": "US",
    "zipcode": "10001"
  },
  "items": [
    {
      "id": "item1",
      "name": "Product A",
      "description": "A great product",
      "price": 100.00,
      "quantity": 2
    },
    {
      "id": "item2",
      "name": "Product B",
      "price": 50.00,
      "quantity": 1
    }
  ]
}
```

**Response:**
```json
{
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "country": "US",
    "zipcode": "10001"
  },
  "items": [
    {
      "item_id": "item1",
      "item_name": "Product A",
      "price": 100.00,
      "quantity": 2,
      "subtotal": 200.00,
      "tax_rate": 8.50,
      "tax_amount": 17.00,
      "total_amount": 217.00
    },
    {
      "item_id": "item2",
      "item_name": "Product B",
      "price": 50.00,
      "quantity": 1,
      "subtotal": 50.00,
      "tax_rate": 8.50,
      "tax_amount": 4.25,
      "total_amount": 54.25
    }
  ],
  "subtotal": 250.00,
  "total_tax": 21.25,
  "grand_total": 271.25,
  "tax_jurisdiction": "NY, US"
}
```

### 2. Health Check

Check if the API is running.

**Endpoint:** `GET /api/v1/health`

**Response:**
```json
{
  "status": "healthy",
  "service": "tax-calculation-api",
  "version": "1.0.0"
}
```

## Request Fields

### Address Object

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| street | string | No | Street address |
| city | string | No | City name |
| state | string | No | State/province code |
| country | string | **Yes** | Country code (US, CA, UK, etc.) |
| zipcode | string | **Yes*** | ZIP/postal code |
| postal_code | string | **Yes*** | Alternative to zipcode |

*Either `zipcode` or `postal_code` must be provided.

### Item Object

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| id | string | **Yes** | Unique item identifier |
| name | string | **Yes** | Item name |
| description | string | No | Item description |
| price | number | **Yes** | Unit price (must be >= 0) |
| quantity | integer | **Yes** | Quantity (must be > 0) |

## Error Responses

The API returns standard HTTP error codes with descriptive messages:

```json
{
  "error": "Bad Request",
  "message": "country is required",
  "code": 400
}
```

### Common Error Codes

- `400 Bad Request` - Invalid request parameters
- `404 Not Found` - Endpoint not found
- `500 Internal Server Error` - Server error

## Testing

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test ./... -cover
```

### Run Tests with Verbose Output

```bash
go test ./... -v
```

### Run Specific Package Tests

```bash
# Test handlers
go test ./handlers -v

# Test services
go test ./services -v
```

## Project Structure

```
tax-calculation/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── handlers/               # HTTP request handlers
│   ├── handlers.go
│   └── handlers_test.go
├── models/                 # Data models
│   └── models.go
├── services/               # Business logic
│   ├── tax_service.go
│   └── tax_service_test.go
├── static/                 # Web UI files
│   ├── index.html         # Main HTML page
│   ├── styles.css         # CSS styling
│   └── app.js             # JavaScript logic
├── README.md               # Documentation
├── API_DOCUMENTATION.md    # Detailed API docs
└── postman_collection.json # Postman collection for testing
```

## Testing with Postman

Import the `postman_collection.json` file into Postman to test all API endpoints with pre-configured requests.

### Quick Start with Postman:
1. Open Postman
2. Click "Import" button
3. Select `postman_collection.json`
4. Use the "Tax Calculation API" collection to test endpoints

## Example Usage with cURL

### Calculate Tax
```bash
curl -X POST http://localhost:8080/api/v1/calculate-tax \
  -H "Content-Type: application/json" \
  -d '{
    "address": {
      "street": "123 Main St",
      "city": "Toronto",
      "state": "ON",
      "country": "CA",
      "zipcode": "M5H2N2"
    },
    "items": [
      {
        "id": "item1",
        "name": "Laptop",
        "price": 999.99,
        "quantity": 1
      }
    ]
  }'
```

### Health Check
```bash
curl http://localhost:8080/api/v1/health
```

## Development

### Adding Support for New Countries

Edit `services/tax_service.go` and add the country to the `getTaxRateForLocation` function:

```go
case "XX", "COUNTRY NAME":
    // Country: tax rate
    baseRate = 0.15
```

### Modifying Tax Calculation Logic

The tax calculation logic is in `services/tax_service.go`. The `CalculateTax` method contains the core business logic.

## Docker Support (Optional)

Create a `Dockerfile`:
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o tax-api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/tax-api .
EXPOSE 8080
CMD ["./tax-api"]
```

Build and run:
```bash
docker build -t tax-calculation-api .
docker run -p 8080:8080 tax-calculation-api
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Author

Vijay Raghava Reddy

## Support

For issues, questions, or contributions, please open an issue on GitHub.
