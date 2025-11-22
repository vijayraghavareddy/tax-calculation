"""API handlers for tax calculation endpoints."""
from fastapi import APIRouter, HTTPException, status
from fastapi.responses import JSONResponse
from models import TaxRequest, TaxResponse, ErrorResponse
from tax_service import TaxService


router = APIRouter(prefix="/api/v1")
tax_service = TaxService()


@router.post("/calculate-tax", response_model=TaxResponse, status_code=status.HTTP_200_OK)
async def calculate_tax(request: TaxRequest) -> TaxResponse:
    """Calculate tax for the given request.
    
    Args:
        request: The tax calculation request
        
    Returns:
        The tax calculation response
        
    Raises:
        HTTPException: If the request is invalid
    """
    try:
        # Normalize postal_code to zipcode if provided
        if not request.address.zipcode and request.address.postal_code:
            request.address.zipcode = request.address.postal_code
        
        response = tax_service.calculate_tax(request)
        return response
    except ValueError as e:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=str(e)
        )


@router.get("/health")
async def health_check() -> dict:
    """Health check endpoint.
    
    Returns:
        Health status information
    """
    return {
        "status": "healthy",
        "service": "tax-calculation-api",
        "version": "1.0.0"
    }
