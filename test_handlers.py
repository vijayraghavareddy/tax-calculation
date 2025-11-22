"""Test suite for API handlers."""
import pytest
from fastapi.testclient import TestClient
from main import app
from models import TaxRequest, Address, Item


client = TestClient(app)


def test_calculate_tax_valid_request():
    """Test valid tax calculation request."""
    req_body = {
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
                "price": 100.00,
                "quantity": 2
            }
        ]
    }
    
    response = client.post("/api/v1/calculate-tax", json=req_body)
    
    assert response.status_code == 200
    data = response.json()
    
    assert data["subtotal"] == 200.00
    assert data["total_tax"] > 0
    assert data["grand_total"] == data["subtotal"] + data["total_tax"]


def test_calculate_tax_invalid_json():
    """Test invalid JSON request."""
    response = client.post(
        "/api/v1/calculate-tax",
        data="invalid json",
        headers={"Content-Type": "application/json"}
    )
    
    assert response.status_code == 422  # Unprocessable Entity


def test_calculate_tax_missing_state():
    """Test error when state is missing."""
    req_body = {
        "address": {
            "street": "",
            "city": "",
            "state": "",
            "country": "US",
            "zipcode": "10001"
        },
        "items": [
            {
                "id": "item1",
                "name": "Product A",
                "price": 100.00,
                "quantity": 1
            }
        ]
    }
    
    response = client.post("/api/v1/calculate-tax", json=req_body)
    
    assert response.status_code == 400


def test_calculate_tax_with_postal_code():
    """Test tax calculation with postal_code field."""
    req_body = {
        "address": {
            "street": "123 Main St",
            "city": "Boston",
            "state": "MA",
            "country": "US",
            "zipcode": "",
            "postal_code": "02101"
        },
        "items": [
            {
                "id": "item1",
                "name": "Product A",
                "price": 100.00,
                "quantity": 1
            }
        ]
    }
    
    response = client.post("/api/v1/calculate-tax", json=req_body)
    
    assert response.status_code == 200


def test_health_check():
    """Test health check endpoint."""
    response = client.get("/api/v1/health")
    
    assert response.status_code == 200
    data = response.json()
    
    assert data["status"] == "healthy"
    assert data["service"] == "tax-calculation-api"
    assert data["version"] == "1.0.0"


def test_calculate_tax_negative_price():
    """Test error with negative price."""
    req_body = {
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
                "price": -100.00,
                "quantity": 1
            }
        ]
    }
    
    response = client.post("/api/v1/calculate-tax", json=req_body)
    
    assert response.status_code == 400


def test_calculate_tax_multiple_items():
    """Test tax calculation with multiple items."""
    req_body = {
        "address": {
            "street": "456 Oak Ave",
            "city": "Los Angeles",
            "state": "CA",
            "country": "US",
            "zipcode": "90001"
        },
        "items": [
            {"id": "item1", "name": "Product A", "price": 50.00, "quantity": 3},
            {"id": "item2", "name": "Product B", "price": 25.00, "quantity": 2},
            {"id": "item3", "name": "Product C", "price": 10.00, "quantity": 5}
        ]
    }
    
    response = client.post("/api/v1/calculate-tax", json=req_body)
    
    assert response.status_code == 200
    data = response.json()
    
    assert len(data["items"]) == 3
    assert data["subtotal"] == 250.00


def test_serve_index():
    """Test serving the index page."""
    response = client.get("/")
    
    assert response.status_code == 200
