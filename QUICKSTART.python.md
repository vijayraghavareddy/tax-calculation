# Tax Calculation API - Quick Start Guide (Python)

Get the Python Tax Calculation API up and running in minutes!

## Prerequisites

- Python 3.11+
- pip

## Quick Start (5 minutes)

### 1. Clone and Setup

```bash
# Clone repository
git clone https://github.com/vijayraghavareddy/tax-calculation.git
cd tax-calculation

# Create virtual environment
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt
```

### 2. Run the Server

```bash
python main.py
```

The server will start at http://localhost:8080

### 3. Test the API

Open your browser and visit:
- **Web UI**: http://localhost:8080
- **API Docs**: http://localhost:8080/docs
- **Health Check**: http://localhost:8080/api/v1/health

Or test with curl:
```bash
curl -X POST http://localhost:8080/api/v1/calculate-tax \
  -H "Content-Type: application/json" \
  -d '{
    "address": {"state": "CA", "country": "US", "zipcode": "90001"},
    "items": [{"id": "1", "name": "Widget", "price": 100, "quantity": 1}]
  }'
```

## Using the Web Interface

1. Go to http://localhost:8080
2. Fill in:
   - State (e.g., "CA")
   - ZIP Code (e.g., "90001")
3. Add items with:
   - Name
   - Price
   - Quantity
4. Click **Calculate Tax**
5. See instant results!

## Testing

```bash
# Run all tests
pytest

# Run with coverage
pytest --cov=.

# Verbose output
pytest -v
```

## Using Make Commands

```bash
# Install dependencies
make -f Makefile.python install

# Run server
make -f Makefile.python run

# Run tests
make -f Makefile.python test

# Clean up
make -f Makefile.python clean
```

## Docker Quick Start

```bash
# Build
make -f Makefile.python docker-build

# Run
make -f Makefile.python docker-run
```

Visit http://localhost:8080

## Example API Call

### Request
```json
POST /api/v1/calculate-tax

{
  "address": {
    "state": "NY",
    "country": "US",
    "zipcode": "10001"
  },
  "items": [
    {
      "id": "item1",
      "name": "Laptop",
      "price": 1000,
      "quantity": 1
    }
  ]
}
```

### Response
```json
{
  "subtotal": 1000.00,
  "total_tax": 85.20,
  "grand_total": 1085.20,
  "tax_jurisdiction": "NY, USA",
  ...
}
```

## Key Features

✅ All 50 US states supported  
✅ Real-time tax calculation  
✅ Beautiful web UI  
✅ Auto-generated API docs  
✅ Full test coverage  
✅ Docker ready  

## Common Issues

**Module not found?**
```bash
pip install -r requirements.txt
```

**Port already in use?**
```bash
PORT=3000 python main.py
```

**Tests failing?**
```bash
pytest -v  # See detailed error messages
```

## Next Steps

- Read the full [README.python.md](README.python.md)
- Try the interactive API docs at http://localhost:8080/docs
- Run the test suite: `pytest -v`
- Customize tax rates in `tax_service.py`

## Support

Questions? Open an issue on GitHub!

Happy coding! 🚀
