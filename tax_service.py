"""Tax calculation service."""
from typing import List
from models import TaxRequest, TaxResponse, ItemTaxDetail, Address


class TaxService:
    """TaxService handles tax calculation logic."""
    
    def calculate_tax(self, req: TaxRequest) -> TaxResponse:
        """Calculate tax for the given request.
        
        Args:
            req: The tax calculation request
            
        Returns:
            The tax calculation response
            
        Raises:
            ValueError: If the request is invalid
        """
        self._validate_request(req)
        
        # Get tax rate based on location
        tax_rate = self._get_tax_rate_for_location(req.address)
        jurisdiction = self._get_tax_jurisdiction(req.address)
        
        item_details: List[ItemTaxDetail] = []
        subtotal = 0.0
        total_tax = 0.0
        
        # Calculate tax for each item
        for item in req.items:
            item_subtotal = item.price * item.quantity
            item_tax = item_subtotal * tax_rate
            item_total = item_subtotal + item_tax
            
            detail = ItemTaxDetail(
                item_id=item.id,
                item_name=item.name,
                price=item.price,
                quantity=item.quantity,
                subtotal=self._round_to_two_decimals(item_subtotal),
                tax_rate=self._round_to_two_decimals(tax_rate * 100),  # Convert to percentage
                tax_amount=self._round_to_two_decimals(item_tax),
                total_amount=self._round_to_two_decimals(item_total)
            )
            
            item_details.append(detail)
            subtotal += item_subtotal
            total_tax += item_tax
        
        response = TaxResponse(
            address=req.address,
            items=item_details,
            subtotal=self._round_to_two_decimals(subtotal),
            total_tax=self._round_to_two_decimals(total_tax),
            grand_total=self._round_to_two_decimals(subtotal + total_tax),
            tax_jurisdiction=jurisdiction
        )
        
        return response
    
    def _validate_request(self, req: TaxRequest) -> None:
        """Validate the tax calculation request.
        
        Args:
            req: The request to validate
            
        Raises:
            ValueError: If the request is invalid
        """
        if not req.address.state:
            raise ValueError("state is required")
        if not req.address.zipcode and not req.address.postal_code:
            raise ValueError("zipcode is required")
        if not req.items:
            raise ValueError("at least one item is required")
        
        for i, item in enumerate(req.items):
            if item.price < 0:
                raise ValueError(f"item {i} has invalid price")
            if item.quantity <= 0:
                raise ValueError(f"item {i} has invalid quantity")
    
    def _get_tax_rate_for_location(self, address: Address) -> float:
        """Get tax rate based on the US state.
        
        Rates are approximate and based on combined state and average local rates.
        
        Args:
            address: The address to get the tax rate for
            
        Returns:
            The tax rate as a decimal (e.g., 0.0852 for 8.52%)
        """
        state = address.state.upper()
        
        # US state sales tax rates (approximate combined rates)
        tax_rates = {
            "AL": 0.0913,  # Alabama
            "AK": 0.0176,  # Alaska
            "AZ": 0.0831,  # Arizona
            "AR": 0.0947,  # Arkansas
            "CA": 0.0850,  # California
            "CO": 0.0763,  # Colorado
            "CT": 0.0635,  # Connecticut
            "DE": 0.0000,  # Delaware - No sales tax
            "FL": 0.0705,  # Florida
            "GA": 0.0733,  # Georgia
            "HI": 0.0444,  # Hawaii
            "ID": 0.0602,  # Idaho
            "IL": 0.0868,  # Illinois
            "IN": 0.0700,  # Indiana
            "IA": 0.0694,  # Iowa
            "KS": 0.0865,  # Kansas
            "KY": 0.0600,  # Kentucky
            "LA": 0.0952,  # Louisiana
            "ME": 0.0550,  # Maine
            "MD": 0.0600,  # Maryland
            "MA": 0.0625,  # Massachusetts
            "MI": 0.0600,  # Michigan
            "MN": 0.0744,  # Minnesota
            "MS": 0.0707,  # Mississippi
            "MO": 0.0824,  # Missouri
            "MT": 0.0000,  # Montana - No sales tax
            "NE": 0.0694,  # Nebraska
            "NV": 0.0823,  # Nevada
            "NH": 0.0000,  # New Hampshire - No sales tax
            "NJ": 0.0663,  # New Jersey
            "NM": 0.0779,  # New Mexico
            "NY": 0.0852,  # New York
            "NC": 0.0698,  # North Carolina
            "ND": 0.0696,  # North Dakota
            "OH": 0.0723,  # Ohio
            "OK": 0.0897,  # Oklahoma
            "OR": 0.0000,  # Oregon - No sales tax
            "PA": 0.0634,  # Pennsylvania
            "RI": 0.0700,  # Rhode Island
            "SC": 0.0744,  # South Carolina
            "SD": 0.0645,  # South Dakota
            "TN": 0.0955,  # Tennessee
            "TX": 0.0820,  # Texas
            "UT": 0.0719,  # Utah
            "VT": 0.0624,  # Vermont
            "VA": 0.0575,  # Virginia
            "WA": 0.0920,  # Washington
            "WV": 0.0650,  # West Virginia
            "WI": 0.0543,  # Wisconsin
            "WY": 0.0536,  # Wyoming
        }
        
        # Default rate if state not recognized
        return tax_rates.get(state, 0.0700)
    
    def _get_tax_jurisdiction(self, address: Address) -> str:
        """Get the tax jurisdiction string.
        
        Args:
            address: The address to get the jurisdiction for
            
        Returns:
            The jurisdiction string
        """
        return f"{address.state}, USA"
    
    @staticmethod
    def _round_to_two_decimals(value: float) -> float:
        """Round a float to 2 decimal places.
        
        Args:
            value: The value to round
            
        Returns:
            The rounded value
        """
        return round(value, 2)
