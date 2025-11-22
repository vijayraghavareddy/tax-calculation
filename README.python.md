# Tax Calculation API (Python)

A RESTful API service built with Python FastAPI that calculates taxes based on address information and item details. The service provides location-based tax rate calculation with support for multiple US states.

## Features

- 🎨 **Modern Web UI** - Beautiful, responsive interface for easy tax calculation
- 🌍 US state-based tax calculation support
- 📦 Item-based tax computation
- 🔍 Address validation (state, zipcode/postal code)
- ✅ Comprehensive unit tests with pytest
- 📚 Automatic API documentation with Swagger/OpenAPI
- 🚀 Simple deployment with FastAPI and uvicorn
- 📱 Mobile-responsive design

## Supported Regions

The API provides realistic tax rates for all 50 US states, including:
- States with no sales tax (OR, DE, MT, NH, AK)
- Combined state and local average rates
- Accurate rates as of 2024

## Prerequisites

- Python 3.11 or higher
- pip (Python package manager)
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/vijayraghavareddy/tax-calculation.git
cd tax-calculation
```

2. Create a virtual environment (recommended):
```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

3. Install dependencies:
```bash
pip install -r requirements.txt
```

## Running the API

### Development Mode

```bash
python main.py
```

Or using uvicorn directly:
```bash
uvicorn main:app --reload --port 8080
```

The server will start on `http://localhost:8080` by default.

**Access Points:**
- **Web UI:** http://localhost:8080
- **API Endpoint:** http://localhost:8080/api/v1/calculate-tax
- **Health Check:** http://localhost:8080/api/v1/health
- **API Documentation (Swagger):** http://localhost:8080/docs
- **Alternative API Docs (ReDoc):** http://localhost:8080/redoc

To use a different port, set the `PORT` environment variable:
```bash
PORT=3000 python main.py
```

## Using the Web UI

1. Open your browser and navigate to `http://localhost:8080`
2. Fill in the address information (state and zip code are required)
3. Add one or more items with name, price, and quantity
4. Click "Calculate Tax" to see the results
5. View detailed breakdown of tax calculations for each item

The UI provides:
- ✨ Real-time validation
- 🎯 Easy item management (add/remove items)
- 📊 Detailed tax breakdown
- 🌐 Support for all US states
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
      "tax_rate": 8.52,
      "tax_amount": 17.04,
      "total_amount": 217.04
    },
    {
      "item_id": "item2",
      "item_name": "Product B",
      "price": 50.00,
      "quantity": 1,
      "subtotal": 50.00,
      "tax_rate": 8.52,
      "tax_amount": 4.26,
      "total_amount": 54.26
    }
  ],
  "subtotal": 250.00,
  "total_tax": 21.30,
  "grand_total": 271.30,
  "tax_jurisdiction": "NY, USA"
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
| state | string | **Yes** | US state code (e.g., NY, CA) |
| country | string | Yes | Country code (US) |
| zipcode | string | **Yes*** | ZIP code |
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
  "detail": "state is required"
}
```

### Common Error Codes

- `400 Bad Request` - Invalid request parameters
- `404 Not Found` - Endpoint not found
- `422 Unprocessable Entity` - Validation error
- `500 Internal Server Error` - Server error

## Testing

### Run All Tests

```bash
pytest
```

### Run Tests with Coverage

```bash
pytest --cov=. --cov-report=html
```

### Run Tests with Verbose Output

```bash
pytest -v
```

### Run Specific Test Files

```bash
# Test service
pytest test_tax_service.py -v

# Test handlers
pytest test_handlers.py -v
```

## Project Structure

```
tax-calculation/
├── main.py                 # FastAPI application entry point
├── handlers.py             # API route handlers
├── models.py               # Pydantic data models
├── tax_service.py          # Tax calculation business logic
├── test_tax_service.py     # Service unit tests
├── test_handlers.py        # Handler/API tests
├── requirements.txt        # Python dependencies
├── Dockerfile.python       # Docker configuration
├── Makefile.python         # Build automation
├── static/                 # Web UI files
│   ├── index.html         # Main HTML page
│   ├── styles.css         # CSS styling
│   └── app.js             # JavaScript logic
├── README.md               # This file
└── .gitignore             # Git ignore rules
```

## Testing with cURL

### Calculate Tax
```bash
curl -X POST http://localhost:8080/api/v1/calculate-tax \
  -H "Content-Type: application/json" \
  -d '{
    "address": {
      "street": "123 Main St",
      "city": "San Francisco",
      "state": "CA",
      "country": "US",
      "zipcode": "94102"
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

### Adding Support for New States

Edit `tax_service.py` and add the state to the `_get_tax_rate_for_location` method in the `tax_rates` dictionary:

```python
"XX": 0.0850,  # State Name
```

### Modifying Tax Calculation Logic

The tax calculation logic is in `tax_service.py`. The `calculate_tax` method contains the core business logic.

## Docker Support

### Build Docker Image

Using the Makefile:
```bash
make -f Makefile.python docker-build
```

Or directly:
```bash
docker build -f Dockerfile.python -t tax-calculation-api:python .
```

### Run Docker Container

Using the Makefile:
```bash
make -f Makefile.python docker-run
```

Or directly:
```bash
docker run -p 8080:8080 tax-calculation-api:python
```

## Using the Makefile

The project includes a Python-specific Makefile with common commands:

```bash
# Install dependencies
make -f Makefile.python install

# Run the application
make -f Makefile.python run

# Run tests
make -f Makefile.python test

# Clean generated files
make -f Makefile.python clean

# Build Docker image
make -f Makefile.python docker-build

# Run Docker container
make -f Makefile.python docker-run
```

## Migration from Go

This project was converted from Go to Python. Key changes:

1. **Web Framework**: Gorilla Mux → FastAPI
2. **Models**: Go structs → Pydantic models
3. **Testing**: Go testing → pytest
4. **Dependencies**: go.mod → requirements.txt
5. **Type System**: Go native types → Python type hints with Pydantic validation

The old Go files are preserved in their original directories and can be referenced if needed.

## API Documentation

FastAPI automatically generates interactive API documentation:

- **Swagger UI**: http://localhost:8080/docs
  - Interactive API testing interface
  - Try out endpoints directly in the browser
  
- **ReDoc**: http://localhost:8080/redoc
  - Clean, readable API documentation
  - Better for viewing and understanding the API

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
