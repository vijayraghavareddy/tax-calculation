# GitHub Copilot Instructions

## JIRA Configuration

When working with JIRA in this repository:

- **Default JIRA Board**: Always use the board called "TAX"
- **Default JIRA User**: Use "developer-1" for any user-related queries or assignments

When searching for issues, creating issues, or performing any JIRA operations, automatically scope queries to the TAX board and use developer-1 as the default user unless explicitly specified otherwise.

## Implementation Philosophy

**PRIORITY: Speed and completion over perfection**

- **Quick Implementation**: Focus on getting working code implemented fast
- **Minimal Testing**: Add only 1-2 simple tests to verify core functionality
- **Single Test Run**: Run tests once to verify basic correctness - do NOT rerun multiple times
- **No Exhaustive Testing**: Skip edge cases, extensive test coverage, and multiple test scenarios
- **Complete Fast**: Aim to complete implementation and move on quickly
- **Avoid Over-Engineering**: Simple, straightforward solutions are preferred

### Branch & Implementation Strategy

- **Always Start Fresh**: Ignore any previous implementations of the feature
- **New Branch**: Create a new branch with unique ID at the end (e.g., `TAX-123-feature-name-a1b2c3`)
- **Clean Implementation**: Start implementation from scratch without referencing old code
- **No Previous Mentions**: Do not mention or document previous implementation attempts
- **Forward Only**: Focus on the current implementation, not what was done before

## Python Development Guidelines

This project uses Python with FastAPI for the Tax Calculation API.

### Code Style & Formatting
- **Formatting**: Use `black` for code formatting (line length: 88)
- **Import Sorting**: Use `isort` for consistent import ordering
- **Linting**: Follow `pylint` and `flake8` recommendations
- **Type Hints**: Always use type hints for function parameters and return values
- **Docstrings**: Use Google-style docstrings for all functions, classes, and modules

### Project Structure
- **Root Level Files**: 
  - `main.py` - FastAPI application entry point
  - `models.py` - Pydantic data models
  - `handlers.py` - API route handlers
  - `tax_service.py` - Business logic
- **Tests**: Test files prefixed with `test_` (e.g., `test_tax_service.py`)
- **Static Files**: Web UI files in `static/` directory

### Error Handling
- **Validation**: Use Pydantic models for automatic request validation
- **Exceptions**: Raise `ValueError` for business logic errors
- **HTTP Exceptions**: Use `HTTPException` from FastAPI for API errors
- **Error Messages**: Provide clear, descriptive error messages

### Testing
- **Framework**: Use `pytest` for all tests
- **Fixtures**: Use pytest fixtures for reusable test setup
- **Coverage**: Maintain high test coverage (aim for >90%)
- **Test Organization**: 
  - Unit tests for service layer (`test_tax_service.py`)
  - Integration tests for API endpoints (`test_handlers.py`)
- **Assertions**: Use descriptive assertion messages

### API Design
- **RESTful**: Follow RESTful conventions
- **Status Codes**: Use appropriate HTTP status codes (200, 400, 404, 422, 500)
- **Response Models**: Define Pydantic response models for all endpoints
- **Documentation**: FastAPI auto-generates docs at `/docs` and `/redoc`
- **CORS**: CORS middleware configured for cross-origin requests

### Naming Conventions
- **Variables/Functions**: Use `snake_case` (e.g., `calculate_tax`, `tax_rate`)
- **Classes**: Use `PascalCase` (e.g., `TaxService`, `TaxRequest`)
- **Constants**: Use `UPPER_SNAKE_CASE` (e.g., `DEFAULT_TAX_RATE`)
- **Private Methods**: Prefix with single underscore (e.g., `_validate_request`)

### Dependencies
- **Core**: FastAPI, uvicorn, pydantic
- **Testing**: pytest, pytest-cov, httpx
- **Management**: Keep `requirements.txt` updated

### Best Practices
- **Async**: Use async/await for I/O operations when beneficial
- **Validation**: Let Pydantic handle data validation
- **DRY**: Don't repeat yourself - extract common logic
- **Single Responsibility**: Each function/class should have one purpose
- **Documentation**: Keep README.python.md and docstrings up to date

### Running the Application
```bash
# Development
python main.py

# Production (with uvicorn)
uvicorn main:app --host 0.0.0.0 --port 8080 --workers 4

# Testing
pytest -v
pytest --cov=. --cov-report=html
```

---

## Legacy Go Code (Archived)

The original Go implementation is preserved in:
- `main.go`
- `handlers/` directory
- `models/` directory  
- `services/` directory
- `go.mod`

These files are kept for reference but are no longer actively maintained.
