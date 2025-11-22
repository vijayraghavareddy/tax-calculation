# GitHub Copilot Instructions

## JIRA Configuration

When working with JIRA in this repository:

- **Default JIRA Board**: Always use the board called "TAX"
- **Default JIRA User**: Use "developer-1" for any user-related queries or assignments

When searching for issues, creating issues, or performing any JIRA operations, automatically scope queries to the TAX board and use developer-1 as the default user unless explicitly specified otherwise.

## Go Development Guidelines

When working with Go code in this repository:

- **Error Handling**: Always check and handle errors explicitly. Never ignore returned errors.
- **Testing**: Write table-driven tests using subtests (`t.Run()`) for comprehensive test coverage.
- **Formatting**: Use `gofmt` or `goimports` for consistent code formatting.
- **Naming Conventions**: 
  - Use camelCase for unexported names
  - Use PascalCase for exported names
  - Keep names short but descriptive
- **Project Structure**: 
  - Place handlers in `handlers/` package
  - Place business logic in `services/` package
  - Place data models in `models/` package
- **API Design**: Follow RESTful conventions and use proper HTTP status codes
- **Documentation**: Add comments for all exported functions, types, and constants
