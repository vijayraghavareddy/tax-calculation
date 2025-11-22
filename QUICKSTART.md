# Quick Start Guide - Tax Calculation API

## 🚀 Getting Started in 3 Steps

### 1. Install & Build
```bash
cd /home/vijayreddy/repo/tax-calculation
go mod download
go build -o tax-api
```

### 2. Run the API
```bash
./tax-api
# Or in development mode:
# go run main.go
```

Server starts on: `http://localhost:8080`

### 3. Use the Application

**Option A: Web UI (Recommended)**
1. Open your browser: `http://localhost:8080`
2. Fill in address and items
3. Click "Calculate Tax"
4. View results instantly! ✨

**Option B: Test with cURL**
```bash
curl -X POST http://localhost:8080/api/v1/calculate-tax \
  -H "Content-Type: application/json" \
  -d '{
    "address": {
      "country": "US",
      "zipcode": "10001",
      "state": "NY",
      "city": "New York"
    },
    "items": [
      {
        "id": "1",
        "name": "Product",
        "price": 100.00,
        "quantity": 1
      }
    ]
  }'
```

**Option C: Postman**
Import `postman_collection.json` and start testing!

---

## 🎨 Web UI Features

The web interface provides an intuitive way to calculate taxes:

- ✅ **Easy Form Input** - No coding required
- ✅ **Dynamic Items** - Add/remove items on the fly
- ✅ **Real-time Validation** - Instant feedback on errors
- ✅ **Detailed Breakdown** - See tax per item
- ✅ **Mobile Responsive** - Works on all devices
- ✅ **Beautiful Design** - Modern, gradient UI

### Screenshots Features:
- Country selector with flags 🇺🇸 🇨🇦 🇬🇧 🇩🇪 🇫🇷 🇮🇳 🇦🇺 🇯🇵
- Item cards with price and quantity
- Results summary with grand total
- Per-item tax breakdown

---

## 📋 Useful Commands

```bash
# Run tests
go test ./... -v

# Run with coverage
go test ./... -cover

# Format code
go fmt ./...

# Build for production
go build -o tax-api

# Run on different port
PORT=3000 go run main.go

# Using Makefile
make build      # Build application
make run        # Run application
make test       # Run tests
make coverage   # Generate coverage report
make clean      # Clean build artifacts
```

---

## 🧪 Testing with Postman

1. Open Postman
2. Click "Import"
3. Select `postman_collection.json`
4. Collection includes:
   - ✅ Health check
   - ✅ Valid requests (US, CA, UK, IN)
   - ✅ Multiple items
   - ✅ Error scenarios
   - ✅ Automated tests

---

## 📦 Request Format

**Minimum Required:**
```json
{
  "address": {
    "country": "US",
    "zipcode": "10001"
  },
  "items": [
    {
      "id": "item1",
      "name": "Product",
      "price": 100.00,
      "quantity": 1
    }
  ]
}
```

**Full Example:**
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
      "description": "Optional description",
      "price": 100.00,
      "quantity": 2
    }
  ]
}
```

---

## 🌍 Supported Countries

| Code | Country | Tax Type | Rate |
|------|---------|----------|------|
| US | United States | Sales Tax | 5-12% |
| CA | Canada | GST/HST | 5-15% |
| UK/GB | United Kingdom | VAT | ~20% |
| DE | Germany | VAT | ~19% |
| FR | France | VAT | ~20% |
| IN | India | GST | 5-28% |
| AU | Australia | GST | ~10% |
| JP | Japan | Consumption Tax | ~10% |

---

## 📊 Response Format

```json
{
  "address": { /* original address */ },
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
    }
  ],
  "subtotal": 200.00,
  "total_tax": 17.00,
  "grand_total": 217.00,
  "tax_jurisdiction": "NY, US"
}
```

---

## ❌ Common Errors

| Error | Cause | Fix |
|-------|-------|-----|
| "country is required" | Missing country field | Add `"country": "US"` |
| "zipcode or postal_code is required" | Missing both | Add either field |
| "at least one item is required" | Empty items array | Add at least one item |
| "invalid price" | Negative price | Use positive numbers |
| "invalid quantity" | Zero/negative quantity | Use positive integers |

---

## 🔍 Health Check

```bash
curl http://localhost:8080/api/v1/health
```

Response:
```json
{
  "status": "healthy",
  "service": "tax-calculation-api",
  "version": "1.0.0"
}
```

---

## 🐳 Docker (Optional)

```bash
# Build image
docker build -t tax-api .

# Run container
docker run -p 8080:8080 tax-api

# Or using Makefile
make docker-build
make docker-run
```

---

## 📁 Project Structure

```
tax-calculation/
├── main.go                    # Entry point
├── handlers/                  # HTTP handlers + tests
├── services/                  # Business logic + tests
├── models/                    # Data structures
├── README.md                  # Full documentation
├── API_DOCUMENTATION.md       # Detailed API docs
├── postman_collection.json    # Postman tests
├── Makefile                   # Build automation
└── Dockerfile                 # Container setup
```

---

## 💡 Tips

- **Development**: Use `go run main.go` for quick iterations
- **Production**: Build with `go build -o tax-api` for better performance
- **Testing**: Always run `go test ./...` before committing
- **Coverage**: Use `make coverage` to generate HTML coverage report
- **Postman**: Import collection for comprehensive API testing
- **Port**: Set `PORT` environment variable to change default port

---

## 🆘 Troubleshooting

**Port already in use:**
```bash
PORT=3000 go run main.go
```

**Module errors:**
```bash
go mod tidy
go mod download
```

**Tests failing:**
```bash
go clean -testcache
go test ./... -v
```

---

## 📚 Documentation

- **README.md** - Complete project documentation
- **API_DOCUMENTATION.md** - Detailed API reference
- **postman_collection.json** - Interactive API testing

---

## ✨ Features

✅ Multi-country support  
✅ Realistic tax calculations  
✅ Address validation  
✅ Comprehensive tests (20+ test cases)  
✅ Error handling  
✅ Docker support  
✅ Postman collection  
✅ Full documentation  
✅ **Beautiful Web UI**  
✅ **Mobile responsive design**  

---

## 🌐 Access Points

Once the server is running:

- **Web Interface:** http://localhost:8080
- **API Endpoint:** http://localhost:8080/api/v1/calculate-tax
- **Health Check:** http://localhost:8080/api/v1/health

---

Happy coding! 🎉
