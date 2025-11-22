"""Data models for tax calculation API."""
from typing import List, Optional
from pydantic import BaseModel, Field, ConfigDict


class Address(BaseModel):
    """Address represents the customer's address for tax calculation."""
    
    model_config = ConfigDict(populate_by_name=True)
    
    street: str
    city: str
    state: str
    country: str
    zipcode: str
    postal_code: Optional[str] = Field(None, alias="postal_code")


class Item(BaseModel):
    """Item represents a product or service to be taxed."""
    
    id: str
    name: str
    description: Optional[str] = None
    price: float
    quantity: int


class TaxRequest(BaseModel):
    """TaxRequest represents the incoming request for tax calculation."""
    
    address: Address
    items: List[Item]


class ItemTaxDetail(BaseModel):
    """ItemTaxDetail represents tax details for a single item."""
    
    item_id: str
    item_name: str
    price: float
    quantity: int
    subtotal: float
    tax_rate: float
    tax_amount: float
    total_amount: float


class TaxResponse(BaseModel):
    """TaxResponse represents the response with calculated taxes."""
    
    address: Address
    items: List[ItemTaxDetail]
    subtotal: float
    total_tax: float
    grand_total: float
    tax_jurisdiction: str


class ErrorResponse(BaseModel):
    """ErrorResponse represents an error response."""
    
    error: str
    message: str
    code: int
