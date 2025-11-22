# Tax Calculation API - Bruno Collection

This is a Bruno collection for testing the Tax Calculation API endpoints. This API calculates taxes based on customer address and items purchased.

## About Bruno

Bruno is a fast and git-friendly open-source API client. Unlike Postman, Bruno stores collections directly in your filesystem using a plain text markup language, making it easy to version control with Git.

## Installation

1. Install Bruno from [https://www.usebruno.com/](https://www.usebruno.com/)
2. Open Bruno
3. Click "Open Collection"
4. Navigate to this `bruno-collection` folder

## Environment Variables

The collection uses the following environment variable:
- `base_url`: The base URL of the API (default: `http://localhost:8080`)

You can modify the environment variables in `environments/Local.bru` or create new environment files.

## Collection Structure

The collection includes 10 requests organized as follows:

### Health Check
- **Health Check**: Verify the API service is running

### Successful Tax Calculations
- **Calculate Tax - US Purchase**: Single item purchase in New York
- **Calculate Tax - Multiple Items**: Multiple items purchase in Toronto, Canada
- **Calculate Tax - UK with Postal Code**: Purchase in London using postal_code
- **Calculate Tax - High Value Purchase**: High-value purchase in California
- **Calculate Tax - India GST**: GST calculation for Indian address

### Error Handling Tests
- **Calculate Tax - Error: Missing Country**: Test validation when country is missing
- **Calculate Tax - Error: Missing Zipcode**: Test validation when zipcode is missing
- **Calculate Tax - Error: No Items**: Test validation when items array is empty
- **Calculate Tax - Error: Negative Price**: Test validation when item has negative price

## Running the Collection

1. Make sure your API server is running on `http://localhost:8080` (or update the environment variable)
2. Open the collection in Bruno
3. Run individual requests or use Bruno's collection runner to execute all requests

## Converting from Postman

This collection was converted from the Postman collection. Key differences:
- Bruno uses `.bru` files instead of JSON
- Tests use slightly different syntax (Bruno's test API vs Postman's `pm` API)
- Each request is a separate file, making it easier to track changes in Git
- Environment variables are stored in plain text files

## More Information

For more information about the API, see:
- `API_DOCUMENTATION.md` - Detailed API documentation
- `README.md` - Project overview
- `QUICKSTART.md` - Getting started guide
