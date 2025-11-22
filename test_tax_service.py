"""Test suite for tax calculation service."""
import pytest
from models import TaxRequest, Address, Item
from tax_service import TaxService


@pytest.fixture
def tax_service():
    """Create a TaxService instance for testing."""
    return TaxService()


def test_calculate_tax_success(tax_service):
    """Test successful tax calculation."""
    req = TaxRequest(
        address=Address(
            street="123 Main St",
            city="New York",
            state="NY",
            country="US",
            zipcode="10001"
        ),
        items=[
            Item(
                id="item1",
                name="Product A",
                price=100.00,
                quantity=2
            ),
            Item(
                id="item2",
                name="Product B",
                price=50.00,
                quantity=1
            )
        ]
    )
    
    resp = tax_service.calculate_tax(req)
    
    assert resp is not None
    
    # Check subtotal
    expected_subtotal = 250.00
    assert resp.subtotal == expected_subtotal
    
    # Check that tax was calculated
    assert resp.total_tax > 0
    
    # Check grand total
    assert resp.grand_total == resp.subtotal + resp.total_tax
    
    # Check items count
    assert len(resp.items) == 2
    
    # Check jurisdiction
    assert resp.tax_jurisdiction != ""


def test_calculate_tax_missing_state(tax_service):
    """Test error when state is missing."""
    req = TaxRequest(
        address=Address(
            street="",
            city="",
            state="",
            country="US",
            zipcode="10001"
        ),
        items=[
            Item(
                id="item1",
                name="Product A",
                price=100.00,
                quantity=1
            )
        ]
    )
    
    with pytest.raises(ValueError, match="state is required"):
        tax_service.calculate_tax(req)


def test_calculate_tax_missing_zipcode(tax_service):
    """Test error when zipcode is missing."""
    req = TaxRequest(
        address=Address(
            street="",
            city="",
            state="NY",
            country="US",
            zipcode=""
        ),
        items=[
            Item(
                id="item1",
                name="Product A",
                price=100.00,
                quantity=1
            )
        ]
    )
    
    with pytest.raises(ValueError, match="zipcode is required"):
        tax_service.calculate_tax(req)


def test_calculate_tax_no_items(tax_service):
    """Test error when no items are provided."""
    req = TaxRequest(
        address=Address(
            street="",
            city="",
            state="NY",
            country="US",
            zipcode="10001"
        ),
        items=[]
    )
    
    with pytest.raises(ValueError, match="at least one item is required"):
        tax_service.calculate_tax(req)


def test_calculate_tax_negative_price(tax_service):
    """Test error when item has negative price."""
    req = TaxRequest(
        address=Address(
            street="",
            city="",
            state="NY",
            country="US",
            zipcode="10001"
        ),
        items=[
            Item(
                id="item1",
                name="Product A",
                price=-100.00,
                quantity=1
            )
        ]
    )
    
    with pytest.raises(ValueError, match="item 0 has invalid price"):
        tax_service.calculate_tax(req)


def test_calculate_tax_invalid_quantity(tax_service):
    """Test error when item has invalid quantity."""
    req = TaxRequest(
        address=Address(
            street="",
            city="",
            state="NY",
            country="US",
            zipcode="10001"
        ),
        items=[
            Item(
                id="item1",
                name="Product A",
                price=100.00,
                quantity=0
            )
        ]
    )
    
    with pytest.raises(ValueError, match="item 0 has invalid quantity"):
        tax_service.calculate_tax(req)


def test_calculate_tax_multiple_items(tax_service):
    """Test tax calculation with multiple items."""
    req = TaxRequest(
        address=Address(
            street="456 Oak Ave",
            city="Los Angeles",
            state="CA",
            country="US",
            zipcode="90001"
        ),
        items=[
            Item(id="item1", name="Product A", price=50.00, quantity=3),
            Item(id="item2", name="Product B", price=25.00, quantity=2),
            Item(id="item3", name="Product C", price=10.00, quantity=5)
        ]
    )
    
    resp = tax_service.calculate_tax(req)
    
    assert len(resp.items) == 3
    assert resp.subtotal == 250.00
    assert resp.total_tax > 0
    assert resp.grand_total == resp.subtotal + resp.total_tax


def test_calculate_tax_zero_tax_state(tax_service):
    """Test tax calculation for states with no sales tax."""
    req = TaxRequest(
        address=Address(
            street="789 Pine St",
            city="Portland",
            state="OR",
            country="US",
            zipcode="97201"
        ),
        items=[
            Item(id="item1", name="Product A", price=100.00, quantity=1)
        ]
    )
    
    resp = tax_service.calculate_tax(req)
    
    assert resp.subtotal == 100.00
    assert resp.total_tax == 0.00
    assert resp.grand_total == 100.00


def test_tax_rate_for_california(tax_service):
    """Test correct tax rate for California."""
    address = Address(
        street="123 Main St",
        city="Los Angeles",
        state="CA",
        country="US",
        zipcode="90001"
    )
    
    rate = tax_service._get_tax_rate_for_location(address)
    assert rate == 0.0850


def test_tax_rate_for_unknown_state(tax_service):
    """Test default tax rate for unknown state."""
    address = Address(
        street="123 Main St",
        city="Unknown City",
        state="XX",
        country="US",
        zipcode="00000"
    )
    
    rate = tax_service._get_tax_rate_for_location(address)
    assert rate == 0.0700  # Default rate


def test_round_to_two_decimals(tax_service):
    """Test rounding to two decimal places."""
    assert tax_service._round_to_two_decimals(10.123) == 10.12
    assert tax_service._round_to_two_decimals(10.126) == 10.13
    assert tax_service._round_to_two_decimals(10.125) == 10.12
    assert tax_service._round_to_two_decimals(10.0) == 10.0
