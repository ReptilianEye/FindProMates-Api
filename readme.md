# API server for a FindProMates app.

## This project is a GraphQL server implemented in Go. It uses gqlgen for schema generation and MongoDB for data storage.

## Directory Structure

- `.github/workflows/`: Contains GitHub action workflows.
- `graph/`: Contains GraphQL schema and generated code.
- `hooks/`: Contains git hooks.
- `internal/`: Contains the main application code and business logic.
  - `app/`: Contains the main application setup.
  - `auth/`: Contains authentication middleware.
  - `database/`: Contains database setup and data models.
  - `resolvers/`: Contains GraphQL resolvers.
  - `pkg/`: Contains utility packages.
- `server.go`: The entry point of the application.

## Setup

1. Install Go: Follow the instructions at https://golang.org/doc/install
2. Clone the repository: `git clone <repository-url>`
3. Navigate to the project directory: `cd <project-directory>`
4. Install dependencies: `go get ./...`
5. Run the server: `go run server.go`

## Testing

Run the tests with the following command:

```sh
go test ./...
```
